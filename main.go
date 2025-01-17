package main

import (
	"net/http"
	"strconv"

	"github.com/bayuiqballl/http-services/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func app(e *echo.Echo, store model.ArticleStore) {
	// curl http://localhost:8080/articles
	e.GET("/articles", func(c echo.Context) error {
		// Process
		articles := store.All()

		// Response
		return c.JSON(http.StatusOK, articles)
	})

	// curl http://localhost:8080/articles/1
	e.GET("/articles/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		article := store.Find(id)

		// Response
		return c.JSON(http.StatusOK, article)
	})

	// curl -d "title=Hello title&body=ini body article" http://localhost:8080/articles
	e.POST("/articles", func(c echo.Context) error {
		// Given
		title := c.FormValue("title")
		body := c.FormValue("body")

		// Create instabce
		article, _ := model.CreateArticle(title, body)

		// Persist
		store.Save(article)

		// Response
		return c.JSON(http.StatusOK, article)
	})

	// curl -X PUT -d "title=Hello title&body=ini body article" http://localhost:8080/articles/1
	e.PUT("/articles/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		article := store.Find(id)
		article.Title = c.FormValue("title")
		article.Body = c.FormValue("body")

		// Persists
		store.Update(article)

		// Response
		return c.JSON(http.StatusOK, article)
	})

	// curl -X DELETE http://localhost:8080/articles/1
	e.DELETE("/articles/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		article := store.Find(id)

		// Remove
		store.Delete(article)

		// Response
		return c.JSON(http.StatusOK, article)
	})
}

func main() {
	// init data store
	// store := model.NewArticleStoreMySQL()
	// driver := "mysql"
	// if driver == "inmemory" {
	// 	store = model.NewArticleStoreInmemory()
	// } else {
	// 	store = model.NewArticleStoreMySQL()
	// }

	store := model.NewArticleStoreMySQL()

	// Create new instance echo framework
	e := echo.New()

	// our apps
	app(e, store)

	e.Logger.Fatal(e.Start(":8080"))
}
