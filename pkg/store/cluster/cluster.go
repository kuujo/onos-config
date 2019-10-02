package cluster

import "os"

const nodeIDEnv = "NODE_ID"

// NodeID is a node identifier provided by the environment
type NodeID string

// GetNodeID gets the local node identifier
func GetNodeID() NodeID {
	nodeID := os.Getenv(nodeIDEnv)
	if nodeID == "" {
		panic("No node ID provided by environment")
	}
	return NodeID(nodeID)
}
