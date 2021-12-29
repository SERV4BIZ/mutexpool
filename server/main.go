package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/SERV4BIZ/mutexpool/server/global"
	"github.com/SERV4BIZ/mutexpool/server/locals"
	"github.com/SERV4BIZ/mutexpool/server/utility"
)

func main() {
	var errConfig error
	global.JSOConfig, errConfig = locals.LoadConfig()
	if errConfig != nil {
		panic(errConfig)
	}

	global.MutexMapStore.Lock()
	global.MapStore = make(map[string]*global.StoreItem)
	global.MutexMapStore.Unlock()

	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint(global.AppName, " Version ", global.AppVersion))
	fmt.Println("Copyright © 2019 Serv4Biz Co.,Ltd. All Rights Reserved.")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint("Directory : ", utility.GetAppDir()))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("Loading configuration file.")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("")
	fmt.Println(global.JSOConfig.ToString())
	fmt.Println("")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	// Force GC to clear up
	go func() {
		for {
			<-time.After(time.Hour)
			runtime.GC()
		}
	}()

	global.Expire = global.JSOConfig.GetInt("int_expire")
	if global.Expire <= 0 {
		global.Expire = 60
	}

	intTime := global.JSOConfig.GetInt("int_timeout")
	if intTime <= 0 {
		intTime = 60
	}

	// Unlock mutex if expired
	go func() {
		for {
			unix := time.Now().Unix()
			global.MutexMapStore.RLock()
			for _, objItem := range global.MapStore {
				if objItem.Stamp > 0 {
					diff := unix - int64(objItem.Stamp)
					if diff > int64(global.Expire) {
						objItem.RUID = ""
						objItem.Stamp = 0
						objItem.Unlock()
					}
				}
			}
			global.MutexMapStore.RUnlock()
			<-time.After(time.Second * time.Duration(global.Expire))
		}
	}()

	router := http.NewServeMux()
	router.HandleFunc("/", WorkHandler)

	appsrv := &http.Server{
		Addr:         fmt.Sprint(":", global.JSOConfig.GetInt("int_port")),
		Handler:      router,
		ReadTimeout:  time.Duration(intTime) * time.Second,
		WriteTimeout: time.Duration(intTime) * time.Second,
	}
	appsrv.ListenAndServe()
}
