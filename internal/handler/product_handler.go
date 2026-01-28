package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lokicodess/CatalogX/internal/domain"
	"github.com/lokicodess/CatalogX/internal/handler/dto"
	"github.com/lokicodess/CatalogX/internal/repository"
)

type ProductHandler struct {
	Repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	// 1) Bind JSON to DTO
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	// basic validation
	if strings.TrimSpace(req.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	if req.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be > 0"})
		return
	}

	// 2) Create domain product from DTO
	p := &domain.Product{
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		Sku:           req.SKU,
		StockQuantity: req.StockQuantity,
		IsActive:      true,
	}

	// 3) Generate slug
	p.Slug = generateSlug(p.Name)

	// 4) Save to DB via repo
	if err := h.Repo.Create(c.Request.Context(), p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	// 5) Return created product
	c.JSON(http.StatusCreated, p)
}

func generateSlug(name string) string {
	s := strings.ToLower(strings.TrimSpace(name))
	s = strings.ReplaceAll(s, " ", "-")

	out := make([]rune, 0, len(s))
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '-' {
			out = append(out, ch)
		}
	}

	slug := string(out)
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}
	return strings.Trim(slug, "-")
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.Repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.Repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
