package netstat

import (
	"fmt"
	"strconv"
	"bytes"
	"os/exec"
	"net/http"
	"github.com/julienschmidt/httprouter"
	mylog "github.com/maxwell92/gokits/log"
	"time"
)

var logger = mylog.Log

const (
	GREP_RESULT_NIL     = "0"
	RUN_NETSTAT_FAILED  = "1"
	RUN_NETSTAT_SUCCESS = "2"
)

var HeaderInfo string = `Active Internet connections (w/o servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
Active UNIX domain sockets (w/o servers)
Proto RefCnt Flags       Type       State         I-Node   Path
`

func Netstat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("---------enter Netstat")
	grep := ""
	port := params.ByName("port")
	established := params.ByName("established")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		logger.Errorf("port=%s 参数错误", port)
		w.Write([]byte("port参数错误"))
		return
	}

	var buf bytes.Buffer

	//约定port=-1表示用户不指定端口
	if portInt != -1 || established == "1" {
		buf.WriteString(HeaderInfo)
		if portInt != -1 {
			grep = " | grep " + port
		}
		if established == "1" {
			grep = grep + " | grep ESTABLISHED"
		}
	}
	fmt.Println("------port", grep)

	//mac对netstat的支持与Linux不完全一致
	//mac不支持-p参数，如果指定run启动报错且可能使程序陷入漫长等待
	//todo 记得加上p
	cmd := exec.Command("/bin/bash", "-c", "netstat -ant"+grep)
	cmd.Stdout = &buf
	//用于测试netstat命令是否成功启动
	s := make([]string, 1)
	go run(cmd, &s)

	time.Sleep(50 * time.Millisecond)
	if s[0] == RUN_NETSTAT_FAILED {
		w.Write([]byte("命令执行失败，请重试"))
		return
	} else if s[0] == GREP_RESULT_NIL {
		w.Write([]byte("没有匹配端口的数据,不带端口或者换个端口尝试吧"))
		return
	}

	timer := time.NewTimer(10 * time.Second)
	timeout := timer.C
	for {
		select {
		case <-timeout:
			{
				w.Write([]byte("命令执行超时，请重试"))
				return
			}
		default:
			{
				if buf.String() != "" {
					//等待数据写入
					time.Sleep(1 * time.Second)
					w.Write([]byte(buf.String()))
					fmt.Println(buf.String())
					return
				}
			}
		}
	}
}

func run(cmd *exec.Cmd, s *[]string) {
	err := cmd.Run()
	if err != nil {
		if isGrepResultNil(s) {
			return
		}
		logger.Errorf("netstat failed. err = %s\n", err)
		(*s)[0] = RUN_NETSTAT_FAILED
		return
	}
	(*s)[0] = RUN_NETSTAT_SUCCESS
}

//当grep xx并没有匹配的数据时，cmd.Run()执行会出错
func isGrepResultNil(s *[]string) bool {
	//todo 记得加上参数P
	err := exec.Command("/bin/bash", "-c", "netstat -ant").Run()
	if err != nil {
		return false
	}
	(*s)[0] = GREP_RESULT_NIL
	return true
}
