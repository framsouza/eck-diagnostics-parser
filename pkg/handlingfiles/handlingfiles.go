package findfile

import (
	"os"
	"path/filepath"
)

var (
	files           []string
	absPathNodes    []string
	absPathES       []string
	absPathKB       []string
	absPathSTS      []string
	absPathPVC      []string
	absPathPods     []string
	absPathDeploy   []string
	absPathService  []string
	absPathEndpoint []string
)

func FindFileAbsPathPods(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		// path is the absolute path.
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				// if the filename is found append the path to the string slice.
				absPathPods = append(absPathPods, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathPods, nil
}

func FindFileAbsPathDeploy(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				absPathDeploy = append(absPathDeploy, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathDeploy, nil
}

func FindFileAbsPathNodes(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				absPathNodes = append(absPathNodes, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathNodes, nil
}

func FindFileAbsPathServices(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				absPathService = append(absPathService, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathService, nil
}

func FindFileAbsPathEndpoint(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				absPathEndpoint = append(absPathEndpoint, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathEndpoint, nil
}

func FindFileAbsPathES(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				absPathES = append(absPathES, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathES, nil
}

func FindFileAbsPathKB(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				absPathKB = append(absPathKB, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathKB, nil
}

func FindFileAbsPathPVC(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				absPathPVC = append(absPathPVC, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathPVC, nil
}

func FindFileAbsPathSTS(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
				absPathSTS = append(absPathSTS, path)
			}
			break
		}

		return nil
	})
	if fileErr != nil {
		return []string{}, fileErr
	}
	return absPathSTS, nil
}

/* TRY TO MAKE IT WORK AS A GLOBAL FUNC FOR ALL THE FILES
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
/*func ListDir(d string) []string {
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
