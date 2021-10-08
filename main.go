package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 返回数据的结构
type result struct {
	Openid      string `json:"openid"`
	Session_key string `json:"session_key"`
	Unionid     string `json:"unionid"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

// type room struct {
// 	Prize  string `json:"prize"`
// 	Amount int    `json:"amount"`
// }

type roomInfo struct {
	Openid string `json:"openid"`
	Rooms  string `json:"roomInfo"`
}

type opid struct {
	Openid string `json:"openid"`
}

func main() {
	db, e := sql.Open("mysql", "root:123456@tcp(localhost)/t2?charset=utf8&parseTime=True&loc=Local")
	if e != nil {
		panic(e)
	}
	fmt.Print(db)

	r := gin.Default()

	r.StaticFile("/HYShangWeiShouShuW.ttf", "./HYShangWeiShouShuW.ttf")
	r.Static("/static", "./static")

	r.GET("/login", func(c *gin.Context) {

		// 用户登录code
		code := c.Query("code")
		// 用户小程序id
		id := `wxb27cb3df6158fc0e`
		// 小程序secret
		secret := `001a15eeb9b2e60acfb21ce896e3885f`
		// grant_type const
		grant_type := `authorization_code`
		// 请求路径

		resp, err := http.Get(`https://api.weixin.qq.com/sns/jscode2session?appid=` + id + `&secret=` + secret + `&js_code=` + code + `&grant_type=` + grant_type)

		if err != nil {
			fmt.Print("err")
		}
		//fmt.Print(resp)
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		var res result
		_ = json.Unmarshal(body, &res)

		c.JSON(200, gin.H{
			"unionid": res.Openid,
		})
		fmt.Printf("%#v", res)
		fmt.Print("\n")
	})

	r.POST("/build", func(c *gin.Context) {
		var res roomInfo
		c.BindJSON(&res)

		fmt.Print(c)
		fmt.Print(res.Openid)
		fmt.Print("\n")
		fmt.Print(res.Rooms)
	})

	r.POST("/history", func(c *gin.Context) {
		var openid opid

		c.BindJSON(&openid)

		fmt.Print(openid.Openid)
		// 测试用静态数据
		c.JSON(200, gin.H{
			"sum": 2000,
			"a":   100,
			"b":   121,
			"c":   322,
			"d":   1233,
			"e":   2123,
			"f":   1311,
		})
	})
	r.Run(":8080")
}
