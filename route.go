package restweb

import (
	"net/http"
	"reflect"
	"strings"
)

// GET（SELECT）：从服务器取出资源（一项或多项）。
// POST（CREATE）：在服务器新建一个资源。
// PUT（UPDATE）：在服务器更新资源（客户端提供改变后的完整资源）。
// PATCH（UPDATE）：在服务器更新资源（客户端提供改变的属性）。
// DELETE（DELETE）：从服务器删除资源。
// HEAD：获取资源的元数据。
// OPTIONS：获取信息，关于资源的哪些属性是客户端可以改变的。

type Router interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Head(w http.ResponseWriter, r *http.Request)
	Options(w http.ResponseWriter, r *http.Request)
}

func CallMethod(c interface{}, m string, rv []reflect.Value) {
	rc := reflect.ValueOf(c)
	rm := rc.MethodByName(m)
	rm.Call(rv)
}

func GetReflectValue(w http.ResponseWriter, r *http.Request) (rv []reflect.Value) {
	rw := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	rv = []reflect.Value{rw, rr}
	return
}

var RouterMap = map[string]Router{}

//添加路由
func AddRouter(pattern string, router Router) {
	RouterMap[pattern] = router
}

var FileMap = map[string]http.Handler{}

//添加静态文件路由
func AddFile(pattern string, fileHandler http.Handler) {
	FileMap[pattern] = fileHandler
}

type Server struct {
}

//路由，先处理静态文件，后处理控件，按照最大匹配原则匹配路由
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path + "/"
	// Logger.Debug(path)
	filemaxlenth := 0
	var realFileHandler http.Handler
	for pattern, fileHandler := range FileMap {
		if len(pattern) > filemaxlenth && strings.HasPrefix(path, pattern) {
			filemaxlenth = len(pattern)
			realFileHandler = fileHandler
		}
	}

	maxlenth := 0
	var realRouter Router
	for pattern, router := range RouterMap {
		if len(pattern) > maxlenth && strings.HasPrefix(path, pattern) {
			// Logger.Debug(pattern)
			maxlenth = len(pattern)
			realRouter = router
		}
	}

	if filemaxlenth > maxlenth {
		realFileHandler.ServeHTTP(w, r)
	} else if maxlenth > 0 {
		switch r.Method {
		case "POST":
			realRouter.Post(w, r)
		case "GET":
			realRouter.Get(w, r)
		case "PUT":
			realRouter.Put(w, r)
		case "DELETE":
			realRouter.Delete(w, r)
		case "PATCH":
			realRouter.Patch(w, r)
		case "HEAD":
			realRouter.Head(w, r)
		case "OPTIONS":
			realRouter.Options(w, r)

		}
	} else {
		http.Error(w, "no such page", 404)
	}
}

// 运行服务器
func Run() error {
	return http.ListenAndServe(":8080", &Server{})
}
