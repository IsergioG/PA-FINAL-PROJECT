package api

type KillRequestDto struct {
	FullName     string `json:"fullName"`
	FaceImageURL string `json:"faceImageUrl"`
	CauseOfDeath string `json:"causeOfDeath,omitempty"`
	Details      string `json:"details,omitempty"`
	DeathTime    string `json:"deathTime,omitempty"`
}

type KillResponseDto struct {
	ID             uint   `json:"id"`
	FullName       string `json:"fullName"`
	FaceImageURL   string `json:"faceImageUrl"`
	CauseOfDeath   string `json:"causeOfDeath,omitempty"`
	Details        string `json:"details,omitempty"`
	CreatedAt      string `json:"createdAt"`
	CauseWrittenAt string `json:"causeWrittenAt,omitempty"`
	DeathTime      string `json:"deathTime,omitempty"`
}

type KillTaskResponseDto struct {
	ID      uint   `json:"id"`
	Status  string `json:"status"`
	DeathAt string `json:"deathAt"` // cuándo morirá (fecha futura)
}

type ErrorResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Message     string `json:"message"`
}
