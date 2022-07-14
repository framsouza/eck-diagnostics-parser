package resources

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func DiagV() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadDiagV("/Users/francismarasouza/finding-file/tmp/version.json")
	fmt.Println("\nDiagnostic version is", config.DiagnosticsVersion.Version)

}

func ECKV() {
	var version string
	b, e := ioutil.ReadFile("/Users/francismarasouza/finding-file/tmp/eck-diagnostics.log")
	if e != nil {
		panic(e)
	}
	array := bytes.Split(b, []byte("\n"))
	version = string(array[2])

	fmt.Printf("%s\n", version)

}

func Collection() {
	config, _ := load.LoadManifest("/Users/francismarasouza/unzip-go/manifest.json")
	fmt.Println("Diagnostic Collection date is", config.CollectionDate)

}
