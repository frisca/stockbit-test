package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	response "stockbit-test/omdb/dto"
	"stockbit-test/omdb/service"

	"github.com/joho/godotenv"
)

var (
	envs map[string]string
	err  error
)

var OmdbService = new(service.OmdbService)

// List ...
func List(c *gin.Context) {
	var res response.Response

	envs, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	omdbkey := envs["OMDB_KEY"]
	search := c.Query("search")
	pagination, _ := strconv.Atoi(c.Query("pagination"))

	res = OmdbService.List(pagination, search, omdbkey)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}

// Get ...
func Get(c *gin.Context) {
	var res response.ResponseDetail

	envs, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	omdbkey := envs["OMDB_KEY"]
	imdbID := c.Param("imdbID")

	res = OmdbService.Get(imdbID, omdbkey)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
