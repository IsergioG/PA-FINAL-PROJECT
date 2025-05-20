package server

import (
	"backend-avanzada/api"
	"backend-avanzada/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) HandleKills(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllKills(w, r)
		return
	case http.MethodPost:
		s.handleCreateKill(w, r)
		return
	}
}

func (s *Server) HandleKillsWithId(w http.ResponseWriter, r *http.Request) {
	s.handleGetKillById(w, r)
}
func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	s.HandleUpdateKillById(w, r)
}

func (s *Server) handleGetAllKills(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	result := []*api.KillResponseDto{}
	kills, err := s.KillRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	for _, v := range kills {
		result = append(result, v.ToKillResponseDto())
	}
	response, err := json.Marshal(result)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

func (s *Server) handleGetKillById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}
	k, err := s.KillRepository.FindById(int(id))
	if k == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path, fmt.Errorf("person with id %d not found", id))
		return
	}
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	resp := k.ToKillResponseDto()
	response, err := json.Marshal(resp)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

func (s *Server) handleCreateKill(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, fmt.Errorf("formato inválido: %v", err))
		return
	}

	fullName := r.FormValue("fullName")
	causeOfDeath := r.FormValue("causeOfDeath")
	details := r.FormValue("details")

	if fullName == "" {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, fmt.Errorf("firstName y lastName son requeridos"))
		return
	}

	// Guardar foto
	file, handler, err := r.FormFile("photo")
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, fmt.Errorf("archivo de imagen requerido"))
		return
	}
	defer file.Close()

	os.MkdirAll("uploads", os.ModePerm)
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
	savePath := fmt.Sprintf("uploads/%s", fileName)
	publicURL := fmt.Sprintf("/static/%s", fileName)

	dst, err := os.Create(savePath)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, fmt.Errorf("no se pudo guardar imagen"))
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	// Crear objeto Kill
	now := time.Now()
	kill := &models.Kill{
		FullName:     fullName,
		FaceImageURL: publicURL,
		CauseOfDeath: causeOfDeath,
		Details:      details,
		CreatedAt:    now,
	}

	// Guardar en la base de datos
	saved, err := s.KillRepository.Save(kill)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	// Responder con DTO estándar
	resp := api.KillResponseDto{
		ID:           saved.ID,
		FullName:     saved.FullName,
		CauseOfDeath: saved.CauseOfDeath,
		Details:      saved.Details,
		FaceImageURL: saved.FaceImageURL,
		CreatedAt:    saved.CreatedAt.Format(time.RFC3339),
	}

	result, err := json.Marshal(resp)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
	s.logger.Info(http.StatusCreated, r.URL.Path, start)
}

func (s *Server) HandleUpdateKillById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var dto api.KillRequestDto
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	kill, err := s.KillRepository.FindById(id)
	if err != nil || kill == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path, fmt.Errorf("Kill not found"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	now := time.Now()
	elapsed := now.Sub(kill.CreatedAt)

	// Validar y asignar causa si no existe aún
	if dto.CauseOfDeath != "" && kill.CauseOfDeath == "" {
		if elapsed > 40*time.Second {
			s.HandleError(w, http.StatusForbidden, r.URL.Path, fmt.Errorf("La causa solo puede escribirse dentro de los primeros 40 segundos"))
			return
		}
		kill.CauseOfDeath = dto.CauseOfDeath
		kill.CauseWrittenAt = &now
	}

	// Validar y asignar detalles si no existe aún
	if dto.Details != "" && kill.Details == "" {
		if elapsed > (6*time.Minute + 40*time.Second) {
			s.HandleError(w, http.StatusForbidden, r.URL.Path, fmt.Errorf("Los detalles solo pueden escribirse dentro de los primeros 6:40 minutos"))
			return
		}
		kill.Details = dto.Details
		kill.CauseWrittenAt = &now
	}

	// Permitir actualización de DeathTime si se pasa explícitamente
	if dto.DeathTime != "" {
		parsedDeathTime, err := time.Parse(time.RFC3339, dto.DeathTime)
		if err != nil {
			s.HandleError(w, http.StatusBadRequest, r.URL.Path, fmt.Errorf("Formato de hora inválido"))
			return
		}
		kill.DeathTime = &parsedDeathTime
	}

	updated, err := s.KillRepository.Update(id, kill)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response, err := json.Marshal(updated.ToKillResponseDto())
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}
