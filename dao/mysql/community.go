package mysql

import (
	"database/sql"
	"fmt"
	"web_app/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id,community_name from community`
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no db")
		err = nil
	}
	fmt.Println(communityList)
	return
}

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select 
			community_id,community_name,introduction,create_time 
			from community 
			where community_id=?`
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}

func GetCommunityByID(id int64) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := `select community_id,community_name from community where community_id=?`
	err = db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorUserNotExist
		return
	}
	return
}
