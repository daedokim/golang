package main

import (
	"encoding/json"

	"os"
	"time"

	"holdempoker/config"
	database "holdempoker/db"
	"holdempoker/models"
	"holdempoker/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoLog "github.com/labstack/gommon/log"

	log "github.com/neko-neko/echo-logrus/log"
	"golang.org/x/net/websocket"
)

//Conf Start
var Conf config.Config
var controller routes.Controller

//DB End

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
		var data models.PacketData
		err := websocket.JSON.Receive(ws, &data)

		if err != nil {
			c.Logger().Error(err)
			return
		}

		if data.PacketNum > 0 {
			returnVal := controller.Handle(data.PacketNum, data.Data.(map[string]interface{}))

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

func main() {

	Conf = config.Config{}
	Conf.LoadConfig()

	database.GetDBInstance().OpenDatabase(&Conf)

	controller = routes.Controller{}
	controller.Init()

	e := echo.New()
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.DEBUG)

	e.Logger = log.Logger()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.GET("/ws", handle)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: Conf.Cors.Hosts,
		AllowMethods: []string{echo.GET},
	}))
	e.Logger.Fatal(e.Start(Conf.Port))
}
