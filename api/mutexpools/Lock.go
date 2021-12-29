package mutexpools

import (
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/mutexpool/api/networks"
)

// Write is write file to coresan
func (me *MutexPoolSession) Lock(txtUID string) bool {
	jsoCmd := new(jsons.JSONObject).Factory()
	jsoCmd.PutString("txt_command", "lock")
	jsoCmd.PutString("txt_ruid", me.UUID)
	jsoCmd.PutString("txt_ouid", txtUID)

	jsoReq := networks.Request(me.MutexPoolHost.JSOConfigHost, jsoCmd)
	return jsoReq.GetInt("status") > 0
}
