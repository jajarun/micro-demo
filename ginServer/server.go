package main

import (
	"encoding/json"
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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
	res := struct {
		ErrCode int
		UserId  int
		Message string
	}{0, 1, "链接成功"}
	msg, _ := json.Marshal(res)
	_ = conn.WriteMessage(websocket.TextMessage, msg)
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access-Control-Allow-Credentials=true和Access-Control-Allow-Origin="*"有冲突
		//故Access-Control-Allow-Origin需要指定具体得跨域origin
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081/")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "content-type")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		//c.Header("Access-Control-Expose-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(Cors())
	router.Use(Auth())

	router.GET("/hello", func(c *gin.Context) {
		logger.Info("Hello, world!")
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	router.GET("/testApi", func(c *gin.Context) {
		logger.Info("Hello, world!")
		c.JSON(200, gin.H{
			"data": "Hello, world!",
		})
	})

	router.GET("/ws", websocketHandler)

	service := web.NewService(
		web.Name("gin-server"),
		web.Address("0.0.0.0:8888"),
		web.Handler(router),
		//web.Registry(consul.NewRegistry()),
	)

	_ = service.Run()
}
