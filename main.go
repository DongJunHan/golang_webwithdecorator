package main

import(
	"net/http"
	"time"
	"log"
	"WEB-INF/golang_webwithdecorator/myapp"
	"WEB-INF/golang_webwithdecorator/decoHandler"

)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler){
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w,r)
	log.Println("[LOGGER1] Completed time: ",time.Since(start).Milliseconds())
}

func NewHandler() http.Handler{
	mux := myapp.NewHandler()
	h := decoHandler.NewDecoHandler(mux,logger)	
	return h

}

func main(){
	mux := NewHandler()

	http.ListenAndServe(":3000",mux)
}
