package mutexpools

import (
	"fmt"

	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/gfp/uuid"
	"github.com/SERV4BIZ/mutexpool/api/networks"
)

// Factory is begin HScaleDB object
func Factory(jsoConfigHost *jsons.JSONObject) (*MutexPool, error) {
	myUUID, errUUID := uuid.NewV4()
	if errUUID != nil {
		return nil, errUUID
	}

	mutexItem := new(MutexPool)
	mutexItem.UUID = myUUID
	mutexItem.JSOConfigHost = jsoConfigHost
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("Mutex Pool Client Factory")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint("UUID : ", mutexItem.UUID))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("Ping test info.")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	jsoReq := networks.Ping(mutexItem.JSOConfigHost)
	fmt.Println(jsoReq.ToString())

	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	return mutexItem, nil
}
