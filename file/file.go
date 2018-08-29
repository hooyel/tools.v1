package file

import (
	"os"
	"strings"
)

//追加写入文件（自带换行）
func AppendFile(strContent string, fileName string) {
	fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	content := strings.Join([]string{strContent, "\n"}, "")
	buf := []byte(content)
	fd.Write(buf)
	fd.Close()
}
