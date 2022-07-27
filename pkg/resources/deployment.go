package resources

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	extract "github.com/framsouza/eck-diagnostics-parser/pkg/extract"
	handlingfiles "github.com/framsouza/eck-diagnostics-parser/pkg/handlingfiles"
	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Deployment() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, ' ', 0)
	defer w.Flush()

	//abspath, _ := handlingfiles.FindFileAbsPathDeploy(extract.Destination, "deployments.json")
	deploypath := handlingfiles.FindFile("deployments.json", extract.Destination)

	//for _, f := range abspath {

	for _, f := range deploypath {
		if f == extract.Destination+"/kube-system/deployments.json" {
			break
		}
		if f == extract.Destination+"/istio-system/deployments.json" {
			break
		}

		config, err := load.LoadDeploy(f)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "\n\n%s\t\t\t%s\t\t%s\t\t", "DEPLOYMENT NAME", "REPLICAS", "NAMESPACE")

		for i := range config.Items {
			fmt.Fprintf(w, "\n%s\t\t\t", config.Items[i].Metadata.Name)
			fmt.Fprintf(w, "%v\t\t", config.Items[i].Spec.Replicas)
			fmt.Fprintf(w, "%v\t", config.Items[i].Metadata.Namespace)

		}
	}
}
