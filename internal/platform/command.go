package platform

import (
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadSettingsFile() {
	// Get Path to Settings
	// Getting current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	defaultSettingsFilePath := "sites/default/settings.php"
	inDrupalDocroot := false

	//Try to figure out site root
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		// @todo Gotta be wary of "sites" dirs that are not from Drupal sites
		if f.Name() == "sites" {
			inDrupalDocroot = true
		}
	}

	if inDrupalDocroot {
		// Build full path
		settingsPath := dir + "/" + defaultSettingsFilePath
		fmt.Println(settingsPath)

		// Open file for reading
		f, err := os.Open(settingsPath)
		check(err)
		defer f.Close()

		// Arbitrary reading right now for testing
		b1 := make([]byte, 5)
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b1))

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("Not in Drupal docroot. Have to do more.")
	}
	// First, Assume in site root directory
	
	// Second, Try to get path to settings.php if in site dir (look for index.php moving backwards?)
	// If we can't find the settings.php, return error

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
