package common

import "os"

func CloseFile(file *os.File) {
	err := file.Close()
	Check(err)
}
