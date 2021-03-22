package zonal

//Zone object type defining zonal objects
type Zone struct {
	status bool
	label  string
	nodes  []int
}

//Setup Zonal layer entry point
func Setup(filePath string) error {
	return ErrProblemSettingUpZonalLayer
}

//Label Gets the label describing the specified zone
func (z *Zone) Label() string {
	return z.label
}

//SetLabel Sets the zone label to the specified string
func (z *Zone) SetLabel(l string) {
	z.label = l
}

//State Gets the activity state of the given zone
func (z *Zone) State() bool {
	return z.status
}

//SetState Sets the activity state of the given zone
func (z *Zone) SetState(state bool) {
	z.status = state
}

//NodeInZone Determines whether the specified node is part of the given zone
func (z *Zone) NodeInZone(nodeID int) bool {
	return z.isNodeInZone(nodeID)
}

//AddNodeToZone Ads the specified node to the given zone
func (z *Zone) AddNodeToZone(nodeID int) {
	if !z.isNodeInZone(nodeID) {
		z.nodes = append(z.nodes, nodeID)
	}
}

// RemoveNodeFromZone Removes the specified node from the given zone
func (z *Zone) RemoveNodeFromZone(nodeID int) {
	if index := z.nodeZoneIndex(nodeID); index != -1 {
		z.nodes = append(z.nodes[:index], z.nodes[index+1:]...)
	}
}

//NodeCount Returns the number of nodes in the given zone.
func (z *Zone) NodeCount() int {
	return len(z.nodes)
}

//ClearZone Remove all nodes from specified zone.
func (z *Zone) ClearZone(status bool) {
	z.status = status
	z.nodes = z.nodes[:0]
}

// private functions
func (z *Zone) isNodeInZone(nodeID int) bool {
	for _, node := range z.nodes {
		if node == nodeID {
			return true
		}
	}

	return false
}

func (z *Zone) nodeZoneIndex(nodeID int) int {
	for index, node := range z.nodes {
		if node == nodeID {
			return index
		}
	}
	return -1
}
