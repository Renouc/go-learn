package web

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO 启动一个web服务器
		srv := &http.Server{
			Addr:    ":8080",
			Handler: http.DefaultServeMux,
		}

		prefix := color.New(color.FgBlue).Sprintf("[web address]")
		message := color.New(color.Underline, color.Bold, color.FgGreen).Sprintf("http://127.0.0.1:8080")
		fmt.Println(prefix, message)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
		})

		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Server failed to start: %v\n", err)
		}
	},
}
