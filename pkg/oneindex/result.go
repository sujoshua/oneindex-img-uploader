package oneindex

import common "oneindex-img-uploader/pkg/uploaderFactory"

type ResultSet struct {
	failedTask  int
	successTask int
	tasks       map[string]common.Result
}

func newResultSet(n int) *ResultSet {
	return &ResultSet{
		failedTask:  0,
		successTask: 0,
		tasks:       make(map[string]common.Result),
	}
}

func (r *ResultSet) Len() int {
	return len(r.tasks)
}

// All return all successful task
func (r *ResultSet) All() []common.Result {
	kp := make([]common.Result, 0, len(r.tasks))
	for _, v := range r.tasks {
		if v.Err == nil {
			kp = append(kp, v)
		}
	}
	return kp
}

func (r *ResultSet) AllFail() []common.Result {
	kp := make([]common.Result, 0, len(r.tasks))
	for _, v := range r.tasks {
		if v.Err != nil {
			kp = append(kp, v)
		}
	}
	return kp
}

func (r *ResultSet) Get(task string) common.Result {
	return r.tasks[task]
}

func (r *ResultSet) Add(task string, url string, err error) {
	if err != nil {
		r.failedTask++
	} else {
		r.successTask++
	}
	r.tasks[task] = common.Result{
		Err: err,
		Url: url,
	}
}

func (r *ResultSet) Adds(results []struct {
	task, url string
	err       error
}) {
	for _, v := range results {
		r.Add(v.task, v.url, v.err)
	}
}

func (r *ResultSet) Success() int {
	return r.successTask
}

func (r *ResultSet) Failed() int {
	return r.failedTask
}
