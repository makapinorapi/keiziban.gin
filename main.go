package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Userresponse struct {
	Title   string
	Content string
}

func main() {

	dsn := "root:nora@tcp(127.0.0.1:3306)/keiziban?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// TODO
	}

	r := gin.Default()

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

	r.GET("/hello", func(c *gin.Context) {
		comments := []Comment{}

		db.Find(&comments) //sqlからuserテーブルの情報を見つけ出し、gormのuserに値を渡している
		//fmt.Println(len(user))

		//fmt.Println(len(user))
		//fmt.Println("AAAAAAAAA", len(user))
		for i := 0; i < len(comments); i++ {
			// ユーザーのTimeとか全部の情報が一通り書き出されてる
		}
		res := make([]Userresponse, len(comments))
		//fmt.Println(len(user))
		//fmt.Println(len(user))
		//fmt.Println(user)
		for i := 0; i < len(res); i++ {
			res[i].Title = comments[i].Title
			res[i].Content = comments[i].Content
		}
		//fmt.Println(len(res))
		//fmt.Println(res)

		c.JSON(200, gin.H{ //↓"message":"hello"みたいな

			"title": res, //BMIみたいに一つだけ取り出す

		})
	})

	db.AutoMigrate(&Comment{})

	comments := Comment{Title: "はるはるはるひ", Content: "ぎゃんぎゃん"}
	//user2 := []User{} //全取得はデータを格納する[]変数の定義
	db.Create(&comments)

	//db.Find(&user2)

	r.POST("/post", func(c *gin.Context) {

		comme := Comment{} //まlとか書いていたやつを消して、代入する
		c.ShouldBindJSON(&comme)

		/*if err != nil {
			// TODO
		}
		*/
		db.Create(&comme)
		fmt.Println(comme, "eeeeeeeeeeeeeeeeeeeeeeeeeeee")

		for i := 0; i < len(comme); i++ {
			// ユーザーのTimeとか全部の情報が一通り書き出されてる
		}
		res := make([]Userresponse, len(comme))
		//fmt.Println(len(user))
		//fmt.Println(len(user))
		//fmt.Println(user)
		for i := 0; i < len(res); i++ {
			res[i].Title = comme[i].Title
			res[i].Content = comme[i].Content*/
		}
		//fmt.Println(len(res))
		//fmt.Println(res)

		c.JSON(200, gin.H{ //↓"message":"hello"みたいな

			"title": res, //BMIみたいに一つだけ取り出す

		})

	})

	r.Run() // listen and serve on 0.0.0.0:8082 (for windows "localhost:8080")
}

type Comment struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Request struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
