package handler

import (
	"net/http"

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
	id := context.Param("id")
	farm, err := h.farmUseCases.GetFarmByID(context, id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Farm data not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, farm)
}

func (h *FarmHandler) GetFarms(context *gin.Context) {
	farms, err := h.farmUseCases.GetFarms(context)
	if err != nil {
		context.IndentedJSON(http.StatusNoContent, gin.H{"message": "No content"})
		return
	}
	context.IndentedJSON(http.StatusOK, farms)
}

func (h *FarmHandler) CreateFarm(context *gin.Context) {

	var farm domain.Farm

	if err := context.ShouldBindJSON(&farm); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	created_farm, err := h.farmUseCases.CreateFarm(context, &farm)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, created_farm)
}

func (h *FarmHandler) UpdateFarm(context *gin.Context) {

	var farm domain.Farm

	if err := context.ShouldBindJSON(&farm); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updated_pond, err := h.farmUseCases.UpdateFarm(context, &farm)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusOK, updated_pond)
		return
	}
	context.IndentedJSON(http.StatusOK, updated_pond)
}

func (h *FarmHandler) DeleteFarm(context *gin.Context) {
	id := context.Param("id")
	delete_err := h.farmUseCases.DeleteFarm(context, id)
	if delete_err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": delete_err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Farm data successufully deleted"})
}
