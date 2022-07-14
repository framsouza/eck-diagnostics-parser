package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Nodes() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadNodes("/Users/francismarasouza/finding-file/tmp/nodes.json")
	fmt.Fprintf(w, "\n%s\t%s\t%s\t%s\t\t%s\t\t%s\t%s\t", "NODE NAME", "CPU CAPACITY", "CPU ALLOCATED", "MEM CAPACITY", "MEM ALLOCATED", "VERSION", "NODE CONDITIONS")

	for i := range config.Items {
		fmt.Fprintf(w, "\n%s\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%s\t", config.Items[i].Status.Capacity.Cpu)
		fmt.Fprintf(w, "%s\t", config.Items[i].Status.Allocatable.Cpu)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Capacity.Memory)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Status.Allocatable.Memory)
		fmt.Fprintf(w, "%s\t", config.Items[i].Status.NodeInfo.KubeletVersion)
		fmt.Fprintf(w, "%s %s, ", config.Items[i].Status.Conditions[0].Type, config.Items[i].Status.Conditions[0].Status)
		fmt.Fprintf(w, "%s %s, ", config.Items[i].Status.Conditions[1].Type, config.Items[i].Status.Conditions[0].Status)
		fmt.Fprintf(w, "%s %s, ", config.Items[i].Status.Conditions[2].Type, config.Items[i].Status.Conditions[0].Status)
		fmt.Fprintf(w, "%s %s", config.Items[i].Status.Conditions[3].Type, config.Items[i].Status.Conditions[0].Status)

	}
}
