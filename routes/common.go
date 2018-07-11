package routes

import (
	"errors"
	"holdempoker/models"
	"time"

	"github.com/jinzhu/gorm"
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

// Login 로그인을 한다.
func Login(data map[string]interface{}) (interface{}, error) {
	var returnVal interface{}
	var err error

	if data["osType"] == nil {
		return nil, errors.New("dd")
	}

	ostype := data["osType"].(float64)

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

	if data["userId"] != nil {
		var auth models.UserAuth
		var user models.User
		var notFound bool

		guestmode := data["guestMode"].(bool)
		userid := data["userId"].(string)
		ostype := int(data["osType"].(float64))

		if guestmode == true {
			notFound = db.Where(&models.UserAuth{UID: userid, OsType: ostype}).First(&auth).RecordNotFound()
		} else {
			notFound = db.Where(&models.User{UserID: userid}).First(&user).RecordNotFound()
		}

		var session *gorm.DB

		if notFound == true {
			session = db.Begin()
			user = models.User{Coin: 1000000, UserID: userid, LoginDate: time.Now(), WriteDate: time.Now(), NickName: "NoName"}

			if err := session.Create(&user).Error; err != nil {
				session.Rollback()
				return nil, err
			}

			if guestmode == true {
				if err := session.Create(&models.UserAuth{UserIndex: user.UserIndex, OsType: ostype, UID: userid}).Error; err != nil {
					session.Rollback()
					return nil, err
				}
			}
			session.Commit()
		} else {
			if err := db.First(&user, "user_id = ?", userid).Error; err != nil {
				return nil, errors.New("Error")
			}

			session = db.Begin()
			if err := session.Model(&user).Update("logn_date", time.Now()).Error; err != nil {
				session.Rollback()
				return nil, err
			}
			session.Commit()
		}

		returnMap := make(map[string]interface{})
		returnMap["User"] = user

		returnVal = returnMap
	}
	return returnVal, nil
}
