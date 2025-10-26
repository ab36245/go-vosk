package vosk

import (
	"slices"

	"github.com/ab36245/go-stream"
)

func New(samplesPerSecond int) *Vosk {
	input := make(chan []byte)

	var stream *stream.Output[VoskResult]
	if samplesPerSecond == 0 {
		stream = doTranscode(input)
	} else {
		stream = doDirect(samplesPerSecond, input)
	}
	return &Vosk{
		input:  input,
		stream: stream,
	}
}

type Vosk struct {
	input  chan []byte
	stream *stream.Output[VoskResult]
}

func (v Vosk) Close() {
	close(v.input)
}

func (v Vosk) Input(bytes []byte) {
	// v.input <- append([]byte{}, bytes...)
	v.input <- slices.Clone(bytes)
}

func (v Vosk) Output() <-chan VoskResult {
	return v.stream.Output()
}

func (v Vosk) Result() <-chan error {
	return v.stream.Result()
}

func doDirect(samplesPerSecond int, input chan []byte) *stream.Output[VoskResult] {
	return stream.New(
		func(output func(VoskResult)) error {
			return recognize(samplesPerSecond, input, output)
		},
		func() {
			close(input)
		},
	)
}

func doTranscode(input chan []byte) *stream.Output[VoskResult] {
	const samplesPerSecond = 16000

	transcoder := stream.New(
		func(output func([]byte)) error {
			return transcode(samplesPerSecond, input, output)
		},
		func() {
			close(input)
		},
	)
	return stream.Add(
		transcoder,
		func(input <-chan []byte, output func(VoskResult)) error {
			return recognize(samplesPerSecond, input, output)
		},
	)
}
