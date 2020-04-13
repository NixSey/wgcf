package trace

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	. "wgcf/cmd/util"
)

var shortMsg = "Prints trace information about the current internet connection"

var Cmd = &cobra.Command{
	Use:   "trace",
	Short: shortMsg,
	Long:  FormatMessage(shortMsg, `
Useful for verifying if Warp and Warp+ are working.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := trace(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
}

func trace() error {
	response, err := http.Get("https://cloudflare.com/cdn-cgi/trace")
	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	log.Println("Trace result:")
	fmt.Println(strings.TrimSpace(string(bodyBytes)))
	return nil
}
