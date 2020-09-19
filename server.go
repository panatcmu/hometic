

package main
import (
   "fmt"
   "github.com/gorilla/mux"
   "log"
   "net/http"
   "os"
)
// git init, git add server.go go.mod, git commit -m "[Nong] init project"
func main(){


   fmt.Println("hello hometic : I'm Gopher!!")
   r := mux.NewRouter()
   r.HandleFunc("/pair-device", PairDeviceHandler).Methods(http.MethodPost)

   addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))


   server := http.Server{
	//   Addr:    "127.0.0.1:2009",
	  Addr:    addr,
      Handler: r,
   }
   log.Println("starting...")
   log.Fatal(server.ListenAndServe())
	}

   func PairDeviceHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(`{"status":"active"}`))
   
}