package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world",
	})
}

func Greeting(ctx *gin.Context) {
	name := ctx.Params.ByName("name")
	ctx.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

func InfoPerson(ctx *gin.Context) {
	ctx.XML(200, Person{
		FirstName: "Yerson",
		LastName:  "Argote",
	})
}

func InfoBreakfast(ctx *gin.Context) {
	ctx.JSON(200, Breakfast{
		Drink: "Coffe",
		Fruit: "Banana",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.GET("/:name", Greeting)
	router.GET("/person", InfoPerson)
	router.GET("/breakfast", InfoBreakfast)
	router.Run()
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

type Breakfast struct {
	Drink string `json:"drink"`
	Fruit string `json:"fruit"`
}

type Recipe struct {
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt  []string `json:"publishedAt"`
}
