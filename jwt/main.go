package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/web"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const secret = "fa!$!@g&(*67"

type MyClaim struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.RequestURI)
		signature := c.Request.Header.Get("Auth-Token")
		if signature == "" {
			c.JSON(401, gin.H{
				"msg": "未登录",
			})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(signature, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			fmt.Println(err)
			c.JSON(401, gin.H{
				"msg": "验签失败",
			})
			c.Abort()
			return
		}
		fmt.Println(token.Claims.(*MyClaim).UserName)
		c.Set("userInfo", token.Claims)
		c.Next()
	}
}

func login(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(400, gin.H{
			"msg": "用户名有误",
		})
		return
	}
	claim := MyClaim{
		UserName: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 600,
		},
	}
	sensitiveToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signature, _ := sensitiveToken.SignedString([]byte(secret))
	c.JSON(200, gin.H{
		"token": signature,
	})
}

func info(c *gin.Context) {
	userInfo, isExist := c.Get("userInfo")
	if !isExist {
		c.JSON(201, gin.H{
			"msg": "登录信息有误",
		})
		return
	}
	claim := userInfo.(*MyClaim)
	c.JSON(200, gin.H{
		"name": claim.UserName,
	})
}

func main() {

	route := gin.Default()

	group1 := route.Group("/group/")
	{
		group1.GET("test1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "test1",
			})
		})
		group1.GET("test2", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "test2",
			})
		})
	}

	route.POST("/user/login", login) //login 在Auth之前 所以不会进行登录认证

	route.Use(Auth()) //中间件按顺序执行  该代码之前的路由不会执行该代码

	route.GET("/user/info", info)

	srv := web.NewService(
		web.Name("jwt-sever"),
		web.Address("0.0.0.0:8889"),
		web.Handler(route),
	)

	_ = srv.Run()
}
