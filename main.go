package main

import (
	"./server"
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strconv"
)

func main() {

	port := flag.Int("p", 80, "Port")
	flag.Parse()
	//echo instance
	e := echo.New()
	e.Logger.SetLevel(log.OFF) //로그 레벨 설정

	// 미들웨어 등록
	e.Use(middleware.Logger()) //http 리퀘스트 기록
	e.Use(middleware.Recover()) //HTTPError 제어, 처리, 복구

	defer func() {
		if err := e.Close(); err != nil {
			panic(err)
		}
	}()

	e.GET("/set-cert-info", server.MainGetTest)




	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*port)))
}