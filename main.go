package main

import (
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/davecgh/go-spew/spew"
)

func Receive(event cloudevents.Event) {
	// do something with event.Context and event.Data (via event.DataAs(foo)
	var data interface{}
	fmt.Println("new event", event.ID(), event.DataAs(&data))
	spew.Dump(data)
}

func main() {
	spew.Config.SpewKeys = false
	spew.Config.SortKeys = false
	spew.Config.DisableCapacities = true
	spew.Config.DisableMethods = true
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisablePointerMethods = true

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(8081),
		cloudevents.WithStructuredEncoding(),
	)
	if err != nil {
		panic(err)
	}
	c, err := cloudevents.NewClient(t)
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), Receive))
}