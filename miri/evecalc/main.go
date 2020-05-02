package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Count struct {
	Id uint
	N  uint
	F  uint
	S  uint
	T  uint
}

func gormcone() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "2rankino5110ru17ta"
	PROTOCOL := ""
	DBNAME := "howmany"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		//panic(err.Error())
		fmt.Println("error!")
	}
	return db
}

func updb(c, i string) {
	db := gormcone()
	defer db.Close()
	ch := make(chan int, 2)
  go func() {
		var cot1 Count
		cot1.Id = 1
		db.Find(&cot1)
		cot1.N += 1
		switch c {
		case "1":
			cot1.F += 1
			cot1.S += 0
			cot1.T += 0
		case "2":
			cot1.F += 0
			cot1.S += 1
			cot1.T += 0
		case "3":
			cot1.F += 0
			cot1.S += 1
			cot1.T += 0
		}
		db.Save(&cot1)
		ch <- 0
	}()

	go func() {
		var cot2 Count
		cot2.Id = 2
		db.Find(&cot2)
		cot2.N += 1
		switch i {
		case "1":
			cot2.F += 1
			cot2.S += 0
			cot2.T += 0
		case "2":
			cot2.F += 0
			cot2.S += 1
			cot2.T += 0
		case "3":
			cot2.F += 0
			cot2.S += 0
			cot2.T += 1
		}
		db.Save(&cot2)
		ch <- 1
	}()
	<- ch
	<- ch
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")
	router.Static("/css", "./css")
	router.Static("/js", "./js")
	router.Static("/img", "./img")
	now := time.Now()
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"hour": now.Hour(), "minute": now.Minute(), "second": now.Second()})
	})

	router.POST("/", func(c *gin.Context) {
		tetx := c.PostForm("test")
		updb(tetx, tetx)
		ci1 := c.PostForm("c")
		ci2 := c.PostForm("i")
		fmt.Printf("c:%T,%v i:%T,%v\n", ci1, ci1, ci2, ci2)
		c.HTML(200, "index.html", gin.H{"tetx": tetx})
	})

	router.POST("/test", func(c *gin.Context) {
		tes := c.PostForm("test")
		fmt.Println(tes)
	})

	router.Run(":8080")
}
