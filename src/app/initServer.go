package app

import (
	// "controllers"
	// "fmt"
	"log"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"time"
)

type myHandle struct{}

type route struct {
	r           string
	cr          *regexp.Regexp
	method      string
	httpHandler http.Handler
}
type Server struct {
	// Config *ServerConfig
	routes []route
	Logger *log.Logger
	Env    map[string]interface{}
	l      net.Listener
}

// var mux map[string]func(http.ResponseWriter, *http.Request)

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.Process(res, req)
}

// func InitServer() {
// 	server := http.Server{
// 		Addr:        ":7070",
// 		Handler:     &myHandle{},
// 		ReadTimeout: 6 * time.Second,
// 	}

// 	mux = make(map[string]func(http.ResponseWriter, *http.Request))
// 	mux["/login"] = controllers.HandleLogin
// 	mux["/logout"] = controllers.Logout
// 	mux["/index"] = controllers.Home
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal(err)
// 		fmt.Println("error")
// 	}
// }

func (s *Server) Run(addr string) {
	mux := http.NewServeMux()
	mux.Handle("/", s)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
	s.l = l
	err = http.Server(s.l, mux)
	s.l.Close()
}
func (s *Server) Process(res http.ResponseWriter, req *http.Request) {
	route := s.routeHandler(req, res)
	if route != nil {
		route.httpHandler.ServeHTTP(res, req)
	}
}

func (s *Server) routeHandler(req *http.Request, res http.ResponseWriter) (unused *route) {
	requestPath := req.URL.Path
	ctx := Context{req, map[string]string{}, s, res}
	ctx.SetHeader("Server", "strucoder go Server", true)
	tm := time.Now().UTC()
	req.ParseForm()
	if len(req.Form) > 0 {
		for k, v := range req.Form {
			ctx.Params[k] = v[0]
		}
	}

	ctx.SetHeader("Date", webTime(tm), true)
	ctx.SetHeader("Content-Type", "text/html; charset=utf-8", true)
	for i := 0; i < len(s.routes); i++ {
		route := s.routes[i]
		cr := route.cr
		if req.Method != route.method {
			continue
		}

		if route.httpHandler != nil {
			unused = &route
			return
		}
		return
	}
}

func (s *Server) addRoute(r string, method string, handler interface{}) {
	cr, err := regexp.Compile(r)
	if err != nil {
		s.Logger.Printf("Error in route regex %q\n", r)
		return
	}
	switch handler.(type) {
	case http.Handler:
		s.routes = append(s.routes, route{r: r, cr: cr, method: method, httpHandler: handler.(http.Handler)})
	case reflect.Value:
		fv := handler.(reflect.Value)
		s.routes = append(s.routes, route{r: r, cr: cr, method: method, handler: fv})
	default:
		fv := reflect.ValueOf(handler)
		s.routes = append(s.routes, route{r: r, cr: cr, method: method, handler: fv})
	}
}

func (s *Server) Get(route string, handler interface{}) {
	s.addRoute(route, "GET", handler)
}

func (s *Server) Post(route string, handler interface{}) {
	s.addRoute(route, "POST", handler)
}
