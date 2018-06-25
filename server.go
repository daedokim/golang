package main

import (
	"encoding/json"
	. "holdempoker/models"
	. "holdempoker/routes"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
)

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

	controller = Controller{}
	controller.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.GET("/ws", handle)
	e.Logger.Fatal(e.Start(":1323"))
}
