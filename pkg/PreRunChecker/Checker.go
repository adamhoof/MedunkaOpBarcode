package prerunchecker

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
)

func CreateUpdateDirIfNotExists(dirToCreate string, mode fs.FileMode) (err error) {
	return os.MkdirAll(dirToCreate, mode)
}

func RequestXlsx2CsvInstallationIfNotExists() (err error) {
	out, err := exec.Command("which", "xlsx2csv").Output()
	if err != nil {
		return fmt.Errorf("failed to execute command %s", err)
	}
	if len(out) == 0 {
		return fmt.Errorf("tool is missing %s", err)
	}
	return err
}
