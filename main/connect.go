package main

import (
	sql "demo02/Sql"
	"demo02/router"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//连接数据库
	sql.ConnextSQL()
	//连接redis
	sql.ConnectRedis()
	//开启网页接口
	rt := router.InitRouter()
	//定义服务监听的地址
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 8080)
	server := &http.Server{
		Addr:    addr,
		Handler: rt,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("服务器启动失败")
		os.Exit(-1)
	}
}
