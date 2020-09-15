package files

import "os"

func RemoveFile(path1 string, path2 string) error {

	return os.Rename(path1, path2)
}
