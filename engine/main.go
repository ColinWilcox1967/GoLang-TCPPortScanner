package main

import (
	"fmt"
	"strings"
	"os"

//	fileutils "../utilities/fileutils"
	configfile "../utilities/configfile"

	access "./access"
	personal "./personal"
	physical "./physical"
	zonal "./zonal"
)

const configFilePath = "./config.txt"

type setupFilePaths struct {
	Physical,
	Access,
	Zonal,
	Personal string
}

func loadConfigFile(filePaths *setupFilePaths) bool {
	fileReadOK, lines := configfile.Read (configFilePath)

	if !fileReadOK  {
		return false
	}

	for _, line := range lines {
		if strings.Contains(line.Key(), "physical") {
			filePaths.Physical = line.Value()
		}
		if strings.Contains(line.Key(), "personal") {
			filePaths.Personal = line.Value()
		}
		if strings.Contains(line.Key(), "zonal") {
			filePaths.Zonal = line.Value()
		}
		if strings.Contains(line.Key(), "access") {
			filePaths.Access = line.Value()
		}
	}

	return false
}

func createEngineLayers(filePaths setupFilePaths) error {

	var err error

	if err = physical.Setup(filePaths.Physical); err != nil {
		return err
	}

	if err = access.Setup(filePaths.Access); err != nil {
		return err
	}

	if err = zonal.Setup(filePaths.Zonal); err != nil {
		return err
	}

	if err = personal.Setup(filePaths.Personal); err != nil {
		return err
	}

	return nil
}

func main() {

	var filePaths setupFilePaths

	if !configfile.Exists(configFilePath) {
		fmt.Printf ("Config file '%s' not found.\n", configFilePath)
		os.Exit(1)
	}
	if err := loadConfigFile(&filePaths); !err{
		status := createEngineLayers(filePaths)
		if status != nil {
			fmt.Printf("Problem creating engine : status %d\n", status)
		}
	} else {
		fmt.Printf("%v", err) // for now
	}
}
