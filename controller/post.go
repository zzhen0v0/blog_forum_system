package controller

import (
	"strconv"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	UserID, err := getCurrentUser(c)
	if err != nil {
		zap.L().Error("GetCurrentUser failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	post.AuthorId = UserID

	err = logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSucess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostByID(id)
	if err != nil {
		zap.L().Error("logic.GetPostByID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSucess(c, data)
}

func GetPost2ListHandler(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}

	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	data, err := logic.GetPost2List(page, size)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSucess(c, data)
	return
}

func GetPostListHandler(c *gin.Context) {
	p := &models.ParamPost{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("c.ShouldBindQuery() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	data, err := logic.GetPostList(p)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	ResponseSucess(c, data)
	return
}

func GetPostListByCommunityHandler(c *gin.Context) {
	p := &models.ParamPostCommunity{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("c.ShouldBindQuery() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	data, err := logic.GetPostListByCommunityID(p)
	if err != nil {
		zap.L().Error("logic.GetPostListByCommunityID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	ResponseSucess(c, data)
	return
}
