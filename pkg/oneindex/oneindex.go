package oneindex

import (
	"fmt"
	"oneindex-img-uploader/pkg/uploaderFactory"
)

type OneIndex struct {
	Url      string // oneindex server index
	Username string // oneindex user
	Password string // oneindex password
}

type resultSet struct {
	path string
	err  error
	url  string
}

func NewUploader(url, username, password string) (*OneIndex, error) {
	if url == "" {
		return nil, fmt.Errorf("url is empty")
	}

	return &OneIndex{
		Url:      url,
		Username: username,
		Password: password,
	}, nil
}

func (o *OneIndex) Upload(taskPaths []string) uploaderFactory.ImageUrlResult {
	uploader := newOneIndexUploader(o.Url, len(taskPaths))
	ec := make(chan resultSet)
	defer close(ec)

	taskFunc := func(path string) {
		if result, err := uploader.newTask().fromFile(path).do().parse().result(); err != nil {
			ec <- resultSet{path: path, err: err}
		} else {
			ec <- resultSet{path: path, url: result}
		}
	}

	for _, path := range taskPaths {
		go taskFunc(path)
	}

	c := 0
	for r := range ec {
		uploader.rs.Add(r.path, r.url, r.err)
		c++
		if c >= len(taskPaths) {
			break
		}
	}
	return uploader.rs
}
