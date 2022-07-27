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

func StatefulSet() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, ' ', 0)
	defer w.Flush()

	absPath, err := handlingfiles.FindFileAbsPathSTS(extract.Destination, "statefulsets.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range absPath {
		if f == extract.Destination+"/kube-system/statefulsets.json" {
			break
		}
		if f == extract.Destination+"/istio-system/statefulsets.json" {
			break
		}

		config, _ := load.LoadStatefulSet(f)

		fmt.Fprintf(w, "\n\n%s\t\t\t%s\t\t%s\t\t%s\t\t", "STATEFULSET NAME", "REPLICAS", "NAMESPACE", "HEAP SIZE")

		for i := range config.Items {
			fmt.Fprintf(w, "\n%s\t\t\t", config.Items[i].Metadata.Name)
			fmt.Fprintf(w, "%v\t\t", config.Items[i].Spec.Replicas)
			fmt.Fprintf(w, "%v\t\t", config.Items[i].Metadata.Namespace)
			for c := range config.Items[i].Spec.Template.Spec.Containers {
				for e := range config.Items[i].Spec.Template.Spec.Containers[c].Env {
					list := config.Items[i].Spec.Template.Spec.Containers[c].Env[e]
					if list.Name == "ES_JAVA_OPTS" {
						fmt.Fprintf(w, "%v\t\t", config.Items[i].Spec.Template.Spec.Containers[c].Env[e].Value)
					}
				}
			}

		}
	}

}
