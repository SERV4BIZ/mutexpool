package mutexpools

import (
	"github.com/SERV4BIZ/gfp/jsons"
)

// New is begin HScaleDB object same Factory
func New(jsoConfigHost *jsons.JSONObject) (*MutexPool, error) {
	return Factory(jsoConfigHost)
}
