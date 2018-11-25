package fstorage

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/OlympBMSTU/exercises/config"
	"github.com/OlympBMSTU/exercises/fstorage/result"
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

// todo refactor close open work eith strings
func WriteFile(fileHdr *multipart.FileHeader) result.FSResult {
	if fileHdr == nil {
		return result.ErrorResult(errors.New("No file presented"))
	}
	conf, _ := config.GetConfigInstance()

	ext := filepath.Ext(fileHdr.Filename)
	newNamePart := ComputeName(fileHdr.Filename)
	staticPath := conf.GetFileStorageName() + "/"
	newDirsPath := staticPath + newNamePart[:6]

	filePathWithExt := staticPath + newNamePart + ext
	idx := 1
	for {
		if FileExist(filePathWithExt) {
			newNamePart += strconv.Itoa(idx)
			filePathWithExt = staticPath + newNamePart + ext
		} else {
			break
		}
		idx++
	}

	inFile, err := fileHdr.Open()
	defer inFile.Close()

	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}

	err = os.MkdirAll(newDirsPath, 0777)
	if err != nil {
		log.Println(err.Error())
		// clear dirs
		return result.ErrorResult(err)
	}

	f, err := os.Create(filePathWithExt)
	defer f.Close() // ? is it
	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}
	_, err = io.Copy(f, inFile)
	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}
	return result.OkResult(newNamePart + ext)
}
