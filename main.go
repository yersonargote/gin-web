package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

func IndexHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world",
	})
}

func NewRecipeHandler(ctx *gin.Context) {
	var recipe Recipe
	if err := ctx.ShouldBindJSON(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	ctx.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.POST("/recipes", NewRecipeHandler)
	router.Run()
}

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}
