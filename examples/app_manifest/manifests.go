package main

import (
	"fmt"

	"github.com/mldelaro/slack"
)

func main() {
	api := slack.New("TOKEN_HERE")
	manifest, err := api.ExportAppManifest("A02TDSWCDDE")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("DisplayName: %s\n", manifest.DisplayInformation.Name)

	// manifestCreateResponse, err := api.CreateAppManifest("{\"_metadata\":{\"major_version\":1,\"minor_version\":1},\"display_information\":{\"name\":\"Some Bot from GoLang\"},\"features\":{\"bot_user\":{\"display_name\":\"delme\",\"always_online\":false},\"slash_commands\":[{\"command\":\"/some-slash-command\",\"url\":\"https://some_lambda.execute-api.us-west-2.amazonaws.com/event/slash-commands\",\"description\":\"Some Bot Description\",\"usage_hint\":\"[help]\",\"should_escape\":false}]},\"oauth_config\":{\"redirect_urls\":[\"https://some_lambda.execute-api.us-west-2.amazonaws.com/event/oauth\"],\"scopes\":{\"bot\":[\"chat:write\",\"commands\",\"im:read\",\"app_mentions:read\",\"channels:history\",\"groups:history\",\"im:history\",\"mpim:history\"]}},\"settings\":{\"event_subscriptions\":{\"request_url\":\"https://some_lambda.execute-api.us-west-2.amazonaws.com/event\",\"bot_events\":[\"message.channels\",\"message.groups\",\"message.im\",\"message.mpim\"]},\"org_deploy_enabled\":false,\"socket_mode_enabled\":false,\"token_rotation_enabled\":false}}")
	// if err != nil {
	// 	fmt.Printf("%s\n", err)
	// 	return
	// }
	// fmt.Printf("AppID: %s\n", manifestCreateResponse.AppId)
	// TODO: Delete me -- A02TZJ50R40
}
