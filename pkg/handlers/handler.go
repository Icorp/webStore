package handler

import (
	"github.com/Icorp/webStore/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	// r is router
	r := gin.New()

	// Auth group
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// Api group
	api := r.Group("/api")
	{
		lists := api.Group("/lists")
		{
			// home directory get lists
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)

			// id directory get list
			lists.POST("/:id", h.updateList)
			lists.GET("/:id", h.getListByID)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				// home directory get lists
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)

				// id directory get list
				items.POST("/:id", h.updateItem)
				items.GET("/:id", h.getItemById)
				items.DELETE("/:id", h.deleteItem)
			}
		}
	}

	return r
}
