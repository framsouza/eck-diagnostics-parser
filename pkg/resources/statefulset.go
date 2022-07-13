package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func StatefulSet() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadStatefulSet("/Users/francismarasouza/finding-file/tmp/default/statefulsets.json")
	fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t", "STATEFULSET NAME", "REPLICAS", "HEAP SIZE")

	for i := range config.Items {
		fmt.Fprintf(w, "\n%s\t\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%v\t\t", config.Items[i].Spec.Replicas)
		//fmt.Fprintf(w, "%v\t", config.Items[i].Spec.Template.Spec.Containers)
		for c := range config.Items[i].Spec.Template.Spec.Containers {
			for e := range config.Items[i].Spec.Template.Spec.Containers[c].Env {
				list := config.Items[i].Spec.Template.Spec.Containers[c].Env[e]
				if list.Name == "ES_JAVA_OPTS" {
					fmt.Fprintf(w, "%v\t\t\t", config.Items[i].Spec.Template.Spec.Containers[c].Env[e].Value)
				}
				// PRINT SOMETHING IF HEAP IS NOT SET OR SEND IT TO THE SUMMARY.GO
			}
		}

	}
}

// Check if JVM was set

//CHECK VOLUMES ATTACHED
