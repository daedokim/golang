package routes

import (
	"errors"
	"holdempoker/models"
	"time"
)

const (
	// OsTypeWeb is 웹
	OsTypeWeb = 0
	// OsTypeIos is IOS
	OsTypeIos = 1
	// OsTypeAndroid is 안드로이드
	OsTypeAndroid = 2

	// NewUserCoin 새유저 코인
	NewUserCoin = 100000000
)

// Login is 로그인을 한다.
func Login(data map[string]interface{}) (interface{}, error) {
	var returnVal interface{}
	var err error

	ostype := data["ostype"].(float64)

	switch ostype {
	case OsTypeWeb:
		returnVal, err = LoginWeb(data)
	case OsTypeIos:
		fallthrough
	case OsTypeAndroid:
		returnVal, err = LoginMobile(data)
	}

	return returnVal, err
}

// LoginWeb is 웹용 로그인
func LoginWeb(data map[string]interface{}) (interface{}, error) {
	var returnVal interface{}

	//id := data["id"]
	//passwd := data["passwd"].(string)

	// if id == "" || passwd == "" {

	// }
	return returnVal, nil
}

// LoginMobile is 모바일용 로그인
func LoginMobile(data map[string]interface{}) (interface{}, error) {
	var returnVal interface{}

	if data["uid"] != nil {
		var auth models.UserAuth
		var user models.User
		uid := data["uid"].(string)
		ostype := int(data["ostype"].(float64))
		notFound := db.Where(&models.UserAuth{UID: uid, OsType: ostype}).First(&auth).RecordNotFound()

		if notFound == true {
			session := db.Begin()
			user = models.User{Coin: 1000000, UserID: uid, LoginDate: time.Now(), WriteDate: time.Now()}

			if err := session.Create(&user).Error; err != nil {
				session.Rollback()
				return nil, err
			}

			if err := session.Create(&models.UserAuth{UserIndex: user.UserIndex, OsType: ostype, UID: uid}).Error; err != nil {
				session.Rollback()
				return nil, err
			}
			session.Commit()
		} else {
			if err := db.First(&user, "user_id = ?", uid).Error; err != nil {
				return nil, errors.New("Error")
			}
		}

		returnMap := make(map[string]interface{})
		returnMap["User"] = user

		returnVal = returnMap
	}
	return returnVal, nil
}
