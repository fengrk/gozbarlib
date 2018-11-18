# gozbarlib: 系统调用zbarimg解析二维码

# 解析二维码
通过系统调用zbarimg, 解析二维码图片

示例:
```
package main

import (
	"fmt"
	"github.com/frkhit/gozbarlib"
)

func main() {
	content, err := gozbarlib.QRCodeReader("xx.png")
	if err!=nil{
		panic(err)
	}
	fmt.Println("content of image is :", content)
}
```

# 安装
## windows系统下安装
- 安装目录下的`zbar-0.10-setup.exe`程序 [最新版本](https://cfhcable.dl.sourceforge.net/project/zbar/zbar/0.10/zbar-0.10-setup.exe)
- 将`zbarimg.exe`所在目录添加至系统`PATH`中
- 安装成功: cmd下执行`zbarimg.exe`, 确定是否安装成功

## linux下安装
- 安装zbar: `sudo apt install libzbar-dev zbar-tools -f -y`
- 安装成功: terminal执行`zbarimg`确定是否安装成功

