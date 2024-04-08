package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xlzpm/internal/users/db/pg"
	"github.com/xlzpm/internal/users/model"
	"github.com/xlzpm/pkg/logger/initlog"
)

type Handler struct {
	repository *pg.Repository
}

func NewHandler(repo *pg.Repository) *Handler {
	return &Handler{
		repository: repo,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/sign-up", h.Register)
	router.GET("/sign-in", h.Autorization)

	return router
}

func (h *Handler) Register(c *gin.Context) {
	var user model.User

	log := initlog.InitLogger()

	if err := c.BindJSON(&user); err != nil {
		log.Error(err.Error())
	}

	err := h.repository.Create(c.Request.Context(), &user)
	if err != nil {
		log.Error(err.Error())
	}
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Autorization(c *gin.Context) {
	log := initlog.InitLogger()
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		log.Error(err.Error())
	}

	user, err := h.repository.FindOne(c.Request.Context(), input.Email, input.Password)
	if err != nil {
		log.Error(err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	})
}
