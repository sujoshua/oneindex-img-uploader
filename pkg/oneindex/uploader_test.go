package oneindex

import (
	"testing"
)

func (t *oneIndexUploaderTask) fakeDo(body string) (tt *oneIndexUploaderTask) {
	tt = t
	if tt.err != nil {
		return
	}
	t.body = []byte(body)
	return
}

type testTask struct {
	uploader        *oneIndexUploader
	fakeBody        string
	wantedResultUrl string
}

/*
func TestParse(test *testing.T) {
	tasks := []*testTask{
		{
			uploader:        newOneIndexUploader("https://test.url", 1),
			fakeBody:        "<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset=\"utf-8\">\n\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0,maximum-scale=1.0, user-scalable=no\"/>\n\t<title>Screenshot_20230930_223412.png - Joshua's OneIndex</title>\n\t<link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/mdui@0.4.1/dist/css/mdui.min.css\" integrity=\"sha256-lCFxSSYsY5OMx6y8gp8/j6NVngvBh3ulMtrf4SX5Z5A=\" crossorigin=\"anonymous\">\n\t<script src=\"https://cdn.jsdelivr.net/npm/mdui@0.4.1/dist/js/mdui.min.js\" integrity=\"sha256-dZxrLDxoyEQADIAGrWhPtWqjDFvZZBigzArprSzkKgI=\" crossorigin=\"anonymous\"></script>\n\t<style>\n\t\t.mdui-appbar .mdui-toolbar{\n\t\t\theight:56px;\n\t\t\tfont-size: 16px;\n\t\t}\n\t\t.mdui-toolbar>*{\n\t\t\tpadding: 0 6px;\n\t\t\tmargin: 0 2px;\n\t\t\topacity:0.5;\n\t\t}\n\t\t.mdui-toolbar>.mdui-typo-headline{\n\t\t\tpadding: 0 16px 0 0;\n\t\t}\n\t\t.mdui-toolbar>i{\n\t\t\tpadding: 0;\n\t\t}\n\t\t.mdui-toolbar>a:hover,a.mdui-typo-headline,a.active{\n\t\t\topacity:1;\n\t\t}\n\t\t.mdui-container{\n\t\t\tmax-width:980px;\n\t\t}\n\t\t.mdui-list-item{\n\t\t\t-webkit-transition:none;\n\t\t\ttransition:none;\n\t\t}\n\t\t.mdui-list>.th{\n\t\t\tbackground-color:initial;\n\t\t}\n\t\t.mdui-list-item>a{\n\t\t\twidth:100%;\n\t\t\tline-height: 48px\n\t\t}\n\t\t.mdui-list-item{\n\t\t\tmargin: 2px 0px;\n\t\t\tpadding:0;\n\t\t}\n\t\t.mdui-toolbar>a:last-child{\n\t\t\topacity:1;\n\t\t}\n\t\t@media screen and (max-width:980px){\n\t\t\t.mdui-list-item .mdui-text-right{\n\t\t\t\tdisplay: none;\n\t\t\t}\n\t\t\t.mdui-container{\n\t\t\t\twidth:100% !important;\n\t\t\t\tmargin:0px;\n\t\t\t}\n\t\t\t.mdui-toolbar>*{\n\t\t\t\tdisplay: none;\n\t\t\t}\n\t\t\t.mdui-toolbar>a:last-child,.mdui-toolbar>.mdui-typo-headline,.mdui-toolbar>i:first-child{\n\t\t\t\tdisplay: block;\n\t\t\t}\n\t\t}\n\t</style>\n</head>\n<body class=\"mdui-theme-primary-blue-grey mdui-theme-accent-blue\">\n\t<header class=\"mdui-appbar mdui-color-theme\">\n\t\t<div class=\"mdui-toolbar mdui-container\">\n\t\t\t<a href=\"/\" class=\"mdui-typo-headline\">Joshua's OneIndex</a>\n\t\t\t\t\t\t<i class=\"mdui-icon material-icons mdui-icon-dark\" style=\"margin:0;\">chevron_right</i>\n\t\t\t<a href=\"/\">/</a>\n\t\t\t\t\t\t<i class=\"mdui-icon material-icons mdui-icon-dark\" style=\"margin:0;\">chevron_right</i>\n\t\t\t<a href=\"/images/\">images</a>\n\t\t\t\t\t\t<i class=\"mdui-icon material-icons mdui-icon-dark\" style=\"margin:0;\">chevron_right</i>\n\t\t\t<a href=\"/images/2023/\">2023</a>\n\t\t\t\t\t\t<i class=\"mdui-icon material-icons mdui-icon-dark\" style=\"margin:0;\">chevron_right</i>\n\t\t\t<a href=\"/images/2023/10/\">10</a>\n\t\t\t\t\t\t<i class=\"mdui-icon material-icons mdui-icon-dark\" style=\"margin:0;\">chevron_right</i>\n\t\t\t<a href=\"/images/2023/10/04/\">04</a>\n\t\t\t\t\t\t<i class=\"mdui-icon material-icons mdui-icon-dark\" style=\"margin:0;\">chevron_right</i>\n\t\t\t<a href=\"/images/2023/10/04/Wknz2Waa4i/\">Wknz2Waa4i</a>\n\t\t\t\t\t\t<i class=\"mdui-icon material-icons mdui-icon-dark\" style=\"margin:0;\">chevron_right</i>\n\t\t\t<a href=\"/images/2023/10/04/Wknz2Waa4i/Screenshot_20230930_223412.png\">Screenshot_20230930_223412.png</a>\n\t\t\t\t\t\t<!--<a href=\"javascript:;\" class=\"mdui-btn mdui-btn-icon\"><i class=\"mdui-icon material-icons\">refresh</i></a>-->\n\t\t</div>\n\t</header>\n\t\n\t<div class=\"mdui-container\">\n    \t\t\n<div class=\"mdui-container-fluid\">\n\t<br>\n\t<img class=\"mdui-img-fluid\" src=\"https://suxinyang-my.sharepoint.com/personal/sunshin_sunshin_ml/_layouts/15/download.aspx?UniqueId=b036c7c9-5ad6-46ef-b5c2-b18c232cf0d0&Translate=false&tempauth=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIwMDAwMDAwMy0wMDAwLTBmZjEtY2UwMC0wMDAwMDAwMDAwMDAvc3V4aW55YW5nLW15LnNoYXJlcG9pbnQuY29tQDFkNDVkM2EzLTMzMzItNDYyZC1iMzM5LWM0Njg4Njk2OWYyMCIsImlzcyI6IjAwMDAwMDAzLTAwMDAtMGZmMS1jZTAwLTAwMDAwMDAwMDAwMCIsIm5iZiI6IjE2OTYzODYyNDIiLCJleHAiOiIxNjk2Mzg5ODQyIiwiZW5kcG9pbnR1cmwiOiJ1L1M2b0s5anpnekdIU05yZEtXOGRPdTVVeVRYQXN3MzI3SDVQODJNbXU0PSIsImVuZHBvaW50dXJsTGVuZ3RoIjoiMTUxIiwiaXNsb29wYmFjayI6IlRydWUiLCJjaWQiOiJRNFI5K2xMbWFVZVdKS0U1anhSaGFRPT0iLCJ2ZXIiOiJoYXNoZWRwcm9vZnRva2VuIiwic2l0ZWlkIjoiTWpSa05qSTRNMk10Wm1ZeE1TMDBObU01TFRrNE9EUXRaVEkzT0dNMFptUTJOVFEyIiwiYXBwX2Rpc3BsYXluYW1lIjoib25laW5kZXgiLCJnaXZlbl9uYW1lIjoi6ZGr6ZizIiwiZmFtaWx5X25hbWUiOiLoi48iLCJhcHBpZCI6ImZhMmFkZmU0LWNiNjEtNDgwYy1hYzVhLTU3ZDJjNTYzMDZjNyIsInRpZCI6IjFkNDVkM2EzLTMzMzItNDYyZC1iMzM5LWM0Njg4Njk2OWYyMCIsInVwbiI6InN1bnNoaW5Ac3Vuc2hpbi5tbCIsInB1aWQiOiIxMDAzMjAwMEQ5QjFCNUVGIiwiY2FjaGVrZXkiOiIwaC5mfG1lbWJlcnNoaXB8MTAwMzIwMDBkOWIxYjVlZkBsaXZlLmNvbSIsInNjcCI6Im15ZmlsZXMucmVhZCBhbGxmaWxlcy5yZWFkIG15ZmlsZXMud3JpdGUgYWxsZmlsZXMud3JpdGUgYWxsc2l0ZXMucmVhZCIsInR0IjoiMiIsImlwYWRkciI6IjIwLjE5MC4xNTQuMzIifQ.oAp_MqDhag8zXYOkMNBm-d1kIvCeZ0WXbefDZZrcALI&ApiVersion=2.0\"/>\n\t<br>\n\t<div class=\"mdui-textfield\">\n\t  <label class=\"mdui-textfield-label\">下载地址</label>\n\t  <input class=\"mdui-textfield-input\" type=\"text\" value=\"http://oneindex.joshua.su/images/2023/10/04/Wknz2Waa4i/Screenshot_20230930_223412.png\"/>\n\t</div>\n\t<div class=\"mdui-textfield\">\n\t  <label class=\"mdui-textfield-label\">HTML 引用地址</label>\n\t  <input class=\"mdui-textfield-input\" type=\"text\" value=\"<img src='http://oneindex.joshua.su/images/2023/10/04/Wknz2Waa4i/Screenshot_20230930_223412.png' />\"/>\n\t</div>\n        <div class=\"mdui-textfield\">\n\t  <label class=\"mdui-textfield-label\">Markdown 引用地址</label>\n\t  <input class=\"mdui-textfield-input\" type=\"text\" value=\"![](http://oneindex.joshua.su/images/2023/10/04/Wknz2Waa4i/Screenshot_20230930_223412.png)\"/>\n\t</div>\n        <br>\n</div>\n<a href=\"http://oneindex.joshua.su/images/2023/10/04/Wknz2Waa4i/Screenshot_20230930_223412.png\" class=\"mdui-fab mdui-fab-fixed mdui-ripple mdui-color-theme-accent\"><i class=\"mdui-icon material-icons\">file_download</i></a>\n  \t</div>\n</body>\n</html>\n",
			wantedResultUrl: "http://oneindex.joshua.su/images/2023/10/04/Wknz2Waa4i/Screenshot_20230930_223412.png",
		},
	}
	for _, t := range tasks {
		if resultUrl, err := t.uploader.newTask().fakeDo(t.fakeBody).parse().result(); resultUrl != t.wantedResultUrl || err != nil {
			if err != nil {
				test.Fatalf("parse error: %s, wanted %s, got %s", err.Error(), t.wantedResultUrl, resultUrl)
			} else {
				test.Fatalf("parse error: wanted %s, got %s", t.wantedResultUrl, resultUrl)
			}
		}
	}
}
*/

func TestUpload(t *testing.T) {
	uploader, err := NewUploader("https://oneindex.joshua.su/images", "", "")
	if err != nil {
		t.Fatal(err)
	}
	r := uploader.Upload([]string{"./cilium_logo.png"})
	for _, v := range r.All() {
		t.Log(v.Url, v.Err)
	}
}
