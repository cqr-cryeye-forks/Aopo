package oracle

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"database/sql"
	_ "github.com/godror/godror"
)

func OracleAuth(host string, port string, user string, pass string) (result bool) {
	//defer common.Oraclewg.Done()
	db, err := sql.Open("godror", user+"/"+pass+"@"+host+":"+port+"/orcl")
	//if err != nil {
	//panic(err)
	//return false
	//}
	if err == nil {
		if db.Ping() == nil {
			db.Close()
			result = true
			return result
		}
	}
	result = false
	if result != false {
		gologger.Infof("Oracle Found: " + host + ":" + port + " " + user + " " + pass)
		common.ResultsMap.Lock()
		common.ResultsMap.Credentials = append(common.ResultsMap.Credentials, common.Credential{Url: host, Port: port, UserName: user, Password: pass, Group: "Oracle"})
		common.ResultsMap.Unlock()
	}
	common.Oraclewg.Done()
	return result
}
