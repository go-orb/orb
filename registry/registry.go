// Package registry is a component for service discovery
package registry

import (
	"errors"

	"github.com/orb-org/config/source/cli"
)

var (
	// Not found error when GetService is called.
	ErrNotFound = errors.New("service not found")
	// Watcher stopped error when watcher is stopped.
	ErrWatcherStopped = errors.New("watcher stopped")
)

func init() {
	err := cli.Flags.Add(cli.NewFlag(
		"registry",
		"mdns",
		cli.CPSlice([]string{"registry", "plugin"}),
		cli.Usage("Registry for discovery. etcd, mdns"),
	))
	if err != nil {
		panic(err)
	}
}

// The registry provides an interface for service discovery
// and an abstraction over varying implementations
// {consul, etcd, zookeeper, ...}.
type Registry interface {
	Init(aConfig any, opts ...Option) error
	Config() any
	Register(*Service, ...RegisterOption) error
	Deregister(*Service, ...DeregisterOption) error
	GetService(string, ...GetOption) ([]*Service, error)
	ListServices(...ListOption) ([]*Service, error)
	Watch(...WatchOption) (Watcher, error)
	String() string
}

type Service struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []*Endpoint       `json:"endpoints"`
	Nodes     []*Node           `json:"nodes"`
}

type Node struct {
	Id       string            `json:"id"`
	Address  string            `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

type Endpoint struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}
