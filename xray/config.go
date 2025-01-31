package xray

import "encoding/json"

type Config struct {
	LogConfig       json.RawMessage `json:"log"`
	RouterConfig    json.RawMessage `json:"routing"`
	DNSConfig       json.RawMessage `json:"dns"`
	InboundConfigs  []InboundConfig `json:"inbounds"`
	OutboundConfigs json.RawMessage `json:"outbounds"`
	Transport       json.RawMessage `json:"transport"`
	Policy          json.RawMessage `json:"policy"`
	API             json.RawMessage `json:"api"`
	Stats           json.RawMessage `json:"stats"`
	Reverse         json.RawMessage `json:"reverse"`
	FakeDNS         json.RawMessage `json:"fakeDns"`
}
