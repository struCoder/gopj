package app

import (
	// "bytes"
	//"code.google.com/p/go.net/websocket"
	"crypto/hmac"
	"crypto/sha1"
	// "crypto/tls"
	// "encoding/base64"
	"fmt"
	// "io/ioutil"
	// "log"
	"mime"
	"net/http"
	// "os"
	// "path"
	// "reflect"
	// "strconv"
	"strings"
	// "time"
)

type Context struct {
	Request *http.Request
	Params  map[string]string
	Server  *Server
	http.ResponseWriter
}

func (ctx *Context) WriteString(content string) {
	ctx.ResponseWriter.write([]byte(content))
}

func (ctx *Context) Abort(status int, body string) {
	ctx.ResponseWriter.WriteHeader(status)
	ctx.ResponseWriter.Write([]byte(body))
}

func (ctx *Context) Redirect(status int, url_ string) {
	ctx.ResponseWriter.Header().Set("Location", url_)
	ctx.ResponseWriter.WriteHeader(status)
	ctx.ResponseWriter.Write([]byte("跳转到: " + url_))
}
func (ctx *Context) NotFound(message string) {
	ctx.ResponseWriter.WriteHeader(404)
	ctx.ResponseWriter.Write([]byte(message))
}

func (ctx *Context) Unauthorized() {
	ctx.ResponseWriter.WriteHeader(401)
}

func (ctx *Context) Forbidden() {
	ctx.ResponseWriter.WriteHeader(403)
}

func (ctx *Context) ContentTye(val string) string {
	var ctype string
	if strings.ContainsRune(val, '/') {
		ctype = val
	} else {
		if !strings.HasPrefix(val, '.') {
			val = "." + val
		}
		ctype = mime.TypeByExtension(val)
	}

	if ctype != "" {
		ctx.Header().Set("Content-Type", ctype)
	}

	return ctype
}

func (ctx *Context) SetHeader(hdr string, val string, unique bool) {
	if unique {
		ctx.Header().Set(hdr, val)
	} else {
		ctx.Header().Add(hdr, val)
	}
}

func (ctx *Context) SetCookie(cookie *http.Cookie) {
	ctx.SetHeader("Set-Cookie", cookie.String(), false)
}

func getCookieSig(key string, val []byte, timestamp string) string {
	hm := hmac.New(sha1.New(), []byte(key))
	hm.Write(val)
	hm.Write([]byte(timestamp))
	hex := fmt.Sprintf("%02x", hm.Sum(nil))
	return hex
}
