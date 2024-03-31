package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignup) (err error) {
	//1 check if user exist

	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	//generate userId
	userID := snowflake.GenID()

	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	//3 insert
	return mysql.InsertUser(&user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	//generate jwt token
	return jwt.GenToken(user.UserID, user.Username)
}
