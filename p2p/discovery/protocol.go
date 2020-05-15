package discovery

import (
	"net"
	"time"

	"github.com/libonomy/libonomy-p2p/log"
	"github.com/libonomy/libonomy-p2p/p2p/node"
	"github.com/libonomy/libonomy-p2p/p2p/p2pcrypto"
	"github.com/libonomy/libonomy-p2p/p2p/server"
	"github.com/libonomy/libonomy-p2p/p2p/service"
)

type protocolRoutingTable interface {
	GetAddress() *KnownAddress
	AddAddresses(n []*node.Info, src *node.Info)
	AddAddress(n *node.Info, src *node.Info)
	AddressCache() []*node.Info
}

type protocol struct {
	local     *node.Info
	table     protocolRoutingTable
	logger    log.Log
	msgServer *server.MessageServer
}

func (p *protocol) SetLocalAddresses(tcp, udp int) {
	p.local.ProtocolPort = uint16(tcp)
	p.local.DiscoveryPort = uint16(udp)
}

// Name is the name if the protocol.
const Name = "/udp/v2/discovery"

// MessageBufSize is the buf size we give to the messages channel
const MessageBufSize = 1000

// MessageTimeout is the timeout we tolerate when waiting for a message reply
const MessageTimeout = time.Second * 5 // TODO: Parametrize

// PingPong is the ping protocol ID
const PingPong = 0

// GetAddresses is the findnode protocol ID
const GetAddresses = 1

// newProtocol is a constructor for a protocol protocol provider.
func newProtocol(local p2pcrypto.PublicKey, rt protocolRoutingTable, svc server.Service, log log.Log) *protocol {
	s := server.NewMsgServer(svc, Name, MessageTimeout, make(chan service.DirectMessage, MessageBufSize), log)
	d := &protocol{
		local:     &node.Info{ID: local.Array(), IP: net.IPv4zero, ProtocolPort: 7513, DiscoveryPort: 7513},
		table:     rt,
		msgServer: s,
		logger:    log,
	}

	// XXX Reminder: for discovery protocol to work you must call SetLocalAddresses with updated ports from the socket.

	d.msgServer.RegisterMsgHandler(PingPong, d.newPingRequestHandler())
	d.msgServer.RegisterMsgHandler(GetAddresses, d.newGetAddressesRequestHandler())
	return d
}

// Close stops the message server protocol from serving requests.
func (p *protocol) Close() {
	p.msgServer.Close()
}
