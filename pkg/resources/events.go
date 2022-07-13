package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func Events() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadEvents("/Users/francismarasouza/finding-file/tmp/default/events.json")
	fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t%s\t\t", "TIMESTAMP", "TYPE", "KIND", "MESSAGE")

	for i := range config.Items {
		fmt.Fprintf(w, "\n%s\t\t", config.Items[i].LastTimestamp)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Type)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].InvolvedObject.Kind)
		fmt.Fprintf(w, "%s\t\t", config.Items[i].Message)

	}

	//ADD NODES CONDITIONS
}
