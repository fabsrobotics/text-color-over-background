package main

import(
	"log"
	"net/http"
	"math/rand"
	"time"
	"github.com/crazy3lf/colorconv"
)

func mainRouter(w http.ResponseWriter, r *http.Request){
	rand.Seed(time.Now().UnixNano())
	h := 360.0*rand.Float64()
	c,err := colorconv.HSVToColor(h,1.0,1.0)
	if err != nil { log.Fatal(err) }
	hex := colorconv.ColorToHex(c)
	color := "#"+hex[2:]
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(color))
}

func main(){
	http.HandleFunc("/",mainRouter)
	err := http.ListenAndServe(":13001",nil)
	if err != nil { log.Fatal(err) }
}