package exp_dcmiddle

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi,you")
}

func tracing(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tracing start")
		log.Printf("url and method,%s,%s", r.URL, r.Method)
		h.ServeHTTP(w, r)
		fmt.Println("tracing end")
	})
}
func logging(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logging start")
		log.Printf("url and method,%s", r.RemoteAddr)
		n.ServeHTTP(w, r)
		fmt.Println("logging end")
	})
}

func timeRecording(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("timeRecording start")
		startTime := time.Now()
		n.ServeHTTP(w, r)
		elapsedTime := time.Since(startTime)
		log.Printf("执行时间,%s", elapsedTime)
		fmt.Println("timeRecording end")
	})
}

func DcmiddleMain() {
	http.Handle("/", tracing(logging(timeRecording(http.HandlerFunc(hi)))))
	http.ListenAndServe(":8089", nil)

}
