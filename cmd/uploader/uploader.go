package uploader

import (
	"fmt"
	"io/fs"
	"oneindex-img-uploader/pkg/fileutil"
	"oneindex-img-uploader/pkg/oneindex"
	"oneindex-img-uploader/pkg/uploaderFactory"
	"path/filepath"
)

// InitUploader init uploader by Config.Platform
func InitUploader() (err error) {
	switch Config.Platform {
	case "oneindex":
		Config.Uploader, err = oneindex.NewUploader(Config.Url, "", "")
	default:
		err = fmt.Errorf("platform %s not support", Config.Platform)
	}
	return
}

// CollectValidatePaths collect all validate image file in given path.
// The rule is:
//
//	check if the path is a directory, if it is a directory, walk through all files and sub dir.
//	check if the path is an image file, if it is an image file, then upload it.
//
// Return all validate image file path in input order, and not filter repeated file. If a Path is invalid, then just ignore parse it.
func CollectValidatePaths(paths []string) (validatePaths []string) {
	validatePaths = collectValidatePaths(paths)
	fmt.Printf("collect %d valid img file:\n", len(validatePaths))
	for _, p := range validatePaths {
		fmt.Println(p)
	}
	return
}

// Upload  all validate image file to the platform Config defined.
func Upload(validatePath []string) uploaderFactory.ImageUrlResult {
	if Config.Uploader == nil {
		fmt.Println("collect failed, uploader not initialized")
		return nil
	}
	return Config.Uploader.Upload(validatePath)
}

// PrintResultByOrder print result by order, the order is defined by the input task order in tasks slice.
// If the task is success, then print the url, otherwise print the error.
func PrintResultByOrder(tasks []string, pairs uploaderFactory.ImageUrlResult) {
	if pairs == nil {
		fmt.Println("print upload failed, none result gotten")
	}
	fmt.Println("upload task finished.")
	fmt.Printf("success: %d, failed: %d\n", pairs.Success(), pairs.Failed())
	for _, t := range tasks {
		if pairs.Get(t).Err != nil {
			fmt.Printf("task %s failed: %s\n", t, pairs.Get(t).Err)
		} else {
			fmt.Println(pairs.Get(t).Url)
		}
	}
}

func collectValidatePaths(paths []string) (validatePaths []string) {
	for _, p := range paths {
		p = fileutil.AbsPath(p)
		_ = filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return fs.SkipDir
			}
			if !d.IsDir() && fileutil.IsImg(path) {
				validatePaths = append(validatePaths, path)
			}
			return nil
		})
	}
	return
}
