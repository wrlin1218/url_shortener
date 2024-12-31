package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wrlin1218/url_shortener/internal/service"
	"github.com/wrlin1218/url_shortener/pkg/logger"
	"net/http"
)

type LinkController struct {
	LinkService service.LinkService
}

func NewLinkController(linkService service.LinkService) *LinkController {
	return &LinkController{LinkService: linkService}
}

type CreateShortLinkRequest struct {
	Username    string `json:"username" binding:"required"`
	OriginalURL string `json:"original_url" binding:"required,url"`
}

func (lc *LinkController) CreateShortLink(c *gin.Context) {
	var req CreateShortLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "illegal params"})
		return
	}

	err, code := lc.LinkService.CreateShortLink(c, req.Username, req.OriginalURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Create short link failed, reason: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Short link created successfully", "code": code})
}

func (lc *LinkController) RedirectToOriginal(c *gin.Context) {
	shortCode := c.Param("short_code")
	err, original_url := lc.LinkService.GetOriginalUrl(c, shortCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get origin url failed, reason: " + err.Error()})
		return
	}
	logger.Info("original_url:" + original_url)
	c.Redirect(http.StatusFound, original_url)
}

func (lc *LinkController) DeleteShortLink(c *gin.Context) {
	username := c.Query("username")
	shortCode := c.Query("short_code")
	err := lc.LinkService.DeleteShortLink(c, username, shortCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delete short link failed, reason: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Short link deleted successfully", "code": http.StatusOK})
}
