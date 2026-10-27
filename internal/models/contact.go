package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	FirstName string `gorm:"not null" json:"first_name" binding:"required"`
	LastName  string `gorm:"not null" json:"last_name" binding:"required"`
	Email     string `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Company   string `json:"company"`
	JobTitle  string `json:"job_title"`
	Notes     string `json:"notes"`
}

type CreateContactRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Company   string `json:"company"`
	JobTitle  string `json:"job_title"`
	Notes     string `json:"notes"`
}

type UpdateContactRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"omitempty,email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Company   string `json:"company"`
	JobTitle  string `json:"job_title"`
	Notes     string `json:"notes"`
}

type ContactResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Company   string    `json:"company"`
	JobTitle  string    `json:"job_title"`
	Notes     string    `json:"notes"`
}

func (c *Contact) ToResponse() ContactResponse {
	return ContactResponse{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		Address:   c.Address,
		Company:   c.Company,
		JobTitle:  c.JobTitle,
		Notes:     c.Notes,
	}
}