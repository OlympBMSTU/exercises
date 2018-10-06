package fstorage

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"math/big"
	"os"

	"github.com/HustonMmmavr/excercieses/config"
)

func FileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ComputeName(filename string) string {
	digest := md5.New()
	digest.Write([]byte(filename))
	hashBytes := digest.Sum(nil)

	z := new(big.Int)
	z.SetBytes(hashBytes)

	// see how correctly do this and there
	newPath := fmt.Sprintf("%x", z)
	newPath = newPath[:2] + "/" + newPath[2:4] + "/" + newPath[4:]
	return newPath
}

func WriteFile(fileData []byte, name string, ext string) error {
	dirs := name[:6]
	conf, _ := config.GetConfigInstance()

	dirPath := conf.GetFileStorageName() + dirs
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		// clear dirs
		return err
	}

	filePath := conf.GetFileStorageName() + name
	filePathWithExt := filePath + ext
	idx := 1
	for {
		if FileExist(filePathWithExt) {
			filePath += string(idx)
			filePathWithExt = filePath + ext
		} else {
			break
		}
		idx++
	}

	f, err := os.Create(filePathWithExt)
	defer f.Close()
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)
	_, err = writer.Write(fileData)

	if err != nil {
		return err
	}

	writer.Flush()

	return nil
}
