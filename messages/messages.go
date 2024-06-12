package messages

import (
	"encoding/json"
	"errors"
)

// Message represents a generic Nomad message.
type Message struct {
	Type       string          `json:"type"`
	Encoding   string          `json:"encoding"`
	Sender     string          `json:"sender"`
	SiteID     string          `json:"site_id"`
	Recipients []string        `json:"recipients"`
	Version    string          `json:"version"`
	Data       json.RawMessage `json:"data"`
}

// PurgeMessage represents a purge message.
type PurgeMessage struct {
	Message
}

// RefreshMessage represents a refresh message.
type RefreshMessage struct {
	Message
}

// RekeyMessage represents a rekey message.
type RekeyMessage struct {
	Message
	Update bool `json:"update"`
}

// ActivityMessage represents an activity message.
type ActivityMessage struct {
	Message
}

// ResponseMessage represents a response message.
type ResponseMessage struct {
	Message
}

// SyncMessage represents a sync message.
type SyncMessage struct {
	Message
}

// NewMessage creates a new generic message.
func NewMessage(msgType, encoding, sender, siteID string, recipients []string, data interface{}) (*Message, error) {
	rawData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &Message{
		Type:       msgType,
		Encoding:   encoding,
		Sender:     sender,
		SiteID:     siteID,
		Recipients: recipients,
		Version:    "12.0",
		Data:       rawData,
	}, nil
}

// ParseMessage parses a JSON-encoded message.
func ParseMessage(jsonData []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(jsonData, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// HandleMessage handles the message based on its type.
func HandleMessage(msg *Message) (interface{}, error) {
	switch msg.Type {
	case "purge":
		return handlePurgeMessage(msg)
	case "refresh":
		return handleRefreshMessage(msg)
	case "rekey":
		return handleRekeyMessage(msg)
	case "activity":
		return handleActivityMessage(msg)
	case "response":
		return handleResponseMessage(msg)
	case "sync":
		return handleSyncMessage(msg)
	default:
		return nil, errors.New("unknown message type")
	}
}

func handlePurgeMessage(msg *Message) (*PurgeMessage, error) {
	var purgeMsg PurgeMessage
	err := json.Unmarshal(msg.Data, &purgeMsg)
	if err != nil {
		return nil, err
	}
	return &purgeMsg, nil
}

func handleRefreshMessage(msg *Message) (*RefreshMessage, error) {
	var refreshMsg RefreshMessage
	err := json.Unmarshal(msg.Data, &refreshMsg)
	if err != nil {
		return nil, err
	}
	return &refreshMsg, nil
}

func handleRekeyMessage(msg *Message) (*RekeyMessage, error) {
	var rekeyMsg RekeyMessage
	err := json.Unmarshal(msg.Data, &rekeyMsg)
	if err != nil {
		return nil, err
	}
	return &rekeyMsg, nil
}

func handleActivityMessage(msg *Message) (*ActivityMessage, error) {
	var activityMsg ActivityMessage
	err := json.Unmarshal(msg.Data, &activityMsg)
	if err != nil {
		return nil, err
	}
	return &activityMsg, nil
}

func handleResponseMessage(msg *Message) (*ResponseMessage, error) {
	var responseMsg ResponseMessage
	err := json.Unmarshal(msg.Data, &responseMsg)
	if err != nil {
		return nil, err
	}
	return &responseMsg, nil
}

func handleSyncMessage(msg *Message) (*SyncMessage, error) {
	var syncMsg SyncMessage
	err := json.Unmarshal(msg.Data, &syncMsg)
	if err != nil {
		return nil, err
	}
	return &syncMsg, nil
}
