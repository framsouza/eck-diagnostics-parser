package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	extract "github.com/framsouza/eck-diagnostics-parser/pkg/extract"
	resources "github.com/framsouza/eck-diagnostics-parser/pkg/resources"
)

var (
	zipfile     = flag.String("zipfile", "", "Enter zip file name")
	destination = flag.String("destination", "./tmp/", "Enter the destination path")
)

func main() {
	flag.Parse()
	if *zipfile == "" {
		fmt.Fprintln(os.Stderr, "Please, enter the file name to be extracted (.zip)")
		flag.Usage()
		os.Exit(2)
	}
	if *destination == "" {
		fmt.Fprintln(os.Stderr, "Folder not found")
		flag.Usage()
		os.Exit(2)
	}

	err := extract.UnzipSource(*zipfile, *destination)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\nWelcome to the ECK diagnostic parser tool\n")
	resources.DiagV()
	resources.ECKV() // Remove time/date from the line
	resources.Collection()
	resources.Nodes()
	resources.Pods()
	resources.Es()
	resources.Kibana()
	resources.StatefulSet()
	resources.Deployment()
	resources.PVC()
	resources.Service()
	summary()

}

func summary() {
	fmt.Print("\n\nBASED ON THE OUTPUT ABOVE, MAKE SURE THAT:\n")
	fmt.Print("- All the services has an endpoint attached to it\n")
	fmt.Print("- All the PVC has a Bound status\n")
	fmt.Print("- All the Elasticsearch Resources are Green and the Phase is \"READY\"\n")
	fmt.Print("- All the Kibana Resources are Green and the Phase is \"READY\"\n")
	fmt.Print("- Every pod is RUNNING\n")
	fmt.Print("- Every pod has the same MEM REQUEST & MEM LIMIT\n\n")

}
