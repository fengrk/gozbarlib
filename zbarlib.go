package gozbarlib

import "os/exec"

var DefaultQRTool = &QRTool{libPath: ""}

type QRTool struct {
	libPath string
}

func (qr *QRTool) QRCodeReader(imagePath string) (string, error) {
	return qr.qrCoderReader(imagePath)
}

func (qr *QRTool) SetLibPath(libPath string) {
	qr.libPath = libPath
}

func (qr *QRTool) FindLibPath() (string, error) {
	zbarName := "zbarimg"
	return exec.LookPath(zbarName)
}

func QRCodeReader(imagePath string) (string, error) {
	return DefaultQRTool.QRCodeReader(imagePath)
}
