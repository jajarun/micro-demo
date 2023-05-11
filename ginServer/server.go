package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var userCons = make(map[string]map[*websocket.Conn]bool)

func handleMessage(msg string, userId string) {
	for conn, _ := range userCons[userId] {
		err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("收到%s的消息%s", userId, msg)))
		if err != nil {
			log.Println("write msg err:", err)
		}
	}
}

func websocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	userId, ok := c.GetQuery("user_id")
	if !ok {
		log.Println("user_id 为空 关闭连接")
		_ = conn.Close()
		return
	}
	if userCons[userId] == nil {
		userCons[userId] = make(map[*websocket.Conn]bool)
	}
	userCons[userId][conn] = true
	_ = conn.WriteMessage(websocket.TextMessage, []byte(userId+"连接成功"))
	log.Println(userId + "连接数" + strconv.Itoa(len(userCons[userId])))
	conn.SetCloseHandler(func(code int, text string) error {
		delete(userCons[userId], conn)
		log.Println(userId + "关闭连接")
		log.Println(userId + "连接数" + strconv.Itoa(len(userCons[userId])))
		return nil
	})
	for {
		messageType, msg, err := conn.ReadMessage()
		log.Println("msg type:", strconv.Itoa(messageType))
		if err != nil {
			_ = conn.Close()
			log.Println("read msg err:", err)
			return
		}
		log.Println("read msg:" + string(msg))
		go handleMessage(string(msg), userId)
	}
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		logger.Info("Hello, world!")
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	router.GET("/ws", websocketHandler)

	service := web.NewService(
		web.Name("gin-server"),
		web.Address("127.0.0.1:80"),
		web.Handler(router),
		//web.Registry(consul.NewRegistry()),
	)

	_ = service.Run()
}
