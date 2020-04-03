package file

import (
	"testing"
)
//创建文件
func TestCreateFile(t *testing.T) {
	err := CreateFile("./a.txt")
	if err != nil {
		t.Fatal(err)
	}
}
//覆盖写文件
func TestWriteFile(t *testing.T) {
	err := WriteFile("./a.txt", "","你好")
	if err !=nil{
		t.Fatal(err)
	}
}
//追加写文件
func TestWriteAppendFile(t *testing.T){
	err := WriteFile("./a.txt", "append","你好\n")
	if err !=nil{
		t.Fatal(err)
	}
}
//读取文件
func TestReadFile(t *testing.T) {
	content ,err := ReadFile("./a.txt")
	if err !=nil{
		t.Fatal(err)
	}
	t.Log(string(content))
}

