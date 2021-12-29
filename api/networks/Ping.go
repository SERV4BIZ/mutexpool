package networks

import (
	"github.com/SERV4BIZ/gfp/jsons"
)

// Ping is test network
func Ping(jsoHost *jsons.JSONObject) *jsons.JSONObject {
	jsoCmd := new(jsons.JSONObject).Factory()
	jsoCmd.PutString("txt_command", "ping")
	return Request(jsoHost, jsoCmd)
}
