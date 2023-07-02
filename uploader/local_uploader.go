package uploader

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"time"

	"github.com/lampnick/doctron/conf"
)

type LocalUploader struct {
	DoctronUploader
}

func (ins *LocalUploader) Upload() (url string, err error) {
	if ins.Key == "" {
		ins.Key = time.Now().Format("20060102-150405")
	}
	start := time.Now()
	defer func() {
		ins.uploadElapsed = time.Since(start)
	}()

	filename := fmt.Sprintf("%s/%s", conf.LocalStoreConfig.Dir, ins.Key)
	err = ioutil.WriteFile(filename, ins.Stream, fs.ModePerm)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (ins *LocalUploader) GetUploadElapsed() time.Duration {
	return ins.uploadElapsed
}
