package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	db "github.com/Isaiah-peter/lostandfound/db/sqlc"
	"github.com/Isaiah-peter/lostandfound/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	FullName  string `json:"full_name"`
	Address   string `json:"address"`
	Contact   string `json:"contact"`
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
		FullName:  req.FullName,
		Address:   req.Address,
		Contact:   req.Contact,
		Username:  req.Username,
		UserImage: req.UserImage,
		Password:  hashpassword,
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
	Id       int32  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Contact  string `json:"contact"`
	Image    string `json:"user_image"`
}

func newUserResponse(user db.User) Userresponse {
	return Userresponse{
		Id:       user.ID,
		FullName: user.FullName,
		Email:    user.Address,
		Contact:  user.Contact,
		Image:    user.UserImage,
	}
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

	userRes := newUserResponse(user)

	fmt.Println(userRes.Contact)

	ctx.JSON(http.StatusOK, userRes)
}

type loginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        Userresponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByUserName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user not found %s", req.Username)})

			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = util.CheckPassword(user.Password, req.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.ID, time.Duration(time.Duration.Minutes(15)))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}
