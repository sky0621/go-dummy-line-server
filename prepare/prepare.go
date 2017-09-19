package prepare

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mcuadros/go-jsonschema-generator"
	"github.com/spf13/cobra"
)

func MakeCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("prepare called!!!")
		sd := &jsonschema.Document{}
		sd.Read(&struct {
			Type linebot.MessageType `json:"type"`
			Text string              `json:"text"`
		}{})
		ioutil.WriteFile("schemajson/reply-textmessage.json", []byte(sd.String()), os.ModePerm)
	}
}
