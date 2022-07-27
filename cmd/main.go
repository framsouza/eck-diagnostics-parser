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
	destination = "./tmp/"
)

func main() {
	flag.Parse()
	if *zipfile == "" {
		fmt.Fprintln(os.Stderr, "Please, enter the file name to be extracted (.zip)")
		flag.Usage()
		os.Exit(2)
	}

	err := extract.UnzipSource(*zipfile, destination)
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
	//resources.PVC()
	resources.Service()
	resources.StorageClass()
	summary()
	// clean up destination dir
	os.RemoveAll(destination)

}

func summary() {
	fmt.Print("\n\nBASED ON THE OUTPUT ABOVE, MAKE SURE THAT:\n")
	fmt.Print("- The Elasticsearch services has an endpoint attached to it, if there's no endpoint the services won't be able to connect to the pods\n")
	fmt.Print("- All the PVC has a Bound status\n")
	fmt.Print("- All the Elasticsearch Resources are Green and the Phase is \"READY\", if the status is \"ApplyingChanges\" check the Operator logs\n")
	fmt.Print("- All the Kibana Resources are Green and the Phase is \"READY\" or \"ESTABLISHED\", if the status is \"ApplyingChanges\" check the Operator logs\n")
	fmt.Print("- All the Elasticsearch and Kibana pods has the RUNNING status, if not, check the pod logs\n")
	fmt.Print("- All the Elasticsearch has the same MEM REQUEST & MEM LIMIT, it's very important to ensure quality of Serivce\n")
	fmt.Print("- Starting with Elasticsearch 7.11, unless manually overridden, heap size is automatically calculated based on the available memory, if the HEAP SIZE column is empty for the Elasticsearch statefulsets, make sure Elasticsearch is >7.11\n")
	fmt.Print("- If you desire to increase the disk size, make sure you are ran the latest ECK diagnostics and check the if ALLOW EXPANSION is true, if that is the case you can easily change the PVC volume size in the Elasticsearch manifest \n\n")

}
