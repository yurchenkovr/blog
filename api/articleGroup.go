package api

import (
	"blog/api/handlers"
	"github.com/labstack/echo"
)

func ArticleGroup(g *echo.Group) {
	g.GET("/:data", handlers.GetArticles)
	g.POST("/new", handlers.AddArticle)
}
