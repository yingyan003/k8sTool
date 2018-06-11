package ping

import (
	"net/http"
	"os/exec"
	"github.com/julienschmidt/httprouter"
	mylog "github.com/maxwell92/gokits/log"
	"time"
	"strconv"
	"bytes"
	"github.com/astaxie/beego/logs"
)

var logger = mylog.Log

func Ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	logs.Info("---------enter Ping")
	dstIP := params.ByName("dstIP")
	count := params.ByName("count")

	countInt, _ := strconv.ParseInt(count, 10, 64)
	if countInt < 1 || countInt > 10 {
		w.Write([]byte("请求次数只能在1-10之间"))
		return
	}

	var buf bytes.Buffer
	cmd := exec.Command("ping", "-c", count, dstIP)
	cmd.Stdout = &buf
	go run(cmd, dstIP)

	time.Sleep(time.Second * time.Duration(countInt))
	if buf.String() != "" {
		w.Write([]byte(buf.String()))
	} else {
		w.Write([]byte("ping " + dstIP + " failed"))
	}
}

func run(cmd *exec.Cmd, dstIP string) {
	err := cmd.Run()
	if err != nil {
		logger.Errorf("ping %s failed. err=", dstIP, err)
	}
}
