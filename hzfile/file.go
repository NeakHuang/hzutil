package hzfile

import (
	"fmt"
	"os"
)

func GetContents() ([]byte, error) {
	data, err := os.ReadFile("./util/file.go")
	return data, err
}

func WriteFile(folderDir, filename string, content []byte) error {
	// 获取文件信息
	if _, err := os.Stat(folderDir); os.IsNotExist(err) {
		// 创建文件所在目录路径
		if err = os.MkdirAll(folderDir, os.ModePerm); err != nil {
			return err
		}
	}

	path := fmt.Sprintf("%v/%v", folderDir, filename)
	if err := os.WriteFile(path, content, 0644); err != nil {
		return err
	}
	return nil
}

func MkDir(dir string) error {
	// 获取文件信息
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 创建文件所在目录路径
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func FilenameList(folderDir string) ([]string, error) {
	// 获取文件夹中的所有文件
	entries, err := os.ReadDir(folderDir)
	if err != nil {
		return nil, err
	}
	// infos := make([]fs.FileInfo, 0, len(entries))
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if info.IsDir() {
			continue
		}
		names = append(names, info.Name())
	}
	return names, nil
}
