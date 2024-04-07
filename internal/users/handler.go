package users

import (
	"encoding/json"
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

func (h *Handler) Autorization(c *gin.Context) {
	log := initlog.InitLogger()

	email := c.Request.URL.Query().Get("email")
	if email != "" {
		log.Error("Empty field email")
	}

	password := c.Request.URL.Query().Get("password")
	if password != "" {
		log.Error("Empty field email")
	}

	user, err := h.repository.FindOne(c.Request.Context(), email, password)
	if err != nil {
		log.Error(err.Error())
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Error(err.Error())
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(userBytes)
}
