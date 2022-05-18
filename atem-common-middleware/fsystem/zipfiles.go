package fsystem

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ZipFiles(filename string, files []string, oldForm, newForm string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = newZipFile.Close()
	}()

	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		_ = zipWriter.Close()
	}()

	// 把files添加到zip中
	for _, file := range files {

		err = func(file string) error {
			zipFile, err := os.Open(file)
			if err != nil {
				return err
			}
			defer zipFile.Close()
			// 获取file的基础信息
			info, err := zipFile.Stat()
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			// 使用上面的FileInforHeader() 就可以把文件保存的路径替换成我们自己想要的了，如下面
			header.Name = strings.Replace(file, oldForm, newForm, -1)

			// 优化压缩
			// 更多参考see http://golang.org/pkg/archive/zip/#pkg-constants
			header.Method = zip.Deflate

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			if _, err = io.Copy(writer, zipFile); err != nil {
				return err
			}
			return nil
		}(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// Zip reads all the files in the directory and creates a new compressed file.
func Zip(dir, path string) error {
	// check dir
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return fmt.Errorf("%v is not exist", dir)
	}
	if !stat.IsDir() {
		return fmt.Errorf("%v is not a dir", dir)
	}

	// check path
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// create zip file
	w := zip.NewWriter(file)
	defer w.Close()

	// func to create zip file
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == dir { // The first path of filepath.Walk is always the root.
			return nil
		}
		targetPath := strings.TrimLeft(path, dir)
		if info.IsDir() { // dir
			if _, err := w.Create(targetPath + "/"); err != nil {
				return err
			}
		} else { // file
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			target, err := w.Create(targetPath)
			if err != nil {
				return err
			}
			if _, err := target.Write(bytes); err != nil {
				return err
			}
		}

		return nil
	}

	// walk all the directories and files in the dir
	if err := filepath.Walk(dir, walkFunc); err != nil {
		return err
	}

	return nil
}

// UnZip decompresses all files to a directory.
func UnZip(filename, dir string) error {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		name := f.Name
		info := f.FileInfo()
		targetPath := filepath.Join(dir, name)
		if info.IsDir() { // dir
			if _, err := os.Stat(targetPath); os.IsNotExist(err) {
				// IMPORTANT: info.Mode() return wrong mode when f is a dir.
				// So, DO NOT use info.Mode() to os.MkdirAll.
				if err := os.MkdirAll(targetPath, 0755); err != nil {
					return err
				}
			}
		} else { // file
			reader, err := f.Open()
			if err != nil {
				return err
			}
			defer reader.Close()
			writer, err := os.OpenFile(targetPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, info.Mode())
			if err != nil {
				return err
			}
			defer writer.Close()
			if _, err := io.Copy(writer, reader); err != nil {
				return err
			}
		}
	}

	return nil
}
