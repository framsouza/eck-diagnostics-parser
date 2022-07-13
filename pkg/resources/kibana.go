package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Kibana() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadKibana("/Users/francismarasouza/finding-file/tmp/default/kibana.json")
	fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t%s\t%s\t", "KB NAME", "STATUS", "VERSION", "PHASE", "NODES")

	for i := range config.Items {

		fmt.Fprintf(w, "\n%s\t\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Health)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Spec.Version)
		fmt.Fprintf(w, "%s\t", config.Items[i].Status.AssociationStatus)
		fmt.Fprintf(w, "%v\t", config.Items[i].Status.AvailableNodes)

		// Print labels to compare with the service labels
	}

	//ADD NODES CONDITIONS
}
