package personal

func loadPersonalMappingFile() error {
	return ErrProblemLoadingPersonalMappingFile
}

// Setup personal layer
func Setup(filePath string) error {
	return loadPersonalMappingFile()
}
