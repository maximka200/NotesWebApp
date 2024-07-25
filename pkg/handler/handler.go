package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maximka200/NotesWebApp/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

// / создаем маршуты и возвращаем router c init routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	// POST - добавить в db новое, PUT - обноновить старое
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
