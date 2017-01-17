package main

import (
	"fmt"
	"os"

	"github.com/crhntr/gomesh/mesh"
)

func init() {

}

func main() {
	desc, err := os.Open("../mesh/testdesc.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := mesh.LoadDescriptors(desc); err != nil {
		fmt.Print(err)
		return
	}
}
