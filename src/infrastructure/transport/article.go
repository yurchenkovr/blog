package transport

import (
	"blog/src/models"
	"blog/src/usecases"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

type serviceArt struct {
	svc usecases.ArticleService
}

func NewService(e echo.Echo, artService usecases.ArticleService, middlewareFunc echo.MiddlewareFunc) {
	articleHTTPsvc := serviceArt{svc: artService}

	art := e.Group("/articles")

	art.GET("", articleHTTPsvc.List)
	art.GET("/:id", articleHTTPsvc.View)
	art.GET("/name/:username", articleHTTPsvc.GetByUsername)
	art.POST("", articleHTTPsvc.Create, middlewareFunc)
	art.DELETE("/:id", articleHTTPsvc.Delete, middlewareFunc)
	art.PATCH("/:id", articleHTTPsvc.Update, middlewareFunc)
}

func (s serviceArt) GetByUsername(c echo.Context) error {
	username := c.Param("username")

	article, err := s.svc.GetByUsername(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, article)
}

func (s serviceArt) Create(c echo.Context) error {
	article := models.Article{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the request body: %s", err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &article)
	if err != nil {
		log.Printf("Fail unmarshaling in addArticle: %s", err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	req := usecases.CreateReqArt{
		Title:    article.Title,
		Username: c.Get("username").(string),
		Content:  article.Content,
		UserID:   c.Get("id").(int),
	}

	err = s.svc.Create(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "We got your article")
}

func (s *serviceArt) Update(c echo.Context) error {
	article := models.Article{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &article)
	if err != nil {
		log.Printf("Fail unmarshaling in updateArticle: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	id := getParamID(c)

	req := usecases.UpdateReqArt{
		Title:   article.Title,
		Content: article.Content,
	}

	err = s.svc.Update(c, id, req)
	if err != nil {
		log.Printf("Error.Transport.Update: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, "We got your updated article")
}

func (s *serviceArt) View(c echo.Context) error {
	id := getParamID(c)

	article, err := s.svc.View(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, article)
}

func (s *serviceArt) Delete(c echo.Context) error {
	id := getParamID(c)

	err := s.svc.Delete(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, "Deleted successfully")
}

func (s *serviceArt) List(c echo.Context) error {
	articles, err := s.svc.List()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, articles)
}
