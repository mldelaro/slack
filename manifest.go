package slack

import (
	"context"
	"encoding/json"
	// "net/url"
)

type exportManifestResponse struct {
	SlackResponse
	Manifest *Manifest `json:"manifest"`
}

type Manifest struct {
	Metadata           *Metadata           `json:"_metadata"`
	DisplayInformation *DisplayInformation `json:"display_information"`
	Features           *Features           `json:"features,omitempty"`
	OAuthConfig        *OAuthConfig        `json:"oauth_config,omitempty"`
	Settings           *Settings           `json:"settings"`
}

type Metadata struct {
	Majorversion int `json:"major_version"`
	Minorversion int `json:"minor_version"`
}

type DisplayInformation struct {
	Name string `json:"name"`
}

type Features struct {
	BotUser       *BotUser       `json:"bot_user,omitempty"`
	SlashCommands []SlashCommandManifest `json:"slash_commands,omitempty"`
}

type BotUser struct {
	DisplayName  string `json:"display_name,omitempty"`
	AlwaysOnline bool   `json: "always_online,omitempty"`
}

// Naming conflict with "SlashCommand" via slash.go
type SlashCommandManifest struct {
	Command      string `json:"command,omitempty"`
	Url          string `json:"url,omitempty"`
	Description  string `json:"description,omitempty"`
	UsageHint    string `json:"usage_hint,omitempty"`
	ShouldEscape bool   `json:"should_escape,omitempty"`
}

type OAuthConfig struct {
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	Scopes       Scopes   `json:"scopes,omitempty"`
}

type Scopes struct {
	bot []string `json:"bot,omitempty"`
}

type EventSubscriptions struct {
	RequestUrl string   `json:"request_url,omitempty"`
	BotEvents  []string `json:"bot_events,omitempty"`
}

type Settings struct {
	EventSubscriptions   *EventSubscriptions `json:"event_subscriptions"`
	OrgDeployEnabled     bool                `json:"org_deploy_enabled"`
	SocketModeEnabled    bool                `json:"socket_mode_enabled"`
	TokenRotationEnabled bool                `json:"token_rotation_enabled"`
}

func (api *Client) ExportAppManifest(appId string) (*Manifest, error) {
	return api.ExportAppManifestContext(context.Background(), appId)
}

// ListEventAuthorizationsContext lists authed users and teams for the given event_context. You must provide an app-level token to the client using OptionAppLevelToken. More info: https://api.slack.com/methods/apps.event.authorizations.list
func (api *Client) ExportAppManifestContext(ctx context.Context, appId string) (*Manifest, error) {
	resp := &exportManifestResponse{}

	request, _ := json.Marshal(map[string]string{
		"app_id": appId,
	})

	err := postJSON(ctx, api.httpclient, api.endpoint+"apps.manifest.export", api.token, request, &resp, api)

	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Err()
	}

	return resp.Manifest, nil
}

/*
func (api *Client) UninstallApp(clientID, clientSecret string) error {
	values := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
	}

	response := SlackResponse{}

	err := api.getMethod(context.Background(), "apps.uninstall", api.token, values, &response)
	if err != nil {
		return err
	}

	return response.Err()
}
*/
