package mutexpools

import (
	"errors"

	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/mutexpool/api/networks"
)

// Write is write file to coresan
func (me *MutexPoolSession) Lock(txtUID string) (bool, error) {
	jsoCmd := new(jsons.JSONObject).Factory()
	jsoCmd.PutString("txt_command", "lock")
	jsoCmd.PutString("txt_ruid", me.UUID)
	jsoCmd.PutString("txt_ouid", txtUID)

	jsoReq := networks.Request(me.MutexPoolHost.JSOConfigHost, jsoCmd)
	if jsoReq.GetInt("status") <= 0 {
		return false, errors.New(jsoReq.GetString("txt_msg"))
	}
	return true, nil
}
