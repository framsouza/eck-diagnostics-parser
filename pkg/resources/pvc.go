package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func PVC() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	config, _ := load.LoadPVC("/Users/francismarasouza/finding-file/tmp/default/persistentvolumeclaims.json")
	fmt.Fprintf(w, "\n\n%s\t%s\t\t%s\t\t%s\t\t", "VOLUME NAME", "STATUS", "CAPACITY", "STORAGE CLASS NAME")

	for i := range config.Items {
		fmt.Fprintf(w, "\n%s\t", config.Items[i].Metadata.Name)
		fmt.Fprintf(w, "%v\t\t", config.Items[i].Status.Phase)
		fmt.Fprintf(w, "%v\t\t", config.Items[i].Status.Capacity.Storage)
		fmt.Fprintf(w, "%v\t", config.Items[i].Spec.StorageClassName)

	}
}

// Check if JVM was set

//CHECK VOLUMES ATTACHED
