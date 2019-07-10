package api

import (
	"github.com/labstack/echo"
	"blog/api/handlers"
)

func ArticleGroup(g *echo.Group) {
	g.GET("/:data", handlers.GetArticles)
	g.POST("/new", handlers.AddArticle)
}
