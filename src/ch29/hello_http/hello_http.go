/*
 * @Description: http服务
 * @Author: leo
 * @Date: 2020-02-14 17:26:36
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-14 18:02:38
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello World")
	})
	http.HandleFunc("/time", func(res http.ResponseWriter, req *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		res.Write([]byte(timeStr))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), logRequest(http.DefaultServeMux))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
