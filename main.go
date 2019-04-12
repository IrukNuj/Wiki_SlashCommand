package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/cmd", func(w http.ResponseWriter, r *http.Request) {
		base_url := "https://ja.wikipedia.org/wiki/"

		s, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !s.ValidateToken(os.Getenv("VERIFICATION_TOKEN")) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch s.Command {
		case "/wiki":
			response := &slack.Msg{Text: base_url + s.Text, ResponseType: "in_channel"}
			print(response)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server listening...")
	http.ListenAndServe(":3000", nil)
}
