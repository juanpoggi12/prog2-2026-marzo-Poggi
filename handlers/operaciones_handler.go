package handlers

import (
	"net/http"
	"strconv"
	"prog2-2026-marzo-Poggi/dto"
	"prog2-2026-marzo-Poggi/services"

	"github.com/gin-gonic/gin"
)

type OperacionesHandler struct {
	service services.OperacionesServicesInterface
}

func NewOperacionesHandler(service services.OperacionesServicesInterface) *OperacionesHandler{
	return &OperacionesHandler{
		service: service,
	}
}

func (handler *OperacionesHandler) CalcularPromedioVelocidad(c *gin.Context) {
	var request dto.PromedioVelocidadRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, response, err := handler.service.CalcularPromedioVelocidad(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (handler *OperacionesHandler) CalcularVariabilidadRendimiento(c *gin.Context) {
	var request dto.VariabilidadRendimientoRequest

	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, response, err := handler.service.CalcularVariabilidadRendimiento(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (handler *OperacionesHandler) CalcularProyeccionMejora(c *gin.Context) {
	rendimientoActualStr := c.Param("rendimiento_actual")
	tasaMejoraStr := c.Param("tasa_mejora")
	semanasStr := c.Param("semanas")

	rendimientoActual, err := strconv.ParseFloat(rendimientoActualStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasaMejora, err := strconv.ParseFloat(tasaMejoraStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	semanas, err := strconv.ParseInt(semanasStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.service.CalcularProyeccionMejora(rendimientoActual, tasaMejora, int(semanas))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}