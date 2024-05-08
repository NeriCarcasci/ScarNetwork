package node

import (
	"Scar/pkg/models"
	"fmt"
)

type BasicNode struct {
	ID             string
	ListeningPort  int
	Neighbours     []string
	RoutingTable   []models.RoutingInfo
	SecCredentials map[string]string
	MessageChannel chan models.Message
}

func NewBasicNode(id string, listeningPort int, neighbours []string, routingTable []models.RoutingInfo, secCredentials map[string]string) *BasicNode {
	return &BasicNode{
		ID:             id,
		ListeningPort:  listeningPort,
		Neighbours:     neighbours,
		RoutingTable:   routingTable,
		SecCredentials: secCredentials,
		MessageChannel: make(chan models.Message),
	}
}

func (node *BasicNode) JoinNetwork() error {
	fmt.Println(">> BasicNode:", node.ID, " Joining Network")
	// start listening thread

	// broadcast joining message to neighbours

	return nil
}

func (node *BasicNode) SendMessage(message models.Message) error {
	fmt.Println(">> BasicNode:", node.ID, " Listening on port:", node.ListeningPort)
	// listen to incoming messages
	go func() {
		for {
			message := <-node.MessageChannel // Receive message from channel
			fmt.Println(">> BasicNode:", node.ID, " Received Message:", message)
			// Process the message...
		}
	}()
	return nil
}

func (node *BasicNode) Listen(port int) error {
	fmt.Println(">> BasicNode:", node.ID, " Listening on port:", port)
	// listen to incoming messages
	return nil
}

func (node *BasicNode) UpdateRoutingTable(info models.RoutingInfo) error {
	fmt.Println(">> BasicNode:", node.ID, " Updating Routing Table")
	// update routing table
	return nil
}

func (node *BasicNode) Initialize() error {
	fmt.Println(">> BasicNode:", node.ID, " Initializing")

	go node.Listen(node.ListeningPort)

	node.JoinNetwork()

	return nil
}

func (node *BasicNode) LifeCycle() error {
	fmt.Println(">> BasicNode:", node.ID, " LifeCycle")
	// start the node
	return nil
}
