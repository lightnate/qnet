package main

import (
	"net/http"
	"qnet"
)

func main() {
	g := qnet.New()
	g.GET("/", func(c *qnet.Context) {
		c.HTML(http.StatusOK, "<h1>lbwnb</h1>")
	})

	g.GET("/hello", func(c *qnet.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"))
	})

	g.POST("/login", func(c *qnet.Context) {
		c.JSON(http.StatusOK, qnet.H{
			"user": c.PostForm("user"),
			"pwd":  c.PostForm("pwd"),
		})
	})

	g.Run(":9999")
}
