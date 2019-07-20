package controllers

import (
	"blog/src/models"
	"blog/src/usecases"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type serviceArt struct {
	svc usecases.ArticleService
}

func NewService(e echo.Echo, artService usecases.ArticleService) {
	articleHTTPsvc := serviceArt{svc: artService}

	g := e.Group("/articles")

	g.GET("", articleHTTPsvc.GetArticles)
	g.POST("", articleHTTPsvc.CreateArticle)
	g.DELETE("/:id", articleHTTPsvc.DeleteArticle)
	g.GET("/:id", articleHTTPsvc.GetArticleByID)
	g.PATCH("/:id", articleHTTPsvc.UpdateArticle)
}

func (s serviceArt) CreateArticle(c echo.Context) error {
	article := models.Article{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &article)
	if err != nil {
		log.Printf("Fail unmarshaling in addArticle: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	req := usecases.CreateReqArt{
		ID:      article.ID,
		Title:   article.Title,
		Author:  article.Author,
		Content: article.Content,
	}
	err = s.svc.SaveArticle(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Printf("This is your article: %#v", article)
	return c.String(http.StatusOK, "We got your article")
}

func (s *serviceArt) UpdateArticle(c echo.Context) error {
	article := models.Article{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &article)
	if err != nil {
		log.Printf("Fail unmarshaling in addArticle: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return errID
	}
	req := usecases.UpdateReqArt{
		Title:   article.Title,
		Author:  article.Author,
		Content: article.Content,
	}
	err = s.svc.UpdateArticle(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Printf("This is your article: %#v", article)
	return c.String(http.StatusOK, "We got your updated article")
}
func (s *serviceArt) GetArticleByID(c echo.Context) error {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return errID
	}
	article, err := s.svc.GetByIDArticle(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, article)
}
func (s *serviceArt) DeleteArticle(c echo.Context) error {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return errID
	}
	err := s.svc.DeleteArticle(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "Deleted successfully")
}
func (s *serviceArt) GetArticles(c echo.Context) error {
	articles, err := s.svc.GetAllArticles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, articles)
}
