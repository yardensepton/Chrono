package routers

import (
	"net/http"

	"my-go-server/model/clinics"
	"my-go-server/services"

	"github.com/gin-gonic/gin"
)

func RegisterClinicRoutes(router *gin.Engine, clinicService *services.ClinicService) {
	group := router.Group("/clinics")
	{
		group.POST("/", createClinic(clinicService))
		group.GET("/", getClinics(clinicService))
		group.GET("/:id", getClinic(clinicService))
		group.PUT("/:id", updateClinic(clinicService))
		group.DELETE("/:id", deleteClinic(clinicService))
	}
}

func createClinic(service *services.ClinicService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req clinics.ClinicRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		clinic, err := service.InsertClinic(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "clinic_created", "clinic": clinic})
	}
}

func getClinics(service *services.ClinicService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// For now, return empty array. In a real app, you'd implement a GetAll method
		c.JSON(http.StatusOK, []clinics.Clinic{})
	}
}

func getClinic(service *services.ClinicService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		clinic, err := service.GetClinicByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "clinic_not_found"})
			return
		}
		c.JSON(http.StatusOK, clinic)
	}
}

func updateClinic(service *services.ClinicService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var clinic clinics.Clinic
		if err := c.ShouldBindJSON(&clinic); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
			return
		}
		clinic.ID = id
		updatedClinic, err := service.UpdateClinic(clinic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error"})
			return
		}
		c.JSON(http.StatusOK, updatedClinic)
	}
}

func deleteClinic(service *services.ClinicService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := service.DeleteClinic(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "clinic_deleted"})
	}
}
