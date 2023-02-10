package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type UserResponse struct {
	ID        uint
	User      string
	Content   string
	GoodCount int
}

func main() {

	dsn := "root:nora@tcp(127.0.0.1:3306)/keiziban?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// TODO
	}

	r := gin.Default()
	db.AutoMigrate(&Comment{})

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"}, //どっからアクセスを許可するか　vue側からginくださいと言ってる
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers", "Content-Length",
			"Accept-Encoding", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/comments", func(c *gin.Context) {
		comments := []Comment{}
		db.Find(&comments) //sqlからuserテーブルの情報を見つけ出し、gormのuserに値を渡している
		res := make([]UserResponse, len(comments))
		for i := 0; i < len(res); i++ {
			res[i].User = comments[i].User
			res[i].Content = comments[i].Content
		}
		c.JSON(200, res)
	})

	r.POST("/comments", func(c *gin.Context) {

		comme := Comment{} //まlとか書いていたやつを消して、代入する
		c.ShouldBindJSON(&comme)
		db.Create(&comme)
		c.JSON(200, gin.H{ //↓"message":"hello"みたいな
			"result": "ok",
		})

	})

	r.POST("/goodCount", func(c *gin.Context) {
		comments := Comment{}
		//fmt.Println(reflect.TypeOf("comments")) // string
		//db.Where(&Comment{ID: 188}).First(&comments)
		//db.Select("id").Find(&comments)
		c.ShouldBindJSON(&comments)
		//fmt.Printf("%T\n", c.ShouldBindJSON, "rrrrrrrr")
		db.Take(&comments) //値を取る
		fmt.Println(comments, "aaaaaaaaa")
		comments.GoodCount++
		//db.Create(&comments)
		fmt.Println(comments, "wwwwwwwwwwww")
		db.Save(&comments) //保存
		//fmt.Println(comments, "jjjjjjjjjjjj")
		c.JSON(200, gin.H{ //↓"message":"hello"みたいな
			"result": "ok",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8082 (for windows "localhost:8080")
}

type Comment struct {
	ID        uint      `json:"id"`
	User      string    `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	GoodCount int       `json:"goodCount"`
}
