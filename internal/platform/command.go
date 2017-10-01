package platform

import (
	"log"
	"os"
	"fmt"
)

func ReadSettingsFile() {
	// Get Path to Settings file
	// First, Assume in site root directory
	
	// Second, Try to get path to settings.php if in site dir (look for index.php moving backwards?)
	// If we can't find the settings.php, return error

	// Testing things below
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

/*
  Modules
 */
func moduleEnable(module string) {
	//
}

func moduleDisable(module string) {
	//
}

func moduleScaffold() {
	//
}

func cacheClear(cache string) {
	//
}
