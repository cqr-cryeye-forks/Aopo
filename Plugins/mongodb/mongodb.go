package mongodb

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"gopkg.in/mgo.v2"
	"time"
)

func MongoAuth(ip string, port string, username string, password string) (result bool, err error) {
	session, err := mgo.DialWithTimeout("mongodb://"+username+":"+password+"@"+ip+":"+port+"/"+"admin", time.Second*3)
	if err == nil && session.Ping() == nil {
		session.Close()
		if err == nil && session.Run("serverStatus", nil) == nil {
			result = true
		}
	}
	if result != false {
		gologger.Infof("MongoDB Found: " + ip + ":" + port + " " + username + " " + password)
		common.ResultsMap.Lock()
		common.ResultsMap.Credentials = append(common.ResultsMap.Credentials, common.Credential{Url: ip, Port: port, UserName: username, Password: password, Group: "MongoDB"})
		common.ResultsMap.Unlock()
	}
	common.Mongowg.Done()
	return result, err
}
