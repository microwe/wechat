package messages

import (
	"encoding/base64"
	"github.com/ArtisanCloud/PowerLibs/object"
)

type DeviceText struct {
	*Message
}

func NewDeviceText(items *object.HashMap) *DeviceText {
	m := &DeviceText{
		NewMessage(items),
	}

	m.Type = "device_text"

	m.Properties = []string{
		"device_type",
		"device_id",
		"content",
		"session_id",
		"open_id",
	}

	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *DeviceText) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {

		return &object.HashMap{
			"DeviceType": msg.Get("device_type", nil),
			"DeviceID":   msg.Get("device_id", nil),
			"SessionID":  msg.Get("session_id", nil),
			"Content":    base64.StdEncoding.DecodeString(msg.GetString("content", "")),
		}
	}
}