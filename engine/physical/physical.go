package physical

const physicalMappingFile = "./physical/mapping"

func loadPhysicalMapDefinition() error {
	return ErrProblemLoadingPhysicalMapFile
}

// Setup physical layer
func Setup(filePath string) error {
	return loadPhysicalMapDefinition()
}
