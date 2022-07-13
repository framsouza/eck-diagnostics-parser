package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Pods() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadPods("/Users/francismarasouza/eck-diagnostics-parser/pods.json")
	fmt.Fprintf(w, "\n\n%s\t%s\t\t%s\t%s\t%s\t%s\t%s\t", "PODS NAME", "STATUS", "MEM REQUEST", "MEM LIMIT", "CPU REQUEST", "CPU LIMIT", "INIT CONTAINER STATUS")

	for i := range config.Items {
		fmt.Fprintf(w, "\n%s\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Phase)
		for r := range config.Items[i].Spec.Containers {
			if config.Items[i].Spec.Containers[r].Resources.Requests.Memory == "" {
				fmt.Fprintf(w, "%s\t", "No Request set")
			} else {
				fmt.Fprintf(w, "%s\t", config.Items[i].Spec.Containers[r].Resources.Requests.Memory)
			}

			if config.Items[i].Spec.Containers[r].Resources.Limits.Memory == "" {
				fmt.Fprintf(w, "%s\t", "No limit set")
			} else {
				fmt.Fprintf(w, "%s\t", config.Items[i].Spec.Containers[r].Resources.Limits.Memory)
			}

			if config.Items[i].Spec.Containers[r].Resources.Requests.Cpu == "" {
				fmt.Fprintf(w, "%s\t", "No Request set")
			} else {
				fmt.Fprintf(w, "%s\t", config.Items[i].Spec.Containers[r].Resources.Requests.Cpu)
			}

			if config.Items[i].Spec.Containers[r].Resources.Limits.Cpu == "" {
				fmt.Fprintf(w, "%s\t", "No limit set")
			} else {
				fmt.Fprintf(w, "%s\t", config.Items[i].Spec.Containers[r].Resources.Limits.Cpu)

			}
		}
		//Check initcontainers
		for c := range config.Items[i].Status.InitContainerStatuses {
			fmt.Fprintf(w, "%s %s, ", config.Items[i].Status.InitContainerStatuses[c].Name, config.Items[i].Status.InitContainerStatuses[c].State.Terminated.Reason)

		}

	}

	// Print labels to compare with the service labels
}
