package telegrambot

import "testing"

func TestHandler_FileTypeVerify(t *testing.T) {
	h := Handler{
		Bot:   nil,
		Owner: User{},
	}
	_, isValidFileType := h.FileTypeVerify(".xls", []string{".csv", ".xlsx", ".xls"})

	if !isValidFileType {
		t.Errorf("FAILED, expectes %t, go %t", true, false)
		return
	}
	t.Logf("PASSED")
}
