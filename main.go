package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/wiki", func(w http.ResponseWriter, r *http.Request) {
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
			reqEnc := base64.StdEncoding.EncodeToString([]byte(s.Text))

			response := &slack.Msg{Text: base_url + , ResponseType: "in_channel"}
			resUrl, err := json.Marshal(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(resUrl)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server listening...")
	http.ListenAndServe(":3000", nil)
}
