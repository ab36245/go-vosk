package vosk

import (
	vosk "github.com/alphacep/vosk-api/go"
)

var model *vosk.VoskModel

func Init(path string) error {
	vosk.SetLogLevel(-1)
	if model == nil {
		var err error
		model, err = vosk.NewModel(path)
		if err != nil {
			return InitError.Wrap(err)
		}
	}
	return nil
}
