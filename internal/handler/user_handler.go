package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lokicodess/CatalogX/internal/domain"
	"github.com/lokicodess/CatalogX/internal/handler/dto"
	"github.com/lokicodess/CatalogX/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	u := &domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashBytes),
	}

	if err := h.Repo.Create(c.Request.Context(), u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	email := c.Query("email")
	u, err := h.Repo.GetByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
		return
	}
	c.JSON(http.StatusOK, u)
}
