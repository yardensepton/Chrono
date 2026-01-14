package routeres

import (
	"github.com/gin-gonic/gin"
)

func RegisterClinicRoutes(router *gin.Engine) {
	group := router.Group("/clinics")
	{
		group.GET("/", getClinics)
		group.GET("/:id", getClinic)
	}
}

func getClinics(c *gin.Context) {
	c.JSON(200, gin.H{"Clinics": []string{"clinic1", "clinic2"}})
}

func getClinic(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": "Clinic with ID " + id})
}
