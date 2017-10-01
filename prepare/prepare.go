package prepare

import (
	"fmt"
	"io/ioutil"
	"os"

	"strings"

	"github.com/ChimeraCoder/gojson"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mcuadros/go-jsonschema-generator"
	"github.com/spf13/cobra"
)

func MakeCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("prepare called!!!")
		sd := &jsonschema.Document{}
		//sd.Read(&struct {
		//	Type linebot.MessageType `json:"type"`
		//	Text string              `json:"text"`
		//}{})
		tm := linebot.NewTextMessage("dummy")
		ba, _ := tm.MarshalJSON()
		fmt.Println(string(ba))
		js, err := gojson.ParseJson(strings.NewReader(string(ba)))
		if err != nil {
			panic(err)
		}
		fmt.Println(js)
		//sd.Read(&linebot.Event{})
		sd.Read(js)
		ioutil.WriteFile("schemajson/reply-textmessage.json", []byte(sd.String()), os.ModePerm)
	}
}
