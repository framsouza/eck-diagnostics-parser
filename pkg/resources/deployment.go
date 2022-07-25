package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Deployment() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadDeploy("/Users/francismarasouza/eck-diagnostics-parser/tmp/default/deployments.json")

	fmt.Fprintf(w, "\n\n%s\t\t\t%s\t\t", "DEPLOYMENT NAME", "REPLICAS")

	for i := range config.Items {
		fmt.Fprintf(w, "\n%s\t\t\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%v\t", config.Items[i].Spec.Replicas)

	}
}
