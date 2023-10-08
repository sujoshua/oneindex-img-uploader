package oneindex

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"oneindex-img-uploader/pkg/fileutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var parsePattern = regexp.MustCompile(`<label.*>下载地址</label>\n*\t*\s*<input.*value="((http(s)?://)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}[-a-zA-Z0-9()@:%_\+.~#?&/=]*)"/>`)

type oneIndexUploader struct {
	url    string // oneindex url
	rs     *ResultSet
	client *http.Client
}

func newOneIndexUploader(url string, n int) *oneIndexUploader {
	return &oneIndexUploader{
		url:    url,
		client: &http.Client{Timeout: 15 * time.Second},
		rs:     newResultSet(n),
	}
}

func (u *oneIndexUploader) newTask() *oneIndexUploaderTask {
	return &oneIndexUploaderTask{
		oneIndexUploader: u,
	}
}

// Upload upload images to server
type oneIndexUploaderTask struct {
	oneIndexUploader *oneIndexUploader
	req              *http.Request
	body             []byte
	err              error
	rs               string // result url
}

// fromFile set task request from file
func (t *oneIndexUploaderTask) fromFile(path string) *oneIndexUploaderTask {
	if t.err != nil {
		return t
	}

	var file *os.File
	var err error
	var payload = &bytes.Buffer{}
	if file, err = os.OpenFile(path, os.O_RDONLY, 0644); err != nil {
		t.err = err
		return t
	}
	defer file.Close()

	var filepart io.Writer
	mulWriter := multipart.NewWriter(payload)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			"file", filepath.Base(file.Name())))
	h.Set("Content-Type", fileutil.GetContentType(path))
	if filepart, err = mulWriter.CreatePart(h); err != nil {
		t.err = err
		return t
	}
	if _, err = io.Copy(filepart, file); err != nil {
		t.err = err
		return t
	}

	if err = mulWriter.Close(); err != nil {
		t.err = err
		return t
	}

	if t.req, err = http.NewRequest("POST", t.oneIndexUploader.url, payload); err != nil {
		t.err = err
		return t
	}
	t.req.Header.Set("Content-Type", mulWriter.FormDataContentType())
	return t
}

// do send request and get response
func (t *oneIndexUploaderTask) do() (tt *oneIndexUploaderTask) {
	tt = t
	if tt.err != nil {
		return
	}
	var resp *http.Response
	if resp, tt.err = tt.oneIndexUploader.client.Do(tt.req); tt.err != nil {
		return
	}
	defer resp.Body.Close()
	if tt.body, tt.err = io.ReadAll(resp.Body); tt.err != nil {
		return
	}
	return
}

// parse parse response body to get Result url
func (t *oneIndexUploaderTask) parse() *oneIndexUploaderTask {
	if t.err != nil {
		return t
	}
	u := parsePattern.FindSubmatch(t.body)
	if u == nil || len(u) < 2 {
		t.err = fmt.Errorf("parse error, body: %s", string(t.body))
	} else {
		t.rs = string(u[1])
	}
	return t
}

// Result get Result url, or error
func (t *oneIndexUploaderTask) result() (url string, err error) {
	if t.err != nil {
		return "", t.err
	}
	return t.rs, nil
}
