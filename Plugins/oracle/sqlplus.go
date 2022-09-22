package oracle

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"fmt"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
	"os/exec"
)

var ScanPort = "1521"

func SqlplusAuth(host string, port string, user string, pass string, db string) bool {
	cmd := exec.Command("sqlplus", "-s", user+"/"+pass+"@"+host+":"+port+"/"+db)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
		return false
	}
	if string(output) == "" {
		return true
	}
	return false
}

func SqlPlusScan2(ScanType string, Target string) {
Loop:
	for _, user := range dic.UserDic() {
		for _, pass := range dic.PassDic() {
			fmt.Println("Checking: " + Target + " " + user + " " + pass)
			res := SqlplusAuth(Target, ScanPort, user, pass, "orcl")
			if res {
				logger.PrintIsok2(ScanType, Target, ScanPort, user, pass)
				common.ResultsMap.Lock()
				common.ResultsMap.Credentials = append(common.ResultsMap.Credentials, common.Credential{Url: Target, Port: ScanPort, UserName: user, Password: pass, Group: "SQL Plus"})
				common.ResultsMap.Unlock()
				break Loop
			}
		}
	}
}
