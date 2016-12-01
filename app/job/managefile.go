package job

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RenameFile(file *os.File, namaFile string) (string, error) {
	file.Close()
	oldPath := file.Name()
	path := strings.Split(oldPath, "/")
	tempPath := strings.Split(path[1], ".")
	fmt.Println(file.Name())
	fmt.Println(path[0] + "/" + namaFile + "." + tempPath[1])
	newPath := path[0] + "/" + namaFile + "." + tempPath[1]
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return "", err
	}
	return namaFile + "." + tempPath[1], nil

}

func CreateDir(root_path string, username string) (bool, error) {
	dst_path := "public/data/" + username
	path := filepath.Join(root_path, dst_path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err2 := os.Mkdir(path, os.ModePerm)
		if err2 != nil {
			return false, err2
		}
		return true, nil
	}
	return false, err
}
