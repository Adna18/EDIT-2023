package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BoredAPIResponse struct {
	Activity string `json:"activity"`
}

func main() {
	r := gin.Default()

	r.GET("/get_activity/:ime", func(c *gin.Context) {

		name := c.Param("ime")

		apiURL := fmt.Sprintf("https://www.boredapi.com/api/activity")
		response, err := http.Get(apiURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "neuspješan poziv api-ja."})
			return
		}

		defer response.Body.Close()

		var apiResponse BoredAPIResponse
		if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "neuspješno parsiranje odgovora."})
			return
		}
		fmt.Println(apiResponse)
		apiResponse.Activity += " " + name
		c.JSON(http.StatusOK, apiResponse)
	})

	r.Run(":8080")
}
