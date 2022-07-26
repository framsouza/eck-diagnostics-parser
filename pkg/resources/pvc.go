package resources

import (
	"fmt"
	"os"
	"text/tabwriter"

	extract "github.com/framsouza/eck-diagnostics-parser/pkg/extract"
	handlingfiles "github.com/framsouza/eck-diagnostics-parser/pkg/handlingfiles"
	load "github.com/framsouza/eck-diagnostics-parser/pkg/load"
)

func PVC() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, ' ', 0)
	defer w.Flush()

	absPath, _ := handlingfiles.FindFileAbsPathPVC(extract.Destination, "persistentvolumeclaims.json")

	for _, f := range absPath {
		if f == extract.Destination+"/kube-system/persistentvolumeclaims.json" {
			break
		}
		if f == extract.Destination+"/istio-system/persistentvolumeclaims.json" {
			break
		}
		if f == extract.Destination+"/elastic-system/persistentvolumeclaims.json" {
			break
		}

		config, _ := load.LoadPVC(f)

		fmt.Fprintf(w, "\n\n%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t", "PVC NAME", "STATUS", "CAPACITY", "VOLUME NAME", "SC NAME", "ACCESS MODE")

		for i := range config.Items {
			fmt.Fprintf(w, "\n%s\t\t", config.Items[i].Metadata.Name)
			fmt.Fprintf(w, "%v\t\t", config.Items[i].Status.Phase)
			fmt.Fprintf(w, "%v\t\t", config.Items[i].Status.Capacity.Storage)
			fmt.Fprintf(w, "%v\t\t", config.Items[i].Spec.VolumeName)
			fmt.Fprintf(w, "%v\t\t", config.Items[i].Spec.StorageClassName)
			fmt.Fprintf(w, "%v\t", config.Items[i].Spec.AccessModes)

		}
	}
}
