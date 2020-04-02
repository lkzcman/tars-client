package tars_client

import (
	"github.com/lkzcman/tars-client/util/endpoint"
)

var cltCfg *clientConfig

// GetClientConfig : Get client config
func GetClientConfig() *clientConfig {
	return cltCfg
}

type adapterConfig struct {
	Endpoint endpoint.Endpoint
	Protocol string
	Obj      string
	Threads  int
}

type clientConfig struct {
	Locator                 string
	stat                    string
	property                string
	modulename              string
	refreshEndpointInterval int
	reportInterval          int
	AsyncInvokeTimeout      int
}
