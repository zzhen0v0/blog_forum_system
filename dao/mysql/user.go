package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"web_app/models"
)

const secret = "zzhen0v0"

func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {

	user.Password = encryptPassword(user.Password)

	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select user_id,username,password from user where username=?"
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if user.Password != encryptPassword(oPassword) {
		return ErrorPasswordInvalid
	}
	return
}

func GetUserByID(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username,password from user where user_id=?`
	err = db.Get(user, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorUserNotExist
		return
	}
	return
}
