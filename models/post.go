package models

import (
	"time"
)

type Post struct {
	PostID      int64     `json:"post_id" db:"post_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}

type ApiPostDetail struct {
	*Post
	VoteNum       int64  `json:"vote_num"`
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
}
