package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/marcusprice/cli-chat-server/user"
)

var u *user.User

func main() {
	u = user.InitUser()

	port := flag.String("port", "8080", "port for server")
	flag.Parse()

	http.HandleFunc("/connect", Connect)

	log.Println("listening on port " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func Connect(w http.ResponseWriter, r *http.Request) {
	un := r.URL.Query().Get("username")
	fn := r.URL.Query().Get("firstName")
	ln := r.URL.Query().Get("username")

	if u.UserExists(un) {
		http.Error(w, "user is already signed in", http.StatusBadRequest)
		return
	}

	log.Println("connection made for user: ", un)

	u.AddUser(user.UserData{
		UserName:  un,
		FirstName: fn,
		LastName:  ln,
	})

	users := u.GetListOfUsers(un)
	fmt.Fprintln(w, users)
}
