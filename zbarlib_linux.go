package gozbarlib

import (
	"fmt"
	"github.com/clsung/grcode"
	"strings"
)

/*
- `sudo apt install libzbar-dev zbar-tools -f -y`
*/

func (qr *QRTool) qrCoderReader(imagePath string) (string, error) {
	results, err := grcode.GetDataFromFile(imagePath)
	if err != nil {
		return "", err
	}
	if len(results) == 0 {
		return "", fmt.Errorf("No qrcode detected from file: %s", imagePath)
	}
	
	return strings.Join(results, ""), nil
}
