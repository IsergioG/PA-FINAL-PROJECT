package models

import (
	"backend-avanzada/api"
	"time"
)

func formatTime(t *time.Time) string {
	if t != nil {
		return t.Format(time.RFC3339)
	}
	return ""
}

type Kill struct {
	ID uint `gorm:"primaryKey" json:"id"`

	FullName       string     `gorm:"-" json:"fullName"`
	FaceImageURL   string     `gorm:"not null" json:"faceImageUrl"`
	CauseOfDeath   string     `json:"causeOfDeath,omitempty"`
	Details        string     `json:"details,omitempty"`
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	CauseWrittenAt *time.Time `json:"CauseWrittenAt,omitempty"`
	DeathTime      *time.Time `json:"deathTime,omitempty"`
}

func (k *Kill) ToKillResponseDto() *api.KillResponseDto {
	return &api.KillResponseDto{
		ID:             k.ID,
		FullName:       k.FullName,
		FaceImageURL:   k.FaceImageURL,
		CauseOfDeath:   k.CauseOfDeath,
		Details:        k.Details,
		CreatedAt:      k.CreatedAt.Format(time.RFC3339),
		CauseWrittenAt: formatTime(k.CauseWrittenAt),
		DeathTime:      formatTime(k.DeathTime),
	}
}
