package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Es() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadES("/Users/francismarasouza/finding-file/tmp/default/elasticsearch.json")
	fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t%s\t\t%s\t", "ES NAME", "STATUS", "VERSION", "PHASE", "NODES")

	for i := range config.Items {
		// Add condition to print "NOT SPECIFIED" in case there's not limit/request specified

		fmt.Fprintf(w, "\n%s\t\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Health)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Spec.Version)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Phase)
		fmt.Fprintf(w, "%v\t", config.Items[i].Status.AvailableNodes)

		// Print labels to compare with the service labels
	}

	//ADD NODES CONDITIONS
}
