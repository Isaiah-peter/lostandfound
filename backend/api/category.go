package api

import (
	"net/http"

	db "github.com/Isaiah-peter/lostandfound/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCategoryReq struct {
	Title       string `json:"title"`
	Discription string `json:"discription"`
}

func (server *Server) bindingError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code , errorResponse(err))
	return
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	arg := db.CreateCategoryParams{
		Title: req.Title,
		Discription: req.Discription,
	}

	cat, err := server.store.CreateCategory(ctx, arg)

	if err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, cat)
}

type GetCategoryReq struct {
	Id int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	var req GetCategoryReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	category, err := server.store.GetCategory(ctx, req.Id)

	if err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, category)
}

type ListCategoryReq struct {
	PageId int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (server *Server) ListCategory(ctx *gin.Context) {
	var req ListCategoryReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	arg := db.ListCategoryParams{
		Limit: req.PageSize,
		Offset: (req.PageId -1) * req.PageSize,
	}

	categories, err := server.store.ListCategory(ctx, arg)

	if err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, categories)
}

type UpdateCategoryReq struct {
	Title string `json:"title"`
	Discription string `json:"discription"`
}

func (server *Server) updateCategory(ctx *gin.Context) {
	var input UpdateCategoryReq

	var id GetCategoryReq

	if err := ctx.ShouldBindJSON(&input); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	if err := ctx.ShouldBindUri(&id); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	cat, err := server.store.GetCategory(ctx, id.Id)
	if err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	arg := db.UpdateCategoryParams{
		ID: id.Id,
		Title: input.Title,
		Discription: input.Discription,
	}

	if arg.Discription == "" {
		arg.Discription = cat.Discription
	}

	if arg.Title == "" {
		arg.Title = cat.Title
	}

	if err := server.store.UpdateCategory(ctx, arg); err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "succesfully update category"})
}

func (server *Server) deleteCategory(ctx *gin.Context) {
	var req GetCategoryReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		server.bindingError(ctx, http.StatusBadRequest, err)
	}

	err := server.store.DeleteCategory(ctx, req.Id)

	if err != nil {
		server.bindingError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "succesfully update category"})
}