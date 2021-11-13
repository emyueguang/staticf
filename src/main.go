package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.Printf("%d,  %+v", len(os.Args), os.Args)
	if argsNum := len(os.Args); argsNum > 2 {
		absPath := os.Args[1]
		if !filepath.IsAbs(absPath) {
			log.Println("error! param1 is not abs-path!")
			return
		}

		addr := os.Args[2]
		log.Printf("static file server. %s  %s", absPath, addr)
		http.ListenAndServe(addr, http.FileServer(http.Dir(absPath)))
	} else {
		log.Println("error! no params!  ./staticf  ABS_PATH  ADDR")
	}
}
