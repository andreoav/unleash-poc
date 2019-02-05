package main

import (
	"fmt"
	"net/http"

	unleash "github.com/Unleash/unleash-client-go/v3"
	"github.com/Unleash/unleash-client-go/v3/context"
)

func init() {
	fmt.Println("Initializing application,")

	headers := http.Header{}
	headers.Add("Authorization", "12345")

	unleash.Initialize(
		unleash.WithCustomHeaders(headers),
		unleash.WithAppName("my-application"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithListener(&unleash.DebugListener{}),
	)

	fmt.Println("Initialization finished.")
}

func main() {
	fmt.Println("Starting application.")

	fmt.Println(unleash.IsEnabled("Feature001", unleash.WithContext(context.Context{
		UserId: "02287813020",
	})))

	fmt.Println(unleash.IsEnabled("Feature001", unleash.WithContext(context.Context{
		UserId: "02287813021",
	})))

	fmt.Println("Exiting application.")
}
