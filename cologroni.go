package cologroni

import (
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/negroni"
)

type logger struct{}

var (
	methodcol   = color.New(color.Bold).SprintFunc()
	pathcol     = color.New(color.Italic).SprintFunc()
	timecol     = pathcol
	infocol     = color.New(color.Bold, color.FgCyan).SprintFunc()
	successcol  = color.New(color.Bold, color.FgGreen).SprintFunc()
	redirectcol = color.New(color.Bold, color.FgYellow).SprintFunc()
	clierrcol   = color.New(color.Bold, color.FgRed).SprintFunc()
	serverrcol  = color.New(color.Bold, color.FgMagenta).SprintFunc()
)

func (l logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, r)
	res := rw.(negroni.ResponseWriter)
	status := res.Status()
	var statstr string
	switch {
	case 200 <= status && status < 300:
		statstr = successcol(status)
	case 300 <= status && status < 400:
		statstr = redirectcol(status)
	case 400 <= status && status < 500:
		statstr = clierrcol(status)
	case 500 <= status:
		statstr = serverrcol(status)
	default:
		statstr = infocol(status)
	}
	log.Printf("%v %v %vB → %v %v %vB",
		methodcol(r.Method), pathcol(r.URL.Path), r.ContentLength,
		statstr, timecol(time.Since(start)), res.Size())
}

func New() logger {
	return logger{}
}
