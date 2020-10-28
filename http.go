package main
import "net/http"
func main(){
http.HandleFunc("/hello-word", func(w http.ResponseWriter,r *http.Request){
w.Write([]byte("hello word"))
})
http.ListenAndServe(":8080",nil)
}
