package main

import (
	"encoding/json"
	"fmt"

	"os"
	"time"

	. "holdempoker/models"
	. "holdempoker/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoLog "github.com/labstack/gommon/log"

	log "github.com/neko-neko/echo-logrus/log"
	viper "github.com/spf13/viper"
	"golang.org/x/net/websocket"
)

//Config Start
type Config struct {
	Debug    bool
	Database struct {
		Driver     string
		Connection string
	}
	Host string
	Port string
}

//Conf Start
var Conf Config

//LoadConfig Start
func LoadConfig() {
	env := os.Getenv("GOENV")
	var confile string
	if env == "" {
		confile = "config.dev.yml"
	} else if env == "prod" {
		confile = "config.yml"
	}
	file, err := os.Open(confile)
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	defer file.Close()
	viper.MergeConfig(file)
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	fmt.Printf("--- Load config from %s ---\n", confile)
}

//DB Start
var DB *gorm.DB

//OpenDatabase start
func OpenDatabase() {
	var err error
	DB, err = gorm.Open(Conf.Database.Driver, Conf.Database.Connection)

	if err != nil {
		panic(err)
	}
	fmt.Println("--- Open database ---")
	DB.LogMode(Conf.Debug)
	DB.SingularTable(true)
}

//DB End

var controller Controller

func handle(c echo.Context) error {

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		Receive(c, ws)
		Push(c, ws)

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

// Push is 데이터를 보낸다.
func Push(c echo.Context, ws *websocket.Conn) {
	for {
		time.Sleep(1000 * time.Millisecond)
	}
}

// Receive is 데이터를 받아온다.
func Receive(c echo.Context, ws *websocket.Conn) {
	for {
		var data PacketData
		err := websocket.JSON.Receive(ws, &data)

		if err != nil {
			c.Logger().Error(err)
			return
		}

		if data.PacketNum > 0 {
			returnVal := controller.Handle(data.PacketNum, data.PacketData.(map[string]interface{}))

			if returnVal != nil {
				jsonString, err := json.Marshal(returnVal)
				if err != nil {
					c.Logger().Error(err)
				}
				err = websocket.JSON.Send(ws, string(jsonString))
				if err != nil {
					c.Logger().Error(err)
				}
			}
		}
	}
}

func main() {

	LoadConfig()
	OpenDatabase()

	controller = Controller{}
	controller.Init()

	e := echo.New()
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.DEBUG)
	e.Logger = log.Logger()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.GET("/ws", handle)
	e.Logger.Fatal(e.Start(":1323"))
}
