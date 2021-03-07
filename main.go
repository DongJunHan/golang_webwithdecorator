package main

import(
	"net/http"
	"WEB-INF/golang_webwithdecorator/myapp"
)

func NewHandler() http.Handler{
	mux := myapp.NewHandler()
	return mux

}

func main(){
	mux := NewHandler()

	http.ListenAndServe(":3000",mux)
}
