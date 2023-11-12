package utils

import (
	"errors"
	"github.com/huoxue1/go-utils/base/log"
	"os"
	"path/filepath"
	"regexp"
)

// ExistDir 判断目录是否存在
func ExistDir(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return true
	}
}

// WorkPath 工作路径
func WorkPath(path ...string) (work string) {
	work, _ = os.Getwd()
	joinPath := filepath.Join(path...)
	return filepath.Join(work, joinPath)
}

// ImgPath 图片存储路径
func ImgPath(fileName string) string {
	return filepath.Join(
		CreateDir("cache", "img"),
		fileName,
	)
}

// GenPath 生成路径
func GenPath(path string, fileName string) string {
	if fileName == "" {
		return CreateDir("cache", "gen", path)
	}

	return filepath.Join(
		CreateDir("cache", "gen", path),
		fileName,
	)
}

// CreateDir 创建文件夹
func CreateDir(path ...string) (dir string) {
	workPath := WorkPath(path...)
	if !ExistDir(workPath) {
		err := os.MkdirAll(workPath, 0755)
		if err != nil {
			log.Errorf("创建目录失败: %s", err)
			return ""
		}
	}
	return workPath
}

func ReadFilesWithCallback(directory string, queryFilename string, callback func(filePath string) error) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	extractWeek := func(str string) string {
		re := regexp.MustCompile(`(第\d+周)`)
		match := re.FindStringSubmatch(str)
		if len(match) > 1 {
			return match[0]
		}
		return ""
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			filePath := filepath.Join(directory, filename)

			if extractWeek(filename) == queryFilename {
				return callback(filePath)
			}
		}
	}
	return errors.New("未找到")
}
