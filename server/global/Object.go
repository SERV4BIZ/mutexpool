package global

import (
	"sync"

	"github.com/SERV4BIZ/gfp/jsons"
)

type StoreItem struct {
	sync.RWMutex
	RUID  string
	OUID  string
	Stamp int
}

// AppName is name of application
var AppName string = "MUTEX POOL SERVER"

// AppVersion is version of application
var AppVersion string = "1.0.0"

// DS is split of path
var DS string = "/"

// MutexJSOConfig is mutex lock of JSOConfig
var MutexJSOConfig sync.RWMutex

// JSOConfig is config json object
var JSOConfig *jsons.JSONObject

var Expire int = 0

var MutexMapStore sync.RWMutex
var MapStore map[string]*StoreItem = nil
