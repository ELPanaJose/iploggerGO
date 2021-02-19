package main
import (
	"fmt"
	"net/http"
	"log"
)
func main() {
	fmt.Println("Use ngrok! `ngrok http 3000`")
	fmt.Println("server on port 3000")
	http.HandleFunc("/", getIP)
	http.ListenAndServe(":3000", nil)
}

func getIP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("x-forwarded-for")
	log.Println("new ip : ", ip)
	http.Redirect(w, r, "https://elpanajose.ga", 301)
}
