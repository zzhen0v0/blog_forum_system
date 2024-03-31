package logic

import (
	"fmt"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	postID := snowflake.GenID()
	post.PostID = postID

	err = mysql.CreatePost(post)
	if err != nil {
		return err
	}
	err = redis.CreatePost(fmt.Sprint(post.PostID), fmt.Sprint(post.CommunityID))
	if err != nil {
		return
	}
	return
}

func GetPostByID(id int64) (data *models.ApiPostDetail, err error) {
	data, err = mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed", zap.String("post_id", string(id)), zap.Error(err))
		return nil, err
	}

	user, err := mysql.GetUserByID(data.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(id) failed", zap.String("post_id", string(data.AuthorId)), zap.Error(err))
		return nil, err
	}
	data.AuthorName = user.Username

	community, err := mysql.GetCommunityByID(data.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GGetCommunityByID(id) failed", zap.String("post_id", string(data.CommunityID)), zap.Error(err))
		return nil, err
	}
	data.CommunityName = community.Name

	return
}

func GetPost2List(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(id) failed", zap.String("post_id", string(post.AuthorId)), zap.Error(err))
			return nil, err
		}
		post.AuthorName = user.Username

		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GGetCommunityByID(id) failed", zap.String("post_id", string(post.CommunityID)), zap.Error(err))
			return nil, err
		}
		post.CommunityName = community.Name
		data = append(data, post)
	}
	return
}

func GetPostList(p *models.ParamPost) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostList(p)
	if err != nil {
		zap.L().Error("redis.GetPostList() failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostList() success,but no data")
		return
	}

	posts, err := mysql.GetPostByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostByIDs() failed", zap.Error(err))
		return
	}

	voteData, err := redis.GetVoteData(ids)
	if err != nil {
		zap.L().Error("redis.GetPostData() failed", zap.Error(err))
		return
	}

	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(id) failed", zap.String("post_id", string(post.AuthorId)), zap.Error(err))
			return nil, err
		}
		post.AuthorName = user.Username

		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GGetCommunityByID(id) failed", zap.String("post_id", string(post.CommunityID)), zap.Error(err))
			return nil, err
		}
		post.CommunityName = community.Name
		post.VoteNum = voteData[idx]
		data = append(data, post)
	}

	return
}

func GetPostListByCommunityID(p *models.ParamPostCommunity) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostListCommunityID(p)
	if err != nil {
		zap.L().Error("redis.GetPostList() failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostList() success,but no data")
		return
	}

	posts, err := mysql.GetPostByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostByIDs() failed", zap.Error(err))
		return
	}

	voteData, err := redis.GetVoteData(ids)
	if err != nil {
		zap.L().Error("redis.GetPostData() failed", zap.Error(err))
		return
	}

	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(id) failed", zap.String("post_id", string(post.AuthorId)), zap.Error(err))
			return nil, err
		}
		post.AuthorName = user.Username

		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GGetCommunityByID(id) failed", zap.String("post_id", string(post.CommunityID)), zap.Error(err))
			return nil, err
		}
		post.CommunityName = community.Name
		post.VoteNum = voteData[idx]
		data = append(data, post)
	}

	return
}
