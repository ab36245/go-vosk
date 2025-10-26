package vosk

import (
	"strconv"

	"github.com/ab36245/go-runner"
)

func transcode(
	samplesPerSecond int,
	input chan []byte,
	output func([]byte),
) error {
	cmd := runner.New(
		"ffmpeg",
		// "-copytb", "1",
		"-loglevel", "warning",
		// "-i", "pipe:0",
		"-i", "-",
		"-ar", strconv.FormatInt(int64(samplesPerSecond), 10),
		"-ac", "1",
		"-f", "s16le",
		// "-f", "wav",
		// "pipe:1",
		"-",
	)

	stdin := runner.StreamInput(input)
	cmd.Stdin(stdin)

	stdout := runner.StreamOutput()
	cmd.Stdout(stdout)

	cmd.Stderr(nil)

	if err := cmd.Start(); err != nil {
		return TranscoderError.Wrap(err)
	}

	for {
		select {
		case bytes, ok := <-stdout.Channel():
			if bytes != nil {
				output(bytes)
			}
			if !ok {
				return nil
			}
		case err := <-cmd.Result():
			return err
		}
	}
}
