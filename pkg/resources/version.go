package resources

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	extract "github.com/framsouza/eck-diagnostics-parser/pkg/extract"
	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func DiagV() {
	diagv := extract.Destination + "/version.json"

	config, err := load.LoadDiagV(diagv)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nDiagnostic version is", config.DiagnosticsVersion.Version)

}

func ECKV() {
	var version string
	eckv := extract.Destination + "/eck-diagnostics.log"

	b, e := ioutil.ReadFile(eckv)
	if e != nil {
		panic(e)
	}
	array := bytes.Split(b, []byte("\n"))
	version = string(array[2])

	fmt.Printf("%s\n", version)

}

func Collection() {
	collec := extract.Destination + "/manifest.json"

	config, _ := load.LoadManifest(collec)
	fmt.Println("Diagnostic Collection date is", config.CollectionDate)

}
