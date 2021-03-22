package zonal

import (
	"testing"
)

//define mock functions to avoid contamination with untested functions
func (z *Zone) mockClear() {
	z.nodes = z.nodes[:0]
	z.label = ""
	z.status = false
}

func (z *Zone) mockCount() int {
	return len(z.nodes)
}
func (z *Zone) mockAddNode(n int) {
	z.nodes = append(z.nodes, n)
}

func (z *Zone) modeDeleteNode(n int) {
	for idx, node := range z.nodes {
		if node == n {
			z.nodes = append(z.nodes[:idx], z.nodes[idx+1:]...)
		}
	}
}

func TestClearZone(test *testing.T) {

	var zone Zone
	var initialSize int

	zone.nodes = make([]int, 5) // arbitary 5 maximum

	//Empty zone
	initialSize = zone.mockCount()

	zone.mockClear()

	if zone.mockCount() != 0 {
		test.Errorf("TestClearZone - Starting with [] (%d) Expected zero got %d ", initialSize, len(zone.nodes))
	}

	//Single node in zone
	zone.mockAddNode(1)
	initialSize = zone.NodeCount()
	zone.ClearZone(false)
	if zone.mockCount() != 0 {
		test.Errorf("TestClearZone - Starting with [1] (%d) Expected zero got %d ", initialSize, len(zone.nodes))
	}

	//Multiple nodes in zone

	zone.mockAddNode(1)
	zone.mockAddNode(2)
	initialSize = zone.NodeCount()
	zone.ClearZone(false)
	if zone.mockCount() != 0 {
		test.Errorf("TestClearZone - Starting with [1,2] (%d) Expected zero got %d ", initialSize, len(zone.nodes))
	}
}

func TestNodeCount(test *testing.T) {

	var zone Zone
	var length int

	//empty zone list
	length = zone.NodeCount()
	if length != 0 {
		test.Errorf("TestNodeCount - Starting with [] (%d) Expected zero got %d ", length, len(zone.nodes))
	}

	// one item in zone list
	zone.mockAddNode(1)
	length = zone.NodeCount()
	if length != 1 {
		test.Errorf("TestNodeCount - Starting with [1] (%d) Expected 1 got %d ", length, len(zone.nodes))
	}

	// multiple items
	for idx := 2; idx <= 10; idx++ {
		zone.mockAddNode(idx)
	}
	length = zone.NodeCount()

	if length != 10 {
		test.Errorf("TestNodeCount - Starting with [1..10] (%d) Expected 10 got %d ", length, len(zone.nodes))
	}
}

func TestState(test *testing.T) {

	var zone Zone

	// State on initial creation
	state := zone.State()
	if state != false {
		test.Errorf("TestState - Newly created zone Expected %v got %v ", false, state)
	}

	//Set state then retrieve - set to true
	zone.SetState(true)
	state = zone.State()
	if state != true {
		test.Errorf("TestState - Zone with state set to %v expected %v got %v ", true, true, state)
	}

	//Set state then retrieve - set to false
	zone.SetState(false)
	state = zone.State()
	if state != false {
		test.Errorf("TestState - Zone with state set to %v expected %v got %v ", true, true, state)
	}

}

