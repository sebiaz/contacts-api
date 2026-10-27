package database

import (
	"contacts-api/internal/models"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type ContactDB interface {
	GetAll() ([]models.Contact, error)
	GetByID(id uint) (*models.Contact, error)
	Create(contact *models.Contact) error
	Update(contact *models.Contact) error
	Delete(id uint) error
	Search(query string) ([]models.Contact, error)
}

type contactDB struct {
	db *gorm.DB
}

func NewContactDB(db *gorm.DB) ContactDB {
	return &contactDB{db: db}
}

func (c *contactDB) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := c.db.Order("created_at DESC").Find(&contacts).Error
	return contacts, err
}

func (c *contactDB) GetByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := c.db.First(&contact, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact not found")
		}
		return nil, err
	}
	return &contact, nil
}

func (c *contactDB) Create(contact *models.Contact) error {
	return c.db.Create(contact).Error
}

func (c *contactDB) Update(contact *models.Contact) error {
	return c.db.Save(contact).Error
}

func (c *contactDB) Delete(id uint) error {
	result := c.db.Delete(&models.Contact{}, id)
	if result.RowsAffected == 0 {
		return errors.New("contact not found")
	}
	return result.Error
}

func (c *contactDB) Search(query string) ([]models.Contact, error) {
	var contacts []models.Contact
	// Use LOWER() for case-insensitive search in SQLite
	lowerQuery := "%" + strings.ToLower(query) + "%"
	err := c.db.Where("LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(email) LIKE ? OR LOWER(company) LIKE ?",
		lowerQuery, lowerQuery, lowerQuery, lowerQuery).
		Order("created_at DESC").Find(&contacts).Error
	return contacts, err
}