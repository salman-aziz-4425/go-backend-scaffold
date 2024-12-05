package models

type Message struct {
	Type     string `json:"type"`
	Group    string `json:"group"`
	PeerID   string `json:"peerId"`
	Muted    bool   `json:"muted"`
	VideoOff bool   `json:"isVideoOff"`
}
