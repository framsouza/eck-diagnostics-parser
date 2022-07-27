package resources

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/framsouza/eck-diagnostics-parser/pkg/extract"
	handlingfiles "github.com/framsouza/eck-diagnostics-parser/pkg/handlingfiles"
	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Pods() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, ' ', 0)
	defer w.Flush()

	abspathpods, _ := handlingfiles.FindFileAbsPathPods(extract.Destination, "pods.json")

	for _, f := range abspathpods {
		if f == extract.Destination+"/istio-system/pods.json" {
			break
		}

		config, err := load.LoadPods(f)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t", "PODS NAME", "NAMESPACE", "STATUS", "MEM REQUEST", "MEM LIMIT", "CPU REQUEST", "CPU LIMIT", "NODE NAME")

		for i := range config.Items {
			fmt.Fprintf(w, "\n%s\t", config.Items[i].Metadata.Name)
			fmt.Fprintf(w, "\t%s\t", config.Items[i].Metadata.Namespace)
			fmt.Fprintf(w, "\t%s\t", config.Items[i].Status.Phase)
			if config.Items[i].Spec.Containers[0].Resources.Requests.Memory == "" {
				fmt.Fprintf(w, "\t%s\t", "Not set")
			} else {
				fmt.Fprintf(w, "\t%s\t", config.Items[i].Spec.Containers[0].Resources.Requests.Memory)
			}

			if config.Items[i].Spec.Containers[0].Resources.Limits.Memory == "" {
				fmt.Fprintf(w, "\t%s\t", "Not set")
			} else {
				fmt.Fprintf(w, "\t%s\t", config.Items[i].Spec.Containers[0].Resources.Limits.Memory)
			}

			if config.Items[i].Spec.Containers[0].Resources.Requests.Cpu == "" {
				fmt.Fprintf(w, "\t%s\t", "Not set")
			} else {
				fmt.Fprintf(w, "\t%s\t", config.Items[i].Spec.Containers[0].Resources.Requests.Cpu)
			}

			if config.Items[i].Spec.Containers[0].Resources.Limits.Cpu == "" {
				fmt.Fprintf(w, "\t%s\t", "Not set")
			} else {
				fmt.Fprintf(w, "\t%s\t", config.Items[i].Spec.Containers[0].Resources.Limits.Cpu)

			}
			fmt.Fprintf(w, "\t%s\t", config.Items[i].Spec.NodeName)

		}

		//Check initcontainers

		//for c := range config.Items[i].Status.InitContainerStatuses {
		//	fmt.Fprintf(w, "%s %s, ", config.Items[i].Status.InitContainerStatuses[c].Name, config.Items[i].Status.InitContainerStatuses[c].State.Terminated.Reason)

		//}

	}
}
