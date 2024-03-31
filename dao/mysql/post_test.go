package mysql

import (
	"testing"
	"web_app/models"
	"web_app/settings"
)

func init() {
	dbCfg := settings.MysqlConfig{
		User:         "root",
		Host:         "127.0.0.1",
		Password:     "1231",
		Port:         3306,
		DB:           "bluebell",
		MaxOpenConns: 200,
		MaxIdleConns: 50,
	}

	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		PostID:      123,
		Title:       "test",
		Content:     "just a test",
		AuthorId:    1,
		CommunityID: 12,
	}

	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed ,err: %v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
