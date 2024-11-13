package main

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"testbeegodemo/models"
	_ "testbeegodemo/routers"
	"time"
)

func main() {
	beego.AddFuncMap("ShowNextPage", getNextPage)
	beego.AddFuncMap("ShowLastPage", getLastPage)
	models.Init()
	beego.Run()

}

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
}

func getNextPage(pageindex int) int {
	pageindex++
	return pageindex
}
func getLastPage(pageindex int) int {
	if pageindex--; pageindex < 0 {
		pageindex = 0
	}
	return pageindex
}
