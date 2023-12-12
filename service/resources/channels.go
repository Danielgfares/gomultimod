package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dgf.io/inter/service/models"
)

var channels = []models.Channel{
	{ID: "1", Name: "channel 1", Context: "grocery"},
	{ID: "2", Name: "channel 2", Context: "cloth"},
	{ID: "3", Name: "channel 3", Context: "greetings"},
}

type Channels struct {
}

// Handler to get channels
// Example
// curl http://localhost:8080/channels \
// --header "Content-Type: application/json" \
// --request "GET"
func (Channels) Get(c *gin.Context) {
	id := c.Param("id")

	if id != "" {
		for _, a := range channels {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	c.IndentedJSON(http.StatusOK, channels)
}

// Handler to add channels
// Example
// curl http://localhost:8080/channels \
// --include \
// --header "Content-Type: application/json" \
// --request "POST" \
// --data '{"id": "4", "name": "channel 4","context": "Services"}'
func (Channels) Post(c *gin.Context) {
	var channel models.Channel

	if err := c.BindJSON(&channel); err != nil {
		return
	}

	channels = append(channels, channel)
	c.IndentedJSON(http.StatusCreated, channel)
}
