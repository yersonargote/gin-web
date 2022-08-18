package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var recipes []Recipe

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func ListRecipesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, recipes)
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

func UpdateRecipeHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var recipe Recipe

	if err := ctx.ShouldBindJSON(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for index, value := range recipes {
		if value.ID == id {
			recipes[index] = recipe
			ctx.JSON(http.StatusOK, recipe)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Recipe not found",
	})
}

func FindRecipeHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, recipe := range recipes {
		if recipe.ID == id {
			ctx.JSON(http.StatusOK, recipe)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Recipe not found",
	})
}

func DeleteRecipeHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, recipe := range recipes {
		if recipe.ID == id {
			recipes = append(recipes[:idx], recipes[idx+1:]...)
			ctx.JSON(http.StatusOK, recipe)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Recipe not found",
	})
}

func SearchRecipeHandler(ctx *gin.Context) {
	tag := ctx.Query("tag")
	queryRecipes := make([]Recipe, 0)

	for _, recipe := range recipes {
		for _, value := range recipe.Tags {
			if strings.EqualFold(value, tag) {
				queryRecipes = append(queryRecipes, recipe)
				break
			}
		}
	}
	ctx.JSON(http.StatusOK, queryRecipes)
}

func init() {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.GET("/recipes/:id", FindRecipeHandler)
	router.GET("/recipes/search", SearchRecipeHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.Run()
}
