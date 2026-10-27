package handlers

import (
	"contacts-api/internal/database"
	"contacts-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	db database.ContactDB
}

func NewContactHandler(db database.ContactDB) *ContactHandler {
	return &ContactHandler{db: db}
}

func (h *ContactHandler) GetContacts(c *gin.Context) {
	// Check for search query
	query := c.Query("q")
	var contacts []models.Contact
	var err error

	if query != "" {
		contacts, err = h.db.Search(query)
	} else {
		contacts, err = h.db.GetAll()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to response format
	response := make([]models.ContactResponse, len(contacts))
	for i, contact := range contacts {
		response[i] = contact.ToResponse()
	}

	c.JSON(http.StatusOK, gin.H{
		"contacts": response,
		"count":    len(response),
	})
}

func (h *ContactHandler) GetContact(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	contact, err := h.db.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contact.ToResponse())
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var req models.CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact := &models.Contact{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		Company:   req.Company,
		JobTitle:  req.JobTitle,
		Notes:     req.Notes,
	}

	if err := h.db.Create(contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, contact.ToResponse())
}

func (h *ContactHandler) UpdateContact(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	// Get existing contact
	contact, err := h.db.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var req models.UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update only provided fields
	if req.FirstName != "" {
		contact.FirstName = req.FirstName
	}
	if req.LastName != "" {
		contact.LastName = req.LastName
	}
	if req.Email != "" {
		contact.Email = req.Email
	}
	if req.Phone != "" {
		contact.Phone = req.Phone
	}
	if req.Address != "" {
		contact.Address = req.Address
	}
	if req.Company != "" {
		contact.Company = req.Company
	}
	if req.JobTitle != "" {
		contact.JobTitle = req.JobTitle
	}
	if req.Notes != "" {
		contact.Notes = req.Notes
	}

	if err := h.db.Update(contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contact.ToResponse())
}

func (h *ContactHandler) DeleteContact(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	if err := h.db.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}