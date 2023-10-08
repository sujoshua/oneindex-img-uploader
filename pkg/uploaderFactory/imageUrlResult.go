package uploaderFactory

type ImageUrlResult interface {
	Len() int
	All() []Result
	AllFail() []Result
	Get(task string) Result
	Add(task string, url string, err error)
	Success() int
	Failed() int
}
