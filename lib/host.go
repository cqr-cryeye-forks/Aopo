package host

import (
	"Aopo/Plugins/ftp"
	"Aopo/Plugins/mongodb"
	"Aopo/Plugins/mssql"
	"Aopo/Plugins/mysql"
	"Aopo/Plugins/oracle"
	"Aopo/Plugins/redis"
	"Aopo/Plugins/smb"
	"Aopo/Plugins/ssh"
	xrayscan "Aopo/Plugins/xray"
	"Aopo/common"
	"Aopo/gologger"
	"fmt"
	"github.com/logrusorgru/aurora"
	"time"
)

var HostList []string

func Host(host string, filename string, nohosts ...string) {
	time.Sleep(time.Duration(1) * time.Second)
	if host != "" {
		gologger.Infof("URL Target: " + host)
	}
	if filename != "" {
		gologger.Infof("File Target: " + filename)
	}
	gologger.Infof("Number of passwords: " + fmt.Sprintln(len(common.Passwords)))
	gologger.Infof("Modules: Host Detection | Port Service Detection | Port Web Service Detection | HTTP Scan | Web Service Vulnerability Detection | Weak Password Detection")
	time.Sleep(time.Duration(1) * time.Second)

	// Host Detection
	gologger.Printf("================================== Host Detection ==================================")
	time.Sleep(time.Duration(1) * time.Second)
	HostList, _ = common.ParseIP(host, filename, nohosts...)
	gologger.Infof("Number of unique hosts: " + fmt.Sprintln(len(HostList)))

	// Port Service Detection
	gologger.Printf("================================== Port Service Detection ==================================")
	time.Sleep(time.Duration(1) * time.Second)
	for i := range HostList {
		gologger.Infof(aurora.Green("Testing Ports： " + HostList[i]).String())
		common.Portwg.Add(1)
		go common.ScanPort(HostList[i])
	}
	common.Portwg.Wait()

	// Port Service Detection
	gologger.Printf("================================== Port Web Service Detection ==================================")
	time.Sleep(time.Duration(1) * time.Second)
	for i := range HostList {
		gologger.Infof(aurora.Green("Testing Web Ports： " + HostList[i]).String())
		common.Portwebwg.Add(1)
		go common.ScanWebPort(HostList[i])
	}
	common.Portwebwg.Wait()

	// HTTP Scan
	gologger.Printf("================================== HTTP Scan ==================================")
	time.Sleep(time.Duration(1) * time.Second)
	common.ResultsMap.Lock()
	for i := range common.ResultsMap.PortWebList {
		for p := range common.ResultsMap.PortWebList[i] {
			gologger.Infof(aurora.Green("Testing： " + i).String())
			common.Httpwg.Add(1)
			go common.HttpScan(i, common.ResultsMap.PortWebList[i][p])
		}
		common.Httpwg.Wait()
	}
	common.ResultsMap.Unlock()

	// WEB Service Vulnerability Detection
	gologger.Printf("================================== WEB Service Vulnerability Detection ==================================")
	time.Sleep(time.Duration(1))
	for i := range common.Xraylist {
		gologger.Infof(aurora.Green("Testing： " + common.Xraylist[i]).String())
		common.Xraywg.Add(1)
		go xrayscan.XrayScan(common.Xraylist[i])
	}
	common.Xraywg.Wait()

	// Weak Password Detection
	gologger.Printf("================================== Weak Password Detection ==================================")
	time.Sleep(time.Duration(1) * time.Second)
	ipNumber := 1
	common.ResultsMap.Lock()
	for i := range common.ResultsMap.HostList {
		common.WeakPass.Add(1)
		gologger.Infof(aurora.Green("Testing： " + common.ResultsMap.HostList[i].Name + ": " + fmt.Sprint(common.ResultsMap.HostList[i].Port) + " (" + common.ResultsMap.HostList[i].Type + ")").String())
		common.Statuses = append(common.Statuses, false)
		go Auth(common.ResultsMap.HostList[i], ipNumber)
		ipNumber += 1
	}
	common.WeakPass.Wait()
	common.ResultsMap.Unlock()
}
func Auth(host common.HostInfo, sum int) {
	defer common.WeakPass.Done()

	if host.Port == common.PORTList["ssh"] {
		gologger.Infof("Testing SSH")
		SshCheck(fmt.Sprint(host.Name+":"+fmt.Sprint(host.Port)), sum)

	} else if host.Port == common.PORTList["mysql"] {
		gologger.Infof("Testing mysql")
		for u := range common.Userdict["MySql"] {
			for p := range common.Passwords {
				common.Mysqlwg.Add(1)
				go mysql.ScanMysql(host.Name, fmt.Sprint(common.PORTList["mysql"]), common.Userdict["mysql"][u], common.Passwords[p])
			}
			common.Mysqlwg.Wait()
		}

	} else if host.Port == common.PORTList["mssql"] {
		gologger.Infof("Testing MsSql")
		for u := range common.Userdict["mssql"] {
			for p := range common.Passwords {
				common.Mssqlwg.Add(1)
				go mssql.MsSqlAuth(host.Name, fmt.Sprint(common.PORTList["mssql"]), common.Userdict["mssql"][u], common.Passwords[p])
			}
		}
		common.Mssqlwg.Wait()

	} else if host.Port == common.PORTList["redis"] {
		gologger.Infof("Testing Redis")
		common.Rediswg.Add(1)
		go redis.RedisNullAuth(host.Name, common.PORTList["redis"])
		common.Rediswg.Wait()

	} else if host.Port == common.PORTList["mgo"] {
		gologger.Infof("Testing MongoDB")
		for u := range common.Userdict["mongodb"] {
			for p := range common.Passwords {
				common.Mongowg.Add(1)
				go mongodb.MongoAuth(host.Name, fmt.Sprint(common.PORTList["mgo"]), common.Userdict["mongodb"][u], common.Passwords[p])
			}
		}
		common.Mongowg.Wait()

	} else if host.Port == common.PORTList["oracle"] {
		gologger.Infof("Testing Oracle")
		for u := range common.Userdict["oracle"] {
			for p := range common.Passwords {
				common.Oraclewg.Add(1)
				go oracle.OracleAuth(host.Name, fmt.Sprint(common.PORTList["oracle"]), common.Userdict["oracle"][u], common.Passwords[p])
			}
		}
		common.Oraclewg.Wait()

	} else if host.Port == common.PORTList["smb"] {
		gologger.Infof("Testing SMB")
		for u := range common.Userdict["smb"] {
			for p := range common.Passwords {
				common.Smbwg.Add(1)
				go smb.SmbAuth(host.Name, fmt.Sprint(common.PORTList["smb"]), common.Userdict["smb"][u], common.Passwords[p])
			}
		}
		common.Smbwg.Wait()

	} else if host.Port == common.PORTList["ftp"] {
		gologger.Infof("Testing FTP")
		for u := range common.Userdict["ftp"] {
			for p := range common.Passwords {
				common.Smbwg.Add(1)
				go ftp.FtpAuth(host.Name, fmt.Sprint(common.PORTList["smb"]), common.Userdict["ftp"][u], common.Passwords[p])
			}
		}
		common.Smbwg.Wait()

	} else {
		gologger.Infof("Skipped: Unsupported type")
	}

}
func SshCheck(port string, sum int) int {
	for u := range common.Userdict["ssh"] {
		for p := range common.Passwords {
			if common.Statuses[sum] == false {
				common.Sshwg.Add(1)
				fmt.Println(fmt.Sprint(port), fmt.Sprint(common.PORTList["ssh"]), common.Userdict["ssh"][u], common.Passwords[p], sum)
				go ssh.ScanSsh(fmt.Sprint(port), fmt.Sprint(common.PORTList["ssh"]), common.Userdict["ssh"][u], common.Passwords[p], sum)
			}
		}
		common.Sshwg.Wait()
	}
	for {
		time.Sleep(10000 * time.Millisecond)
		if common.Statuses[sum] == true {
			fmt.Println(common.Statuses[sum])
			common.Sshwg.Wait()
			common.WeakPass.Done()
			return 1
		} else {
			common.Sshwg.Wait()
			common.WeakPass.Done()
			return 0
		}
	}
	//return 0
}
