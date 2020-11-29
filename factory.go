package dockerrun

import (
	"net"
	"sync"

	"github.com/containerssh/sshserver"
)

// New creates a new NetworkConnectionHandler for a specific client.
//goland:noinspection GoUnusedExportedFunction
func New(client net.TCPAddr, connectionID []byte, config Config) (sshserver.NetworkConnectionHandler, error) {
	return &networkHandler{
		mutex:        &sync.Mutex{},
		client:       client,
		connectionID: connectionID,
		config:       config,
	}, nil
}
