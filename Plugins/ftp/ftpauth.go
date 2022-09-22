package ftp

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func FtpAuth(ip string, port string, user string, pass string) {
	conn, _ := ftp.DialTimeout(ip+":"+port, time.Duration(5)*time.Second)
	err := conn.Login(user, pass)
	fmt.Println(err)
	if err == nil {
		conn.Logout()
		gologger.Infof("FTP Found: " + ip + ":" + port + " " + user + " " + pass)
		common.ResultsMap.Lock()
		common.ResultsMap.Credentials = append(common.ResultsMap.Credentials, common.Credential{Url: ip, Port: port, UserName: user, Password: pass, Group: "FTP"})
		common.ResultsMap.Unlock()
	}

}
