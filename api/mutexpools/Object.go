package mutexpools

import (
	"sync"

	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/gfp/uuid"
)

// MutexPool is main object
type MutexPool struct {
	sync.RWMutex

	UUID          string
	JSOConfigHost *jsons.JSONObject
}

type MutexPoolSession struct {
	MutexPoolHost *MutexPool
	UUID          string
}

func (me *MutexPool) NewSession() (*MutexPoolSession, error) {
	myUUID, errUUID := uuid.NewV4()
	if errUUID != nil {
		return nil, errUUID
	}

	sess := new(MutexPoolSession)
	sess.MutexPoolHost = me
	sess.UUID = myUUID
	return sess, nil
}
