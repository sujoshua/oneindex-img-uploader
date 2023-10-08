package uploaderFactory

type Result struct {
	Url string
	Err error
}

type Uploader interface {
	Upload(taskPaths []string) ImageUrlResult // upload images to server, taskPaths is image paths to upload and path must be a valid image file
}
