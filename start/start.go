package start

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
)

func MakeCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("start called!!")

		schema := gojsonschema.NewReferenceLoader("schemajson/reply-textmessage.json")
		//tmpjson := gojsonschema.NewReferenceLoader("tmp.json")
		tmpjson := gojsonschema.NewStringLoader(`{"type":"idid","text":"txtx"}`)
		res, err := gojsonschema.Validate(schema, tmpjson)
		if err != nil {
			panic(err)
		}

		if res.Valid() {
			fmt.Println("Valid!")
		} else {
			fmt.Println("Not Valid!!")
		}
	}
}
