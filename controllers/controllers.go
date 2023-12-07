package controllers

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/heavenland/hl-game-loginhandler/dtos"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Basic Login Functions with Gorilla Websocket

func stringToByte(input string) []byte {
	return []byte(input)
}

var upgrader = websocket.Upgrader{}

func TestController() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lets upgrade the connection to Websocket
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Problem Occured during converting connection to Websocket!")
			return
		}
		// Defer waits till other functions returns
		// It will not close connection until err is returned or Socket handling is done
		defer ws.Close()
		// var userLogin dtos.Login
		// if err := c.BindJSON(&userLogin); err != nil {
		// 	c.AbortWithError(http.StatusBadRequest, err)
		// 	return
		// }
		// Keep websocket on
		for {
			//Read Message from client
			// message itself is bidirectional, therefore it can be overwritten both by client and server
			// When client sends it will be readable by ws.ReadMessage
			// When Server sends it will be sent with ws.WriteMessage
			mt, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			//If client message is ping will return pong
			if string(message) == "ping" {
				// message = []byte("pong")
				message = stringToByte("pong")
			}
			//If client message is json will return JSON
			// It will use DTo which is Struct by Marhsalling it
			// More information on https://www.dotnetperls.com/json-go
			if string(message) == "json" {
				exampleLogin := dtos.Login{
					Email:    "TestUser@test.cz",
					Password: "123456CommonPassword",
				}
				// json Marshall will return []byte and error
				message, err = json.Marshal(exampleLogin)
				if err != nil {
					fmt.Println(err)
					break
				}
			}
			//Response message to client
			// If Message is not ping then it will write message back to user
			err = ws.WriteMessage(mt, message)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}

func TestMarhsallUnMarshall() gin.HandlerFunc {
	return func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Problem Occured during converting connection to Websocket!")
			return
		}
		defer ws.Close()
		for {
			// Read the message
			mt, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			// Declare input as user login arrays(JSON)
			var userLogin dtos.Login
			// Unmarshall it to userLogin
			json.Unmarshal(message, &userLogin)
			fmt.Println("JSON : ", userLogin)
			fmt.Println("JSON Message: ", message)
			// Edit the message and send back to User
			userLogin.Email = "LetsCorrectThat"
			userLogin.Password = "AlsoPassword"
			// Marshall It back
			message, err = json.Marshal(userLogin)
			if err != nil {
				fmt.Println(err)
				break
			}
			err = ws.WriteMessage(mt, message)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}
