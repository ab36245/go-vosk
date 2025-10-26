package vosk

import "github.com/ab36245/go-writer"

type VoskResult struct {
	Alternatives []VoskAlternative `json:"alternatives"`
	Partial      string            `json:"partial"`
}

func (v VoskResult) NotEmpty() bool {
	return len(v.Alternatives) > 0 || v.Partial != ""
}

func (v VoskResult) String() string {
	return writer.Value(v)
}

type VoskAlternative struct {
	Confidence float32    `json:"confidence"`
	Result     []VoskWord `json:"result"`
	Text       string     `json:"text"`
}

func (v VoskAlternative) String() string {
	return writer.Value(v)
}

type VoskWord struct {
	End   float32 `json:"end"`
	Start float32 `json:"start"`
	Word  string  `json:"word"`
}

func (v VoskWord) String() string {
	return writer.Value(v)
}
