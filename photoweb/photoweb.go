package main

import (
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const UPLOAD_DIR = "C:/Users/ZOOL/Desktop/uploads"

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		renderHtml(w, "photoweb/uploads/upload.html", nil)
		return
	}

	if r.Method == "POST" {

		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, errors.New("请上传图片").Error(), http.StatusForbidden)
			return
		}
		fileName := header.Filename
		defer file.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + fileName)
		if err != nil {
			http.Error(w, "图片上传失败", http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, file); err != nil {
			http.Error(w, "图片存储失败", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+fileName, http.StatusFound)
	}

}

func viewHandle(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + fileName
	if !isExists(imagePath) {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, e := os.Stat(path)
	if e == nil {
		return true
	}
	return os.IsExist(e)
}

func listHandle(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locals := make(map[string]interface{})
	var images []string
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "photoweb/uploads/list.html", locals)
}

//渲染Html模板
func renderHtml(w http.ResponseWriter, htmlPath string, data interface{}) {
	t, err := template.ParseFiles(htmlPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/uploads", uploadHandler)
	http.HandleFunc("/view", viewHandle)
	http.HandleFunc("/", listHandle)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("端口监听异常", err.Error())
	}
}
