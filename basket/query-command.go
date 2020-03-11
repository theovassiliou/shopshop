package shopshop

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Query and display all shopshop lists
func (sl *Basket) Query(dir string) {
	fmt.Println("Available lists: ")

	files, err := ioutil.ReadDir(os.ExpandEnv(dir))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if path.Ext(f.Name()) == ".shopshop" {
			listName := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
			if listName == sl.ListName() {
				fmt.Printf("* %s\n", listName)
			} else {
				fmt.Printf("  %s\n", listName)
			}
		}
	}

}
