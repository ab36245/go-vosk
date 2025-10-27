module github.com/ab36245/go-vosk

go 1.24.2

replace github.com/ab36245/go-errors => ../go-errors

require (
	github.com/ab36245/go-errors v0.0.0-20250428061939-8b056c3b905e
	github.com/ab36245/go-runner v0.0.0-20250408015956-639ed8110e13
	github.com/ab36245/go-stream v0.0.0-20250610105203-92d1af5dbbb7
	github.com/ab36245/go-writer v0.0.0-20250625064440-6b35dc74f95b
	github.com/alphacep/vosk-api/go v0.3.50
)

replace github.com/ab36245/go-writer => ../go-writer

replace github.com/ab36245/go-runner => ../go-runner

replace github.com/ab36245/go-stream => ../go-stream
