package main

import (
	"Scar/pkg/interfaces"
	"Scar/pkg/models"
	"Scar/pkg/node"
)

type NodeConfig struct {
	ID             string
	Port           int
	Neighbours     []string
	RoutingTable   []models.RoutingInfo
	SecCredentials map[string]string
}

func main() {

	// simulationg network: there is no common network entity that is aware of all nodes
	var nodesList []interfaces.INode

	// Define the configuration for each node
	nodeConfigs := []NodeConfig{
		{
			ID:             "A",
			Port:           8080,
			Neighbours:     []string{"B", "C"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "B",
			Port:           8080,
			Neighbours:     []string{"A", "C", "D"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "C",
			Port:           8080,
			Neighbours:     []string{"B", "A", "E"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "D",
			Port:           8080,
			Neighbours:     []string{"B"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "E",
			Port:           8080,
			Neighbours:     []string{"C", "F"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "F",
			Port:           8080,
			Neighbours:     []string{"E", "G"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "G",
			Port:           8080,
			Neighbours:     []string{"F", "H", "I", "J"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "H",
			Port:           8080,
			Neighbours:     []string{"G", "I", "J"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "I",
			Port:           8080,
			Neighbours:     []string{"G", "H", "J"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
		{
			ID:             "J",
			Port:           8080,
			Neighbours:     []string{"H", "I", "G"},
			RoutingTable:   []models.RoutingInfo{},
			SecCredentials: map[string]string{},
		},
	}

	//#		A -- B ---- D
	//#	    |   /
	//#	    |  /
	//#	    | /
	//		|/
	//#		C ---- E ---- F ---- G ---- H
	//#							 |	 X  |
	//#					 		 I ---- J

	for _, config := range nodeConfigs {
		var nodeInstance interfaces.INode = node.NewBasicNode(config.ID, config.Port, config.Neighbours, config.RoutingTable, config.SecCredentials)
		nodesList = append(nodesList, nodeInstance)
	}

}
