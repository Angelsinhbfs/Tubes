package discovery

import (
	"encoding/json"
	"errors"
	"net/http"
)

// DiscoveryDocument represents the structure of the discovery document.
type DiscoveryDocument struct {
	ID              string     `json:"id"`
	Aliases         []string   `json:"aliases"`
	PrimaryLocation Location   `json:"primary_location"`
	PublicKey       string     `json:"public_key"`
	Name            string     `json:"name"`
	Username        string     `json:"username"`
	Photo           Photo      `json:"photo"`
	ChannelRole     string     `json:"channel_role"`
	Searchable      bool       `json:"searchable"`
	AdultContent    bool       `json:"adult_content"`
	PublicForum     bool       `json:"public_forum"`
	Profile         Profile    `json:"profile"`
	Permissions     string     `json:"permissions"`
	Locations       []Location `json:"locations"`
	Site            Site       `json:"site"`
}

// Location represents the primary location of the channel.
type Location struct {
	Address        string `json:"address"`
	URL            string `json:"url"`
	ConnectionsURL string `json:"connections_url"`
	FollowURL      string `json:"follow_url"`
}

// Photo represents the photo of the channel.
type Photo struct {
	URL     string `json:"url"`
	Type    string `json:"type"`
	Updated string `json:"updated"`
}

// Profile represents the profile information of the channel.
type Profile struct {
	Description string `json:"description"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
	Marital     string `json:"marital"`
	Sexual      string `json:"sexual"`
	Locale      string `json:"locale"`
	Region      string `json:"region"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	About       string `json:"about"`
	Homepage    string `json:"homepage"`
	Hometown    string `json:"hometown"`
}

// Site represents the site information.
type Site struct {
	URL             string   `json:"url"`
	SiteSig         string   `json:"site_sig"`
	Post            string   `json:"post"`
	OpenWebAuth     string   `json:"openWebAuth"`
	AuthRedirect    string   `json:"authRedirect"`
	SiteKey         string   `json:"sitekey"`
	DirectoryMode   string   `json:"directory_mode"`
	Encryption      []string `json:"encryption"`
	ProtocolVersion string   `json:"protocol_version"`
	RegisterPolicy  string   `json:"register_policy"`
	AccessPolicy    string   `json:"access_policy"`
	Accounts        int      `json:"accounts"`
	Channels        int      `json:"channels"`
	Admin           string   `json:"admin"`
	About           string   `json:"about"`
	SiteHash        string   `json:"sitehash"`
	SiteName        string   `json:"sitename"`
	Logo            string   `json:"logo"`
	SellPage        string   `json:"sellpage"`
	Location        string   `json:"location"`
	Community       string   `json:"community"`
}

// DiscoverChannel discovers the channel using the provided URL.
func DiscoverChannel(url string) (*DiscoveryDocument, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/x-nomad+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to discover channel")
	}

	var doc DiscoveryDocument
	err = json.NewDecoder(resp.Body).Decode(&doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}
