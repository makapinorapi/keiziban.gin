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
		AllowOrigins: []string{"http://localhost:8080/hello"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers", "Content-Length",
			"Accept-Encoding", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	fmt.Println("aaa") //↓ハローが呼び出されたときの関数
	r.GET("/hello", func(c *gin.Context) {
		user := []User{}

		db.Find(&user) //sqlからuserテーブルの情報を見つけ出し、gormのuserに値を渡している
		fmt.Println(len(user))
		//db.Take(&user)
		//users[0].Title
		//users[0].Content
		//db.Select([]string{"title", "content"}).Find(&user)
		//fmt.Println("AAAAAAAAAAAA", user[9])
		fmt.Println(len(user))
		fmt.Println("AAAAAAAAA", len(user))
		for i := 0; i < len(user); i++ {
			// ユーザーのTimeとか全部の情報が一通り書き出されてる
		}
		res := make([]Userresponse, len(user))
		fmt.Println(len(user))
		fmt.Println(len(user))
		//fmt.Println(user)
		for i := 0; i < len(res); i++ {
			res[i].Title = user[i].Title
			res[i].Content = user[i].Content
		}
		fmt.Println(len(res))
		fmt.Println(res)
		//fmt.Println(reflect.TypeOf(res))
		//fmt.Println(reflect.TypeOf(user))
		c.JSON(200, gin.H{ //↓"message":"hello"みたいな

			"title": res, //BMIみたいに一つだけ取り出す

		})
	})

	db.AutoMigrate(&User{})
	//db.Migrator().CreateTable(&User{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})

	user1 := User{Title: "はるはるはるひ", Content: "ぎゃんぎゃん"}
	user2 := []User{} //全取得はデータを格納する[]変数の定義
	db.Create(&user1)

	//db.Take(&user2)
	//fmt.Println(user2)

	//db.Find(&user2)
	//fmt.Println(user2)

	db.Find(&user2)
	//fmt.Println("CCCCCCCCCCCCCCC", user2)

	//human:= map[string]float64{"weight":46.0,"height":1.56}
	//var total float64
	//total = human["weight"]/(human["height"] * human["height"])
	//fmt.Println(total)

	/*r.POST("/post", func(c *gin.Context) {
		s := c.PostForm("str")
		n := c.PostForm("num")
		b := c.PostForm("bool")
		l := c.DefaultPostForm("limit", "10")
		message := fmt.Sprintf("s: %v, n: %v, b: %v, l: %v", s, n, b, l)
		c.String(http.StatusOK, message)
	})*/

	r.POST("/post", func(c *gin.Context) {
		//comment := Comment{}
		//db.AutoMigrate(&Comment{}) //テーブル作成
		//db.Migrator().CreateTable(&Comment{})
		//come := Comment{Title: "まl", Content: "あ"}
		//fmt.Println("wwwwwwwww", come)
		//db.Create(&come) // sqlに反映する

		//fmt.Println("ffffffffffffff", comment)
		//db.Find(&comment)//データを見つける

		request := Comment{} //まlとか書いていたやつを消して、代入する
		c.ShouldBindJSON(&request)

		if err != nil {
			// TODO
		}
		fmt.Println(request)
		db.Create(&request)

	})

	/*r.POST("/post2", func(c *gin.Context) {
	request := Request{Title: "aaaa", Content: "bbbbbb"}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title": "aaaaa",
	})
	*/
	/*if e != nil {
		// TODO
		db.Create(&request)
		db.Select("Name", "Age", "CreatedAt").Create(&request)
	}*/
	//fmt.Println(request)
	///*if e != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	//} else {
	//	c.String(http.StatusOK, "ok")
	//}*/

	//})
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
