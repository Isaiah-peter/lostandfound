package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/Isaiah-peter/lostandfound/db/sqlc"
	"github.com/Isaiah-peter/lostandfound/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	FullName  string         `json:"full_name"`
	Address   string `json:"address"`
	Contact   string         `json:"contact"`
	Username  string `json:"username"`
	UserImage string `json:"user_image"`
	Password  string `json:"password"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	hashpassword, _ := util.HashPassword(req.Password)

	fmt.Println(hashpassword)

	arg := db.CreateUserParams{
		FullName: req.FullName,
		Address: sql.NullString{String: req.Address, Valid: true},
		Contact: req.Contact,
		Username: sql.NullString{String: req.Username, Valid: true},
		UserImage: sql.NullString{String: req.UserImage, Valid: true},
		Password: sql.NullString{String: hashpassword, Valid: true},
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type GetUserRequest struct {
	ID int64 `uri:"id" binding: "required,min=1"`
}

type Userresponse struct {
	Id int32 `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Contact string `json:"contact"`
	Image string `json:"user_image"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	user, err := server.store.GetUser(ctx, int32(req.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userRes := Userresponse{
		Id: user.ID,
		FullName: user.FullName,
		Email: user.Address.String,
		Contact: user.Contact,
		Image: user.UserImage.String,
	}

	fmt.Println(userRes.Contact)

	

	ctx.JSON(http.StatusOK, userRes)
}
