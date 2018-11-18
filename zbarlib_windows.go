package gozbarlib

import (
	"bytes"
	"os/exec"
	"strings"
)

/*
- install zbar-0.10-setup.exe
- Add path including zbarimg.exe to `PATH`
*/

func (qr *QRTool) qrCoderReader(imagePath string) (string, error) {
	if len(qr.libPath) == 0 {
		libPath, libErr := qr.FindLibPath()
		if libErr != nil {
			return "", libErr
		}
		qr.SetLibPath(libPath)
	}
	
	cmd := exec.Command(qr.libPath, "-D", "--raw", imagePath)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdout.Bytes())), nil
}
