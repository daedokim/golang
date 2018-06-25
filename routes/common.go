package routes

// . "holdempoker/models"

const (
	// OsTypeWeb is 웹
	OsTypeWeb = 0
	// OsTypeIos is IOS
	OsTypeIos = 1
	// OsTypeAndroid is 안드로이드
	OsTypeAndroid = 2
)

// Login is 로그인을 한다.
func Login(data map[string]interface{}) interface{} {
	var returnVal interface{}
	osType := data["ostype"].(int)

	switch osType {
	case OsTypeWeb:
		returnVal = LoginWeb(data)
	case OsTypeIos:
		fallthrough
	case OsTypeAndroid:
		returnVal = LoginMobile(data)
	}

	return returnVal
}

// LoginWeb is 웹용 로그인
func LoginWeb(data map[string]interface{}) interface{} {
	var returnVal interface{}

	id := data["id"].(string)
	passwd := data["passwd"].(string)

	if id == "" || passwd == "" {

	}
	return returnVal
}

// LoginMobile is 모바일용 로그인
func LoginMobile(data map[string]interface{}) interface{} {
	var returnVal interface{}

	uid := data["uid"].(string)

	if uid == "" {

	}

	return returnVal
}
