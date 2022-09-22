package redis

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"Aopo/common"
	"Aopo/gologger"
	"github.com/go-redis/redis"
	"strconv"
	"time"

	"fmt"
)

func RedisNullAuth(host string, iport int) (err error, result bool) {
	portt := strconv.Itoa(iport)
	opt := redis.Options{Addr: fmt.Sprintf("%v:%v", host, portt),
		Password: "", DB: 0, DialTimeout: 2 * time.Second}
	client := redis.NewClient(&opt)
	_, err = client.Ping().Result()
	defer client.Close()
	if err == nil {
		gologger.Infof("Redis service has empty password: " + host + ":" + fmt.Sprintln(iport))
		common.ResultsMap.Lock()
		common.ResultsMap.Credentials = append(common.ResultsMap.Credentials, common.Credential{Url: host, Port: portt, UserName: host, Password: "", Group: "Redis"})
		common.ResultsMap.Unlock()
		result = true
	}
	common.Rediswg.Done()
	return err, result
}
func RedisAuth(host string, iport int, password string) (err error, result bool) {
	portt := strconv.Itoa(iport)
	opt := redis.Options{Addr: fmt.Sprintf("%v:%v", host, portt),
		Password: "", DB: 0, DialTimeout: 5 * time.Second}
	client := redis.NewClient(&opt)
	_, err = client.Ping().Result()
	client.Close()
	if err == nil {
		gologger.Infof("Redis has empty password: " + host + ":" + fmt.Sprintln(iport))
		common.ResultsMap.Lock()
		common.ResultsMap.Credentials = append(common.ResultsMap.Credentials, common.Credential{Url: host, Port: portt, UserName: host, Password: "", Group: "Redis"})
		common.ResultsMap.Unlock()
		result = true
	}
	common.Rediswg.Done()
	return err, result
}
