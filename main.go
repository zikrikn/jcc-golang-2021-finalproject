package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//making a lot of varian of endpoints
	r.GET("/", rootHandler) //GET untuk mengambil atau meminta data atau consume
	r.GET("/hello", helloHandler)
	r.GET("movies/:id/:title", booksHandler)
	r.GET("/query", queryHandler)
	r.POST("/movies", postBooksHandler) //POST untuk meng-create sesuatu

	r.Run("localhost:8888") // nanti bisa di Run() aja untuk di deploy di website
	//untuk local pake localhost:8080 dulu
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"name":    "Zikri KN is Awesome",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping",
		"name":    "Zikri KN is Nice",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

type moviesInput struct {
	Title string
	Price int `json:"price"` 
	SubTitle string `json:"sub_title"` //untuk mendeklarasikan di variabel json
}

func postBooksHandler(c *gin.Context) {
	//title, ranting
	var movieInput moviesInput

	err := c.ShouldBindJSON(&movieInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title": movieInput.Title,
		"price": movieInput.Price,
		"sub_title": movieInput.SubTitle,
	})
}

//3. HTTPS POST & VALIDATION