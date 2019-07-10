package handlers

import (
	"blog/api/models"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

func GetArticles(c echo.Context) error {
	id := c.QueryParam("id")
	title := c.QueryParam("title")
	author := c.QueryParam("author")
	content := c.QueryParam("content")

	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("id: %s\ntitle: %s\nauthor: %s\nContent: %s", id, title, author, content))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"id":      id,
			"title":   title,
			"author":  author,
			"content": content,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "please choose the type",
	})
}

func AddArticle(c echo.Context) error {
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
	log.Printf("This is your article: %#v", article)
	return c.String(http.StatusOK, "We got your article")
}
