package findfile

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	files []string
)

func FindFile(filename, d string) []string {
	filepath.Walk(d, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(filename, f.Name())
			if err == nil && r {
				absolutefilepath, err := filepath.Abs(path)
				if err != nil {
					log.Fatal(err)
				}

				files = append(files, absolutefilepath)

			}

		}
		return nil
	})

	return files

}

/*
func FindOpNS() string {
	str := "2022/03/10 07:44:23 ECK diagnostics with parameters: {DiagnosticImage:docker.elastic.co/eck-dev/support-diagnostics:8.1.4 ECKVersion: Kubeconfig: OperatorNamespaces:[elastic-system] ResourcesNamespaces:[default kube-system] OutputDir: RunStackDiagnostics:true Verbose:false}"
	split := strings.Split(str, " ")

	operatorns := string(split[9]) // OPERATOR NAMESPACE
	slipop := strings.Split(operatorns, ":")
	triml := strings.TrimLeft(string(slipop[1]), "[")
	trimr := strings.TrimRight(triml, "]")

	return trimr

}

func FindRsNS() string {
	str := "2022/03/10 07:44:23 ECK diagnostics with parameters: {DiagnosticImage:docker.elastic.co/eck-dev/support-diagnostics:8.1.4 ECKVersion: Kubeconfig: OperatorNamespaces:[elastic-system] ResourcesNamespaces:[default kube-system] OutputDir: RunStackDiagnostics:true Verbose:false}"
	split := strings.Split(str, ":")

	rsns := split[9] // RESOURCES NAMESPACE
	rm := strings.TrimSuffix(rsns, " OutputDir")
	triml := strings.TrimLeft(rm, "[")
	trimr := strings.TrimRight(triml, "]")

	return trimr

}

// RETURN WORKS ONLY WHEN I CALL IT USING fmt.PrintLN and the function name (see main.go)
func ListDir(d string) []string {
	var files []string
	//var file string
	err := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(files)

	//return files, err
	return files

}
*/
