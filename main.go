package main

import (
	"Aopo/common"
	"Aopo/gologger"
	host "Aopo/lib"
	"bufio"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"time"
)

func main() {
	banner := `
                            
       /\                     
      /  \   ___  _ __   ___  
     / /\ \ / _ \| '_ \ / _ \ 
    / ____ \ (_) | |_) | (_) |
   /_/    \_\___/| .__/ \___/ 
                 | |    By：` + aurora.Red("ExpLang & zeinlol").String() + `      
                 |_|    Version: ` + aurora.Green(common.Version).String() + `
 Github: ` + aurora.Blue("github.com/cqr-cryeye-forks/Aopo").String() + `

`
	print(banner)
	start := time.Now()
	var Info common.Flaginfo
	var AllTarget = "192.168.0.0/16,10.0.0.0/8,172.0.0.0/8"
	common.Flag(&Info)
	if Info.Passdic != "" {
		testPasswords(Info.Passdic)
	}
	if Info.All != false {
		if Info.Ip != "" || Info.Ipfile != "" {
			gologger.Errorf("Error: Too many arguments")
			os.Exit(0)
		}
		common.ResultsMap.Lock()
		common.ResultsMap.Target = AllTarget
		common.ResultsMap.Unlock()
		host.Host(AllTarget, "", "")
	} else if Info.Ip != "" {
		if Info.Ipfile != "" {
			gologger.Errorf("Error：Too many parameters")
			os.Exit(0)
		}
		common.ResultsMap.Lock()
		common.ResultsMap.Target = Info.Ip
		common.ResultsMap.Unlock()
		host.Host(Info.Ip, "", "")
	} else if Info.Ipfile != "" {
		if Info.Ip != "" {
			gologger.Errorf("Error：Too many parameters")
			os.Exit(0)
		}
		common.ResultsMap.Lock()
		common.ResultsMap.Target = Info.Ipfile
		common.ResultsMap.Unlock()
		host.Host("", Info.Ipfile, "")
	}
	t := time.Now().Sub(start)
	common.ResultsMap.Lock()
	common.ResultsMap.TotalTime = t
	common.ResultsMap.Unlock()
	common.SaveJson(Info.JsonFile)
	gologger.Printf("================================== END ==================================")
	gologger.Infof("The scan is over. Total time: " + fmt.Sprint(t) + "\n")
}

func testPasswords(file string) {
	fp, err := os.Open(file)
	if err != nil {
		fmt.Println(err) //打开文件错误
		return
	}
	buf := bufio.NewScanner(fp)
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		if IsContain(common.Passwords, line) {
			continue
		} else {
			common.Passwords = append(common.Passwords, line)
			gologger.Infof("Successfully merged: " + line + " passwords to the built-in dictionary")
		}
	}
}
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
