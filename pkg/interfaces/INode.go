package interfaces

import "Scar/pkg/models"

type INode interface {
	LifeCycle() error
	Initialize() error
	SendMessage(message models.Message) error
	Listen(port int) error
	JoinNetwork() error
	UpdateRoutingTable(info models.RoutingInfo) error
}
