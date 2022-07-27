package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	extract "github.com/framsouza/eck-diagnostics-parser/pkg/extract"
	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

var memorycapacity string

func Nodes() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, ' ', 0)
	defer w.Flush()

	nodespath := extract.Destination + "/nodes.json"

	config, _ := load.LoadNodes(nodespath)

	fmt.Fprintf(w, "\n%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t", "NODE NAME", "CPU CAPACITY", "CPU ALLOCATED", "MEM CAPACITY", "MEM ALLOCATED", "VERSION", "NODE READY")

	for i := range config.Items {
		fmt.Fprintf(w, "\n%s\t\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Capacity.Cpu)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Allocatable.Cpu)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Capacity.Memory)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Allocatable.Memory)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.NodeInfo.KubeletVersion)
		for r := range config.Items[i].Status.Conditions {
			if config.Items[i].Status.Conditions[r].Type == "Ready" && config.Items[i].Status.Conditions[r].Status == "True" {
				fmt.Fprintf(w, "%s", config.Items[i].Status.Conditions[r].Status)
			}
		}

	}
}
