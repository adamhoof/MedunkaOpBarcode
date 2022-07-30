package telegrambot

import (
	"os/exec"
	"testing"
)

func TestExecuteXlsxToCsv(t *testing.T) {
	fileTypeToInsert := ".xlsx"
	output, err := exec.Command("xlsx2csv", "/tmp/Products/sklad"+fileTypeToInsert, "/tmp/Products/update.csv", "-d", ";").Output()
	if err != nil {
		t.Errorf("FAILED %s", err)
		return
	}
	t.Logf("PASSED %s", output)
}
