package main

import (
	"net/http"
	"strconv"

	"github.com/bayuiqballl/http-services/model"
	"github.com/labstack/echo"
)

func main() {

	store := model.NewArticleStoreInMemory()

	e := echo.New()
	e.GET("/articles", func(c echo.Context) error {
		articles := store.ArticleMap
		return c.JSON(http.StatusOK, articles)
	})

	e.POST("/articles", func(c echo.Context) error {
		// 1. mengambil data dari form value
		title := c.FormValue("title")
		body := c.FormValue("body")

		// 2. create article instance
		article, _ := model.CreateArticle(title, body)

		// 3. save ke store
		store.Save(article)
		return c.JSON(http.StatusOK, article)
	})

	e.PUT("/articles/:id", func(c echo.Context) error {
		title := c.FormValue("title")
		body := c.FormValue("body")
		id, _ := strconv.Atoi(c.Param("id"))

		articles := store.ArticleMap
		articles[id-1] = model.Article{ID: id, Title: title, Body: body}

		return c.JSON(http.StatusOK, articles)
	})

	e.DELETE("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		articles := store.ArticleMap
		del := append(articles[:id-1], articles[id:]...)
		return c.JSON(http.StatusOK, del)
	})

	e.GET("/articles/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		articles := store.ArticleMap
		return c.JSON(http.StatusOK, articles[id-1])
	})

	e.Logger.Fatal(e.Start(":8080"))
}
