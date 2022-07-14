package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Service() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadService("/Users/francismarasouza/finding-file/tmp/default/services.json")
	ep, _ := load.LoadEndPoint("/Users/francismarasouza/finding-file/tmp/default/endpoints.json")

	fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t", "SERVICE NAME", "TYPE", "ENDPOINTS")

	for i := range config.Items {
		fmt.Fprintf(w, "\n%s\t\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Spec.Type)
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
