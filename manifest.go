package slack

import (
	"context"
	"encoding/json"
	// "net/url"
)

type ExportManifestResponse struct {
	SlackResponse
	Manifest *Manifest `json:"manifest"`
}

type Manifest struct {
	DisplayInformation *DisplayInformation `json:"display_information"`
	Features           *Features           `json:"features,omitempty"`
	OAuthConfig        *OAuthConfig        `json:"oauth_config,omitempty"`
	Settings           *Settings           `json:"settings"`
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
	Scopes       *Scopes   `json:"scopes,omitempty"`
}

type Scopes struct {
	Bot []string `json:"bot,omitempty"`
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

type NewManifestResponse struct {
	SlackResponse
	AppId              string        `json:"app_id"`
	OAuthAuthorizeUrl  string        `json:"oauth_authorize_url"`
	Credentials        *Credentials  `json:"credentials"`
}

type Credentials struct {
	ClientId          string  `json:"client_id"`
	ClientSecret      string  `json:"client_secret"`
	VerificationToken string  `json:"verification_token"`
	SigningSecret     string  `json:"signing_secret"`
}

type UpdateManifestResponse struct {
	SlackResponse
	AppId               string  `json:"app_id"`
	PermissionsUpdated  bool    `json:"permissions_updated"`
}


func (api *Client) ExportAppManifest(appId string) (*Manifest, error) {
	return api.ExportAppManifestContext(context.Background(), appId)
}

func (api *Client) CreateAppManifest(appManifest string) (*NewManifestResponse, error) {
	return api.CreateAppManifestContext(context.Background(), appManifest)
}

func (api *Client) UpdateAppManifest(appId string, appManifest string) (*UpdateManifestResponse, error) {
	return api.UpdateAppManifestContext(context.Background(), appId, appManifest)
}

func (api *Client) DeleteAppManifest(appId string) (*SlackResponse, error) {
	return api.DeleteAppManifestContext(context.Background(), appId)
}

// ExportAppManifestContext gets the manifest file for a given slack-app-id. You must provide an app-level token to the client using OptionAppLevelToken. More info: https://api.slack.com/methods/apps.event.authorizations.list
func (api *Client) ExportAppManifestContext(ctx context.Context, appId string) (*Manifest, error) {
	resp := &ExportManifestResponse{}

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


// CreateAppManifestContext creates a new app for a given app-manifest. You must provide an app-level token to the client using OptionAppLevelToken. More info: https://api.slack.com/methods/apps.event.authorizations.list
func (api *Client) CreateAppManifestContext(ctx context.Context, appManifest string) (*NewManifestResponse, error) {
	resp := &NewManifestResponse{}

	request, _ := json.Marshal(map[string]string{
		"manifest": appManifest,
	})

	err := postJSON(ctx, api.httpclient, api.endpoint+"apps.manifest.create", api.token, request, &resp, api)

	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Err()
	}

	return resp, nil
}

// UpdateAppManifestContext updates a given app's manifest. You must provide an app-level token to the client using OptionAppLevelToken. More info: https://api.slack.com/methods/apps.event.authorizations.list
func (api *Client) UpdateAppManifestContext(ctx context.Context, appId string, appManifest string) (*UpdateManifestResponse, error) {
	resp := &UpdateManifestResponse{}

	request, _ := json.Marshal(map[string]string{
		"app_id": appId,
		"manifest": appManifest,
	})

	err := postJSON(ctx, api.httpclient, api.endpoint+"apps.manifest.update", api.token, request, &resp, api)

	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Err()
	}

	return resp, nil
}

func (api *Client) DeleteAppManifestContext(ctx context.Context, appId string) (*SlackResponse, error) {
	resp := &SlackResponse{}

	request, _ := json.Marshal(map[string]string{
		"app_id": appId,
	})

	err := postJSON(ctx, api.httpclient, api.endpoint+"apps.manifest.delete", api.token, request, &resp, api)

	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Err()
	}

	return resp, nil
}
