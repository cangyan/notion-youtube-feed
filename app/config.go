package app

type Config struct {
	NotionToken         string `json:"notion_token,omitempty"`
	NotionDatabaseId    string `json:"notion_database_id,omitempty"`
	YouTubeApiKey       string `json:"you_tube_api_key,omitempty"`
	YouTubeClientId     string `json:"you_tube_client_id,omitempty"`
	YouTubeClientSecret string `json:"you_tube_client_secret,omitempty"`
}
