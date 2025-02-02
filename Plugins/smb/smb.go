package smb

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"github.com/stacktitan/smb/smb"
)

//Not Support 2003
func SmbAuth(ip string, port string, username string, password string) (result bool, err error) {
	//defer common.Smbwg.Done()
	result = false

	options := smb.Options{
		Host:        ip,
		Port:        445,
		User:        username,
		Password:    password,
		Domain:      "",
		Workstation: "",
	}

	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			result = true
		}
	}
	if result != false {
		gologger.Infof("SMB Found: " + ip + ":" + port + " " + username + " " + password)
		common.ResultsMap.Lock()
		common.ResultsMap.Credentials = append(common.ResultsMap.Credentials, common.Credential{Url: ip, Port: port, UserName: username, Password: password, Group: "SMB"})
		common.ResultsMap.Unlock()
		//gologger.Infof(aurora.Red("SMB Found: %v:%v %v %v").String(), ip, port, username, password)
	}
	common.Smbwg.Done()
	return result, err
}
