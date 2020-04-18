package domain

import (
	"github.com/usagiga/distable/entity"
)

// ServerArrayModelImpl is struct implemented `ServerArrayModel`.
type ServerArrayModelImpl struct {}

// NewServerArrayModel initializes `ServerArrayModel`.
func NewServerArrayModel() ServerArrayModel {
	return &ServerArrayModelImpl{}
}

// GetMasters gets all master servers.
func (s *ServerArrayModelImpl) GetMasters(servers []entity.ServerContext) (masters []entity.ServerContext) {
	masters = []entity.ServerContext{}

	// Append all master servers into `masters`
	for _, serv := range servers {
		if serv.ServerType != entity.Master {
			continue
		}

		masters = append(masters, serv)
	}

	return masters
}

// GetSlaves gets all slave servers.
func (s *ServerArrayModelImpl) GetSlaves(servers []entity.ServerContext) (slaves []entity.ServerContext) {
	slaves = []entity.ServerContext{}

	// Append all master servers into `slaves`
	for _, serv := range servers {
		if serv.ServerType != entity.Slave {
			continue
		}

		slaves = append(slaves, serv)
	}

	return slaves
}

