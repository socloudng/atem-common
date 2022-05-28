package utils

import (
	"fmt"
	"strings"
)

func SliceString(objFiles []string, slice string) []string {
	for i := 0; i < len(objFiles); i++ {
		if objFiles[i] == slice {
			objFiles = append(objFiles[0:i], objFiles[i+1:]...)
			i--
		}
	}
	return objFiles
}

func RemoveAt(str string, at int, size int) string {
	fmt.Print("1============" + str + "============\n")
	if len(str) >= at {
		str = str[0:at] + str[at+size:]
	}
	fmt.Print("2============" + str + "============\n")
	return str
}

func GetExtension(str string) string {
	var arr = strings.Split(str, ".")
	if len(arr) > 1 {
		var ext = arr[1]
		var hasCall = strings.Index(str, "?")
		if hasCall != -1 {
			return ext[0:hasCall]
		} else {
			return ext
		}
	}
	return ""
}

func GetFilename(str string, withExt bool) string {
	var start = 0
	if strings.Contains(str, "\\") {
		start = strings.LastIndex(str, "\\")
	} else {
		start = strings.LastIndex(str, "/")
	}
	if start != -1 {
		str = str[start:]
		if strings.Index(str, "\\") == 0 || strings.Index(str, "/") == 0 {
			str = str[1:]
		}
	}
	if withExt {
		return str
	} else {
		return strings.Split(str, ".")[0]
	}
}

func GetPath(str string) string {
	end := 0
	if strings.Contains(str, "\\") {
		end = strings.LastIndex(str, "\\")
	} else {
		end = strings.LastIndex(str, "/")
	}
	if end == -1 {
		return str
	} else {
		return str[0:end]
	}
}

func AppendPath(path string, append string) string {
	return strings.TrimRight(path, "/") + "/" + strings.TrimRight(strings.TrimLeft(append, "/"), "/")
}

func AppendUrl(str string, url string) string {
	result := ""
	fileName := GetFilename(str, true)
	if strings.HasPrefix(fileName, "?") {
		result = AppendPath(GetPath(str), url) + fileName
	} else {
		result = AppendPath(str, url)
	}
	return result
}
