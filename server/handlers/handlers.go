package handlers

import (
	"gogin/data"
	"gogin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var mockData = data.Data.Albums

func Albums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mockData)
}

func GetAlbumById(c *gin.Context) {
	for _, album := range mockData {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithError(404, err)
		}
		if album.Id == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func CreateAlbumById(c *gin.Context) {
	var requestBody models.Album

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(500, err)
	}

	for index := range mockData {
		requestBody.Id = mockData[index].Id + 1
	}
	mockData = append(mockData, requestBody)
	c.IndentedJSON(http.StatusCreated, requestBody)
}

func UpdateAlbumById(c *gin.Context) {

	var requestBody models.Album

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(404, err)
	}

	for index, album := range mockData {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithError(404, err)
		}
		if album.Id == id {
			mockData[index].Price = requestBody.Price
			c.IndentedJSON(http.StatusOK, mockData[index])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func DeleteAlbumById(c *gin.Context) {
	for index, album := range mockData {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithError(404, err)
		}
		if album.Id == id {
			mockData = append(mockData[:index], mockData[index+1:]...)
			c.IndentedJSON(http.StatusOK, mockData)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
