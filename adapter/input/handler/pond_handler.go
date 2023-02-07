package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/usecase"
)

type PondHandler struct {
	pondUseCase usecase.PondUseCases
}

func NewPondHandler(pondUseCase usecase.PondUseCases) *PondHandler {
	return &PondHandler{
		pondUseCase: pondUseCase,
	}
}

func (h *PondHandler) GetPondById(context *gin.Context) {
	id := context.Param("id")
	pond, err := h.pondUseCase.GetPondByID(context, id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, pond)
}

func (h *PondHandler) GetPonds(context *gin.Context) {
	pond, err := h.pondUseCase.GetPonds(context)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, pond)
}

func (h *PondHandler) GetPondsByFarmID(context *gin.Context) {
	farm_id := context.Param("farm_id")
	pond, err := h.pondUseCase.GetPondsByFarmID(context, farm_id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, pond)
}

func (h *PondHandler) CreatePond(context *gin.Context) {

	var pond domain.Pond

	if err := context.ShouldBindJSON(&pond); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	created_pond, err := h.pondUseCase.CreatePond(context, pond)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, created_pond)
}

func (h *PondHandler) UpdatePond(context *gin.Context) {

	var pond domain.Pond

	if err := context.ShouldBindJSON(&pond); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	created_pond, err := h.pondUseCase.UpdatePond(context, pond)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, created_pond)
}

func (h *PondHandler) DeletePond(context *gin.Context) {
	id := context.Param("id")
	delete_err := h.pondUseCase.DeletePond(context, id)
	if delete_err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": delete_err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Pond data successufully deleted"})
}

func (h *PondHandler) DeletePondsByFarmID(context *gin.Context) {
	id := context.Param("farm_id")
	delete_err := h.pondUseCase.DeletePondsByFarmID(context, id)
	if delete_err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": delete_err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Pond data successufully deleted"})
}
