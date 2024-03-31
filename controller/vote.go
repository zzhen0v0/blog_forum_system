package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/dao/redis"

	"github.com/gin-gonic/gin"
)

type VoteData struct {
	PostID    string  `json:"post_id" binding:"required"`
	Direction float64 `json:"direction" binding:"required,omitempty"`
}

func VoteHandler(c *gin.Context) {
	//1 check param
	var vote VoteData
	if err := c.ShouldBindJSON(&vote); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs)
		return
	}

	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	//2 redis update
	if err := redis.PostVote(vote.PostID, fmt.Sprint(userID), vote.Direction); err != nil {
		zap.L().Error("redis.PostVote() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//3response
	ResponseSucess(c, nil)
}
