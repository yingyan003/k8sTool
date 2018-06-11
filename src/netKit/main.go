package main

import (
	"net/http"
	"netKit/ping"
	mylog "github.com/maxwell92/gokits/log"
	"github.com/julienschmidt/httprouter"
	"netKit/netstat"
)

var logger = mylog.Log

func main() {
	router:=httprouter.New()
	router.GET("/ping/:dstIP/:count",ping.Ping)
	router.GET("/netstat/:port/:established",netstat.Netstat)

	logger.Infoln("troubleshooting listen at port: 8090")
	http.ListenAndServe(":8090", router)
}
