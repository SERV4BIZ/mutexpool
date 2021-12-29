package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/mutexpool/server/global"
)

func Lock(ruid string, ouid string) bool {
	global.MutexMapStore.RLock()
	objItem, blnOk := global.MapStore[ouid]
	global.MutexMapStore.RUnlock()

	if !blnOk {
		objItem = new(global.StoreItem)
		objItem.RUID = ruid
		objItem.OUID = ouid

		global.MutexMapStore.Lock()
		global.MapStore[ouid] = objItem
		global.MutexMapStore.Unlock()
	}

	objItem.Lock()
	objItem.RUID = ruid
	objItem.Stamp = int(time.Now().Unix())
	return true
}

func Unlock(ruid string, ouid string) bool {
	global.MutexMapStore.RLock()
	objItem, blnOk := global.MapStore[ouid]
	global.MutexMapStore.RUnlock()

	if blnOk && objItem.RUID == ruid && objItem.RUID != "" && objItem.Stamp > 0 {
		objItem.RUID = ""
		objItem.Stamp = 0
		objItem.Unlock()
		return true
	}
	return false
}

func WorkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsoResult := jsons.JSONObjectFactory()
	jsoResult.PutInt("status", 0)

	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not read body from http request [ ", err, " ]"))
	} else {
		jsoCmd, errCmd := jsons.JSONObjectFromString(string(buffer))
		if errCmd != nil {
			jsoResult.PutString("txt_msg", fmt.Sprint("Can not load command from json string buffer [ ", errCmd, " ]"))
		} else {
			txtCmd := jsoCmd.GetString("txt_command")
			txtRUID := jsoCmd.GetString("txt_ruid")
			txtOUID := jsoCmd.GetString("txt_ouid")
			if txtCmd == "lock" {
				if Lock(txtRUID, txtOUID) {
					jsoResult.PutInt("status", 1)
				}
			} else if txtCmd == "unlock" {
				if Unlock(txtRUID, txtOUID) {
					jsoResult.PutInt("status", 1)
				}
			} else if txtCmd == "ping" {
				jsoResult.PutInt("status", 1)
			}
		}
	}
	w.Write([]byte(jsoResult.ToString()))
}
