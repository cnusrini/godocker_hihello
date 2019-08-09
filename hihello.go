package main

import (
  "fmt"
  "log"
  "net/http"
  "io"
)

func Hello(w http.ResponseWriter,r *http.Request){
  w.Header().Set("content-type","text/plain;charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("hello in w write" + "\n"))
  log.Println("hello func executed")
}
func goodbye(res http.ResponseWriter, req *http.Request){
res.Header().Set("content-type","text/plain;chatset=utf-8")
res.WriteHeader(http.StatusOK)
io.WriteString(res , "googbye" + "\n")
log.Println("executed goodbye func")

}
func main(){
  fmt.Println("in hihello go fun")
  http.HandleFunc("/", Hello)
  http.HandleFunc("/hi",goodbye)
  log.Fatal(http.ListenAndServe(":8000",nil))
}
