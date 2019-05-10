package main

import (
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	UPLOAD_DIR   = "C:/Users/ZOOL/Desktop/uploads"
	TEMPLATE_DIR = "photoweb/views"
)

var templates = make(map[string]*template.Template)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		renderHtml(w, "upload", nil)
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
		check(err)

		defer t.Close()
		_, errCopy := io.Copy(t, file)
		check(errCopy)
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
	check(err)

	locals := make(map[string]interface{})
	var images []string
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "list", locals)
}

//渲染Html模板
func renderHtml(w http.ResponseWriter, htmlName string, data interface{}) error {
	return templates[htmlName].Execute(w, data)
}

func init() {

	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}

	var templateName, templatePath string

	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		//path.Ext获取文件的扩展名
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("正在加载Template模板", templatePath)
		//确保了模板不能解析成功时，一定会触发错误处理流程
		t := template.Must(template.ParseFiles(templatePath))
		newName := strings.Replace(templateName, ".html", "", -1)
		templates[newName] = t
	}
}

//处理异常
func check(err error) {
	if err != nil {
		panic(err)
	}
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
