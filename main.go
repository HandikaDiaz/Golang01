package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/form-blog", formAddBlog)
	e.GET("/blog-detail/:id", blogDetail)

	e.POST("/add-blog", addBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	data := map[string]interface{}{
		"Login": true,
	}

	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddBlog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":      id,
		"Title":   "Yamada is Pro Player in Game",
		"Content": "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Saepe itaque omnis repellat aliquam repellendus et! Voluptatibus corrupti ratione cupiditate. Excepturi nulla soluta sed quasi omnis blanditiis inventore aliquam dicta quae.",
	}

	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("input-tittle")
	description := c.FormValue("input-description")

	println("Title : " + title)
	println("Description : " + description)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
