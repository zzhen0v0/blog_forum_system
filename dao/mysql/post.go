package mysql

import (
	"database/sql"
	"web_app/models"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (
post_id,title,content,author_id,community_id)
values(?,?,?,?,?) `
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorId, p.CommunityID)
	return
}

func GetPostByID(id int64) (data *models.ApiPostDetail, err error) {
	data = new(models.ApiPostDetail)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?`
	err = db.Get(data, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("querry failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

func GetPostList(page, size int64) (posts []*models.ApiPostDetail, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	limit ?,?
	`
	posts = make([]*models.ApiPostDetail, 0, size)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	if err != nil {
		err = ErrorQueryFailed
		return
	}
	return
}

func GetPostByIDs(ids []string) (posts []*models.ApiPostDetail, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)`
	// 动态填充id
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&posts, query, args...)
	return
}
