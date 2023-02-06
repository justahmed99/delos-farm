package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/usecase"
)

type FarmHandler struct {
	farmUseCases usecase.FarmUseCases
}

func NewFarmHandler(farmUseCase usecase.FarmUseCases) *FarmHandler {
	return &FarmHandler{
		farmUseCases: farmUseCase,
	}
}

func (h *FarmHandler) GetFarmById(context *gin.Context) {
	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	farm, err := h.farmUseCases.GetFarmByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Farm data not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, farm)
}

func (h *FarmHandler) GetFarms(context *gin.Context) {
	farms, err := h.farmUseCases.GetFarms()
	if err != nil {
		context.IndentedJSON(http.StatusNoContent, gin.H{"message": "No content"})
	}
	context.IndentedJSON(http.StatusOK, farms)
}

func (h *FarmHandler) CreateFarm(context *gin.Context) {

	var farm domain.Farm

	if err := context.ShouldBindJSON(&farm); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	updated_farm, err := h.farmUseCases.CreateFarm(farm)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	context.IndentedJSON(http.StatusOK, updated_farm)
}

func (h *FarmHandler) UpdateFarm(context *gin.Context) {

	var farm domain.Farm

	if err := context.ShouldBindJSON(&farm); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updated_pond, err := h.farmUseCases.UpdateFarm(farm)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusOK, updated_pond)
		return
	}
	context.IndentedJSON(http.StatusOK, updated_pond)
}

func (h *FarmHandler) DeleteFarm(context *gin.Context) {
	id, err_input := strconv.ParseInt(context.Param("id"), 10, 64)
	if err_input != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "input must be integer"})
		return
	}
	_, err := h.farmUseCases.GetFarmByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Farm data not found"})
		return
	}

	delete_err := h.farmUseCases.DeleteFarm(id)
	if delete_err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Farm data successufully deleted"})
}
