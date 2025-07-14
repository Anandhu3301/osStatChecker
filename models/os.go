package models

type SystemValues struct {
	OperatingSystem string `json:"operatingSystem"`
	Architecture  string `json:"architecture"`
	Memory Stats
	Error string `json:"error"`
}

type SystemGenerator func(*SystemValues);