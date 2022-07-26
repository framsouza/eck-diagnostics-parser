package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	extract "github.com/framsouza/eck-diagnostics-parser/pkg/extract"
	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func StorageClass() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, ' ', 0)
	defer w.Flush()

	scpath := extract.Destination + "/storageclasses.json"

	if _, err := os.Stat(scpath); err == nil {

		config, _ := load.LoadStorageClass(scpath)
		fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t", "STORAGE CLASS NAME", "PROVISIONER", "ALLOW EXPANSION", "VOLUME BIND MODE", "RECLAIM POLICY")

		for i := range config.Items {
			fmt.Fprintf(w, "\n%s\t\t", config.Items[i].Metadata.Name)
			fmt.Fprintf(w, "%s\t\t", config.Items[i].Provisioner)
			fmt.Fprintf(w, "%v\t\t", config.Items[i].AllowVolumeExpansion)
			fmt.Fprintf(w, "%s\t\t", config.Items[i].VolumeBindingMode)
			fmt.Fprintf(w, "%s\t\t", config.Items[i].ReclaimPolicy)

		}
	} else {
		fmt.Fprintf(w, "\n\n%s\t", "Storage Class file not found, make sure to use the latests version of ECK Diagnostics")
	}
}
