package vosk

import "github.com/ab36245/go-errors"

var Error = errors.Make("vosk", nil)

var InitError = Error.Make("init")

var JsonError = Error.Make("json")

var RecognizerError = Error.Make("recognizer")

var TranscoderError = Error.Make("transcoder")
