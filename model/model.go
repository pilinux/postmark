package model

import (
	"time"

	"gorm.io/gorm"
)

// PostmarkOutbound - model to handle JSON outbound data from postmark
type PostmarkOutbound struct {
	ID          uint64 `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	RecordType  string         `json:"RecordType,omitempty"`
	Type        string         `json:"Type,omitempty"`
	TypeCode    int            `json:"TypeCode,omitempty"`
	MessageID   string         `json:"MessageID,omitempty"`
	Tag         string         `json:"Tag,omitempty"`
	From        string         `json:"From,omitempty"`
	To          string         `json:"To,omitempty"`
	Email       string         `gorm:"-" json:"Email,omitempty"`
	Recipient   string         `gorm:"-" json:"Recipient,omitempty"`
	EventAt     time.Time      `json:"EventAt,omitempty"`
	DeliveredAt time.Time      `gorm:"-" json:"DeliveredAt,omitempty"`
	BouncedAt   time.Time      `gorm:"-" json:"BouncedAt,omitempty"`
	ReceivedAt  time.Time      `gorm:"-" json:"ReceivedAt,omitempty"`
	ChangedAt   time.Time      `gorm:"-" json:"ChangedAt,omitempty"`
	ServerID    int            `json:"ServerID,omitempty"`
}
