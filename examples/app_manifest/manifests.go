package main

import (
	"fmt"

	"github.com/mldelaro/slack"
)

func main() {
	api := slack.New("YOUR_TOKEN_HERE")
	manifest, err := api.ExportAppManifest("A02TDSWCDDE")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("DisplayName: %s\n", manifest.DisplayInformation.Name)
}
