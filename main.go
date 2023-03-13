package main

import (
	"log"
	"runtime/debug"
	"server/client/recreation_gov"
	"server/request"
)

func main() {
	client := recreation_gov.NewHttpClient()
	{
		var cmd, err = request.NewGetActivitiesCmd(client)
		if err != nil {
			log.Fatalln(err)
		}

		err = cmd.Execute()
		if err != nil {
			debug.PrintStack()
			log.Fatalf("Failed executing req [%v]: %s\n", cmd.RawReq, err)
		}

		log.Printf("Activities are: %s", cmd.ParsedResponse)
	}

	{
		cmd, err := request.NewGetRecAreasRequest(client)
		if err != nil {
			log.Fatalln(err)
		}

		err = cmd.Execute()
		if err != nil {
			debug.PrintStack()
			log.Fatalf("failed executing req [%v]: %s", cmd.RawReq.URL, err)
		}

		log.Printf("RecAreas are: %s", cmd.ParsedResponse)
	}
}
