package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	//启动一个httpsEREVER 内注入对应的
	// http.Handler
	// gin.HANLER
	// behi./roughen
	// echo
	// 自我诶你对server的宝座
	n := negroni.Classic()
	n.UseHandler(mux)
	hostString := fmt.Sprintf(":%s", port)
	n.Run(hostString)

}
func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from Go!")
}

//简单的使用的 http.NewMux路由进行处理
//使用negorRun组我诶 nginx
// API-RPOXY
// router
// formater

//适配器模式
// gin echo beego iris
// handler
// ngix
// oep
// oweklr
//