func TestLabels(test *testing.T) {

	var zone Zone

	zone.mockClear()
	label := zone.Label()
	if len(label) != 0 {
		test.Errorf("TestLabel - Newly created zones should have no label, expected '%s', got '%s'", "", label)
	}

	// Set basic label and make sure we can retrieve it
	zone.mockClear()
	initialLabel := "Sample"
	zone.SetLabel(initialLabel)
	label = zone.Label()
	if label != initialLabel {
		test.Errorf("TestLabel - Label set and retries, expected '%s', got '%s'", initialLabel, label)
	}

	//Exchange labels between two zones
	var zone1 Zone
	var zone2 Zone

	zone1.mockClear()
	zone2.mockClear()

	label1 := "First"
	label2 := "Second"
	zone1.SetLabel(label1)
	zone2.SetLabel(label2)

	// exchange them
	got1 := zone1.Label()
	got2 := zone2.Label()

	zone1.SetLabel(got2)
	zone2.SetLabel(got1)

	got1 = zone1.Label()
	got2 = zone2.Label()

	if (got1 != label2) && (got2 != label1) {
		test.Errorf("TestLabel - labels between zones swapped incorrectly")
	}

}
func TestNodeInZone(test *testing.T) {

	var zone1 Zone
	var zone2 Zone

	// empty zone
	found := zone1.NodeInZone(0)
	if found {
		test.Errorf("TestNodeInZone - Empty zone, expected %v, got %v", false, found)
	}

	// Add single node to zone and check its there, and then not
	zone1.mockAddNode(1)
	found = zone1.NodeInZone(1)
	if !found {
		test.Errorf("TestNodeInZone - Zone has single node, expected %v, got %v", true, found)
	}

	found = zone1.NodeInZone(2)
	if found {
		test.Errorf("TestNodeInZone - Zone has single node, expected %v, got %v", false, found)
	}

	//Multiple nodes test
	zone1.mockClear()
	zone2.mockClear()

	zone1.mockAddNode(1)
	zone1.mockAddNode(2)
	zone1.mockAddNode(3)

	zone2.mockAddNode(4)
	zone2.mockAddNode(5)

	found1 := zone1.NodeInZone(1) //true
	if !found1 {
		test.Errorf("TestNodeInZone - Zone has single node, expected %v, got %v", false, found)
	}
	found2 := zone1.NodeInZone(5) // false
	if found2 {
		test.Errorf("TestNodeInZone - Zone has single node, expected %v, got %v", false, found)
	}
	found3 := zone1.NodeInZone(3) // true
	if !found3 {
		test.Errorf("TestNodeInZone - Zone has single node, expected %v, got %v", false, found)
	}
	found4 := zone2.NodeInZone(5) // true
	if !found4 {
		test.Errorf("TestNodeInZone - Zone has single node, expected %v, got %v", false, found)
	}
	found5 := zone2.NodeInZone(1) // false
	if found5 {
		test.Errorf("TestNodeInZone - Zone has single node, expected %v, got %v", false, found)
	}

}

func TestRemoveNodeFromZone(test *testing.T) {

	var zone Zone
	var initialLength, finalLength int

	// remove from empty zone
	zone.mockClear()
	initialLength = zone.mockCount()
	zone.RemoveNodeFromZone(1) // doesnt exist so no change
	finalLength = zone.mockCount()
	if finalLength != 0 {
		test.Errorf("TestRemoveNodeFromZone- Starting with [] (%d) Expected 0 got %d ", initialLength, finalLength)
	}

	// remove from non empty zone but id doesnt exists
	zone.mockClear()
	zone.mockAddNode(1)
	initialLength = zone.mockCount()
	zone.RemoveNodeFromZone(2) // doesnt exist so no change
	finalLength = zone.mockCount()
	if finalLength != initialLength {
		test.Errorf("TestRemoveNodeFromZone- Starting with [1] (%d) Expected 1 got %d ", initialLength, finalLength)
	}

	//remove from non empty zone id that exists
	zone.mockClear()
	zone.mockAddNode(1)
	initialLength = zone.mockCount()
	zone.RemoveNodeFromZone(1)
	finalLength = zone.mockCount()
	if finalLength != 0 {
		test.Errorf("TestRemoveNodeFromZone- Starting with [1] (%d) Expected 0 got %d ", initialLength, finalLength)
	}

	//remove from zone with multiple entries and ID exists
	zone.mockClear()
	zone.mockAddNode(1)
	zone.mockAddNode(2)
	initialLength = zone.mockCount()
	zone.RemoveNodeFromZone(1)
	finalLength = zone.mockCount()
	if finalLength != initialLength-1 {
		test.Errorf("TestRemoveNodeFromZone- Starting with [1 2] (%d) Expected %d got %d ", initialLength, initialLength-1, finalLength)
	}
}
