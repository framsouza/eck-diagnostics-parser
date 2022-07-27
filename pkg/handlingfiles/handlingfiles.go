package findfile

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	absPathES       []string
	absPathKB       []string
	absPathSTS      []string
	absPathPVC      []string
	absPathPods     []string
	absPathDeploy   []string
	absPathService  []string
	absPathEndpoint []string
	DeployPath      []string
)

func FindFileAbsPathPods(d, filename string) ([]string, error) {
	fileErr := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for {
			if info.Name() == filename {
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

// TRY TO MAKE IT WORK AS A GLOBAL FUNC FOR ALL THE FILES
func FindFile(filename, d string) []string {
	filepath.Walk(d, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(filename, f.Name())
			if err == nil && r {
				absolutefilepath, err := filepath.Abs(path)
				if err != nil {
					log.Fatal(err)
				}

				DeployPath = append(DeployPath, absolutefilepath)

			}

		}
		return nil
	})
	return DeployPath

}
