

package main
import (
   "fmt"
   "github.com/gorilla/mux"
   "log"
   "net/http"
   "os"
   "encoding/json"
)

type Pair struct {
	DeviceID int64
	UserID int64
}
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
	var p Pair
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer r.Body.Close()
	fmt.Printf("pair:%#v\n",p)
	w.Write([]byte(`{"status":"active"}`))
   
}