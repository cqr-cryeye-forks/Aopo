package common

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/gologger"
	"net"
	"strconv"
	"sync"
	"time"
)

var scan sync.WaitGroup

func ScanWebPort(host string) {
	defer Portwebwg.Done()
	for _, p := range WebPort {
		scan.Add(1)
		//CheckPort(net.ParseIP(host),p)
		go PortWebCheck(host, p)
	}
	scan.Wait()
	//Portwebwg.Done()
}
func ScanPort(host string) {
	defer Portwg.Done()
	for p := range PORTList {
		scan.Add(1)
		go PortCheck(host, p, PORTList[p])
		//var IsPortOpen = go PortCheck(host, PORTList[p])
		//if IsPortOpen {
		//	TotalResults.HostData = append(TotalResults.HostData,
		//		HostInfo{
		//			Name:   host,
		//			Port:   PORTList[p],
		//			Status: "Open"})
		//}
	}
	scan.Wait()
	//Portwg.Done()
}
func PortWebCheck(host string, port int) bool {
	defer scan.Done()
	p := strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", host+":"+p, time.Second*3)
	if err != nil {
		//fmt.Println(host, p, "Close")
		return false
	} else {
		gologger.Infof("WEB Service: " + host + " " + p + " Open")
		//_, ok := PortList[host]
		ResultsMap.Lock()
		ResultsMap.PortWebList[host] = append(ResultsMap.PortWebList[host], port)
		ResultsMap.HostList = append(ResultsMap.HostList, HostInfo{Name: host, Type: "web", Port: port})
		ResultsMap.Unlock()
		conn.Close()
		//scan.Done()
		return true
	}
}
func PortCheck(host string, name string, port int) bool {
	defer scan.Done()
	p := strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", host+":"+p, time.Second*3)
	if err != nil {
		//fmt.Println(host, p, "Close")
		return false
	} else {
		gologger.Infof("Host: " + host + " " + p + " Open")
		//_, ok := PortList[host]
		ResultsMap.Lock()
		ResultsMap.PortList[host] = append(ResultsMap.PortList[host], port)
		ResultsMap.HostList = append(ResultsMap.HostList, HostInfo{Name: host, Type: name, Port: port})
		ResultsMap.Unlock()
		conn.Close()
		return true
	}
}
