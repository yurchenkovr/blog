package controllers

import (
	"blog/src/models"
	"blog/src/usecases"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	g := e.Group("/a")

	auth := g.Group("/log", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("my_secret_key"),
		TokenLookup: "cookie:token",
	}))
	adm := auth.Group("/admin")

	g.GET("", articleHTTPsvc.GetArticles)                  // Any user can see articles
	g.GET("/:id", articleHTTPsvc.GetArticleByID)           //			-||-
	g.GET("/name/:username", articleHTTPsvc.GetByUsername) //			-||-

	auth.POST("", articleHTTPsvc.CreateArticle)   //Auth user can Create only his article
	auth.DELETE("/:id", articleHTTPsvc.DeleteArt) // 			-||- Delete -||-
	auth.PATCH("/:id", articleHTTPsvc.UpdateArt)  // 			-||= Update -||-

	adm.DELETE("/:id", articleHTTPsvc.DeleteAnyArticle) // Admin can Delete any article
}
func (s serviceArt) GetByUsername(c echo.Context) error {
	username := c.Param("username")

	article, err := s.svc.GetByUsername(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, article)
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
	claims, err := getClaims(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if claims.Blocked == true {
		return c.String(http.StatusBadRequest, "Your profile was blocked.")
	}
	req := usecases.CreateReqArt{
		Title:    article.Title,
		Username: claims.Username,
		Content:  article.Content,
		UserID:   claims.ID,
	}
	err = s.svc.SaveArticle(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
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
		log.Printf("Fail unmarshaling in updateArticle: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return errID
	}
	req := usecases.UpdateReqArt{
		Title:   article.Title,
		Content: article.Content,
	}
	err = s.svc.UpdateArticle(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "We got your updated article")
}
func (s *serviceArt) GetArticleByID(c echo.Context) error {
	id := getParamID(c)
	article, err := s.svc.GetByIDArticle(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, article)
}
func (s *serviceArt) DeleteArticle(c echo.Context) error {
	id := getParamID(c)
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

func (s *serviceArt) isMine(c echo.Context) bool {
	id := getParamID(c)
	claims, errClaims := getClaims(c)
	if errClaims != nil {
		return false
	}
	art, err := s.svc.GetByIDArticle(id)
	if err != nil {
		return false
	}
	if art.Username != claims.Username {
		log.Printf("Error: Please choose your article\n")
		return false
	}
	return true
}

func (s *serviceArt) UpdateArt(c echo.Context) error {
	if mine := s.isMine(c); mine != true {
		log.Printf("Error: It`s not your article\n")
		return c.String(http.StatusBadRequest, "Please choose your article")
	}
	if err := s.UpdateArticle(c); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "")
}
func (s *serviceArt) DeleteArt(c echo.Context) error {
	if mine := s.isMine(c); mine != true {
		log.Printf("Error: It`s not your article\n")
		return c.String(http.StatusBadRequest, "Please choose your article")
	}
	if err := s.DeleteArticle(c); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "")
}
func (s *serviceArt) DeleteAnyArticle(c echo.Context) error {
	claims, err := getClaims(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if isAdmin(c, claims.ID) == true {
		if err := s.DeleteArticle(c); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.String(http.StatusOK, fmt.Sprintf("Deleted successfully\nRole: %s", claims.Role))
	}
	return c.String(http.StatusOK, "Deleted successfully")
}
