package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	extract "github.com/framsouza/eck-diagnostics-parser/pkg/extract"
	findfile "github.com/framsouza/eck-diagnostics-parser/pkg/findfile"
	resources "github.com/framsouza/eck-diagnostics-parser/pkg/resources"
)

var (
	zipfile     = flag.String("zipfile", "", "Enter zip file name")
	destination = flag.String("destination", "", "Enter the destination path")
)

// DON'T FORGET!!! REMEMBER TO FIND ANOTHER WAY TO PASS THE FILE NAMES TO THE FUNCITONS!!!!!

// CHECK IF STORAGECLASS FILE EXISTS AND PRINT IT

// TREAT ERRORS AND LOGS

// ADD OPERATOR INFORMATIONS

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
	fmt.Print("\n\n")
	findfile.FindOpNS()
	findfile.FindRsNS()
	fmt.Print("FILES")
	findfile.ListDir(*destination)
	//resources.Events()

}
