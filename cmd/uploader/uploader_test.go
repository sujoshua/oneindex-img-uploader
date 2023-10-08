package uploader

import (
	"os"
	"testing"
)

func TestCollectValidateImg(t *testing.T) {
	validatePaths := collectValidatePaths(os.Args[1:])
	for _, v := range validatePaths {
		t.Log(v)
	}
}
