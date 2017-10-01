package platform

import (
	"log"
	"os"
	"fmt"
	filepath "path/filepath"
)

func ReadSettingsFile() {
	// Get Path to Settings file
	// First, Assume in site root directory
	
	// Second, Try to get path to settings.php if in site dir (look for index.php moving backwards?)
	// If we can't find the settings.php, return error

	// Testing things below
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
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
