package server

import (
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second*5)
	w.WriteHeader(200)
	w.Write([]byte("welcome home"))

}

func Run(cmd *cobra.Command, args []string) error {
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		os.Exit(1)
		return err
	}
	if port == "" {
		port = "3000"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	http.ListenAndServe(":"+port, mux)
	return nil
}
