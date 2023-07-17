package api

import (
	"net/http"

	db "github.com/Isaiah-peter/lostandfound/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createLostItemReq struct {
	Title       string `json:"title"`
	Discription string `json:"discription"`
}

func (server *Server) createLostItem(ctx *gin.Context) {
	var req createLostItemReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	arg := db.CreateLostItemParams{
		Title:       req.Title,
		Discription: req.Discription,
	}

	cat, err := server.store.CreateLostItem(ctx, arg)

	if err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, cat)
}

type GetLostItemReq struct {
	Id int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getLostItem(ctx *gin.Context) {
	var req GetLostItemReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	LostItem, err := server.store.GetLostItem(ctx, req.Id)

	if err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, LostItem)
}

type ListLostItemReq struct {
	PageId   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

type StatusReq struct {
	Status string `form:"status"`
}

func (server *Server) ListLostItem(ctx *gin.Context) {
	var req ListLostItemReq

	var status StatusReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	if err := ctx.ShouldBindQuery(&status); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	arg := db.ListLostItemParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}

	categories, err := server.store.ListLostItem(ctx, arg)

	if err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}

	if status.Status != "" {
		var result []db.LostItem

		for _, element := range categories {
			// Check if the element has the desired property
			if element.Status == db.ItemStatus(status.Status) {
				result = append(result, element)
			}
		}

		ctx.JSON(http.StatusOK, result)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

type UpdateLostItemReq struct {
	Status db.ItemStatus `json:"status"`
}

func (server *Server) updateLostItem(ctx *gin.Context) {
	var input UpdateLostItemReq
	var id GetLostItemReq

	if err := ctx.ShouldBindJSON(&input); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	if err := ctx.ShouldBindUri(&id); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	arg := db.UpdateLostItemStatusParams{
		ID:     id.Id,
		Status: input.Status,
	}

	if err := server.store.UpdateLostItemStatus(ctx, arg); err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "succesfully update LostItem"})
}

func (server *Server) deleteLostItem(ctx *gin.Context) {
	var req GetLostItemReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	err := server.store.DeleteLostItem(ctx, req.Id)

	if err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "succesfully update LostItem"})
}
