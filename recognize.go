package vosk

import (
	"encoding/json"

	vosk "github.com/alphacep/vosk-api/go"
)

func recognize(
	samplesPerSecond int,
	input <-chan []byte,
	output func(VoskResult),
) error {
	recognizer, err := vosk.NewRecognizer(model, float64(samplesPerSecond))
	if err != nil {
		return RecognizerError.Wrap(err)
	}

	// Be very, very careful if changing these
	// The format of the JSON results returned by the recognizer changes
	// dramatically depending on whether alternatives and/or words are required
	recognizer.SetMaxAlternatives(5)
	recognizer.SetWords(1)

	for {
		bytes, ok := <-input
		if bytes != nil {
			var str string
			if recognizer.AcceptWaveform(bytes) != 0 {
				str = recognizer.Result()
			} else {
				str = recognizer.PartialResult()
			}
			var result VoskResult
			if err := json.Unmarshal([]byte(str), &result); err != nil {
				return JsonError.Wrap(err)
			}
			output(result)
		}
		if !ok {
			str := recognizer.FinalResult()
			var result VoskResult
			if err := json.Unmarshal([]byte(str), &result); err != nil {
				return JsonError.Wrap(err)
			}
			output(result)
			return nil
		}
	}
}
