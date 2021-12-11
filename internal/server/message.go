package server

const (
	messageTypeError = "error"
	messageTypeOK    = "ok"
)

type message struct {
	Type      string `json:"type"`
	ErrorInfo string `json:"errorInfo,omitempty"`
	Data      data   `json:"data"`
}

type data struct {
	Diff    string `json:"diff"`
	Weather string `json:"weather"`
}

func getDiffWeatherMessage(diff string, weather string) *message {
	return &message{
		Type: messageTypeOK,
		Data: data{
			Diff:    diff,
			Weather: weather,
		},
	}
}
