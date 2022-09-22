package common

import (
	"Aopo/gologger"
	"bytes"
	"fmt"
	"github.com/k8gege/LadonGo/str"
	"io"
	"net/http"
	"regexp"
	"time"

	//"os"
	"strings"
)

func HttpScan(url string, port int) {
	defer Httpwg.Done()
	//fmt.Println(url, port)
	result, err := HttpBanner(url, port)
	if err != nil {
		//gologger.Errorf(err.Error())
		return
	}
	ResultsMap.UrlData = append(ResultsMap.UrlData, result)
}
func IsUrl(url string, port string) (result string) {
	if !strings.Contains(url, "http") {
		url := "http://" + url + ":" + port
		return url
	} else if !strings.Contains(url, "https") {
		url := "https://" + url + ":" + port
		return url
	}
	return url
}
func HttpBanner(url string, port int) (result HttpBannerData, err error) {

	url2 := IsUrl(url, fmt.Sprint(port))
	response, err := http.Head(url2)
	title := ScanTitle(url2)
	if err != nil {
		//fmt.Println(err.Error())
		//os.Exit(2)
		return HttpBannerData{}, err
	}

	//fmt.Println(response)
	//fmt.Println(response.Status)
	Xraylist = append(Xraylist, url2)
	gologger.Infof("URL: " + url2 + " Status: " + fmt.Sprint(response.StatusCode) + " Content length: " + fmt.Sprint(response.ContentLength) + " Title: " + title)
	return HttpBannerData{Url: url2, StatusCode: response.StatusCode, ContentLength: response.ContentLength, Title: title}, err

}
func GetHtml(url string) string {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		//panic(err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			//panic(err)
			return ""
		}
	}
	return result.String()
}

func GetTitle(html string) string {
	re, _ := regexp.Compile("<[\\S\\s]+?>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)
	html = strings.Replace(html, "\n", "", -1)
	title := strings.Trim(str.GetBetween(html, "<title>", "</title>"), " ")
	return title
}

func ScanTitle(host string) (title string) {
	if strings.Contains(host, ":") {
		title = GetTitle(GetHtml(host))
	} else {
		url := "http://" + host
		title = GetTitle(GetHtml(url))

		url = "https://" + host
		title = GetTitle(GetHtml(url))
	}
	return title
}
