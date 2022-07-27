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

func Service() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, ' ', 0)
	defer w.Flush()

	absPath, _ := handlingfiles.FindFileAbsPathServices(extract.Destination, "services.json")
	endPoint, _ := handlingfiles.FindFileAbsPathEndpoint(extract.Destination, "endpoints.json")

	for _, f := range absPath {
		if f == extract.Destination+"/kube-system/services.json" {
			break
		}
		if f == extract.Destination+"/elastic-system/services.json" {
			break
		}
		if f == extract.Destination+"/istio-system/services.json" {
			break
		}

		config, err := load.LoadService(f)
		if err != nil {
			log.Fatal(err)
		}
		ep, err := load.LoadEndPoint(endPoint[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t%s\t\t", "SERVICE NAME", "TYPE", "NAMESPACE", "ENDPOINTS")

		for i := range config.Items {

			fmt.Fprintf(w, "\n%s\t\t", config.Items[i].Metadata.Name)
			fmt.Fprintf(w, "%s\t\t", config.Items[i].Spec.Type)
			fmt.Fprintf(w, "%s\t\t", config.Items[i].Metadata.Namespace)
			for e := range ep.Items {
				if ep.Items[e].Metadata.Name == config.Items[i].Metadata.Name {
					for a := range ep.Items[e].Subsets {
						for ip := range ep.Items[e].Subsets[a].Addresses {
							fmt.Fprintf(w, "%s, ", ep.Items[e].Subsets[a].Addresses[ip].Ip)
						}
					}
				}

			}

		}

	}
}
