package main

import (
	"fmt"

	resources "github.com/framsouza/eck-diagnostics-parser/pkg/resources"
)

// DON'T FORGET!!! REMEMBER TO FIND ANOTHER WAY TO PASS THE FILE NAMES TO THE FUNCITONS!!!!!
// CREATE A SUMMARY summary.go CHECKING IF THE SVC LABELS EXISTS, AND WITH SOME RECOMMENDATIONS
// CHECK IF STORAGECLASS FILE EXISTS AND PRINT IT
// TREAT ERRORS AND LOGS

func main() {
	fmt.Print("\nWelcome to the ECK diagnostic parser tool\n")
	// Print ECK version
	resources.DiagV()
	resources.ECKV() // Remove time/date from the line
	resources.Nodes()
	resources.Pods()
	resources.Es()
	resources.Kibana()
	resources.StatefulSet()
	resources.Deployment()
	resources.PVC()
	//resources.Events()

}
