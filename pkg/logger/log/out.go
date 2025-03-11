package log

import (
	"github.com/rs/zerolog"
	"io"
	"os"
)

type OutAndErrOutput struct {
	output    io.Writer
	errOutput io.Writer
}

func (l OutAndErrOutput) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (l OutAndErrOutput) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level > zerolog.InfoLevel {
		return l.errOutput.Write(p)
	}
	return l.output.Write(p)
}

func DefaultOutput(isPretty bool) OutAndErrOutput {
	if isPretty {
		return OutAndErrOutput{
			output:    zerolog.ConsoleWriter{Out: os.Stdout},
			errOutput: zerolog.ConsoleWriter{Out: os.Stderr},
		}
	}
	return OutAndErrOutput{
		output:    os.Stdout,
		errOutput: os.Stderr,
	}
}

type OutDump struct {
	Dump []byte
}

func (d *OutDump) Write(p []byte) (n int, err error) {
	d.Dump = p
	return len(p), nil
}

func NewOutDump() *OutDump {
	return &OutDump{}
}

type OutMultiDump struct {
	Dumps [][]byte
}

func (d *OutMultiDump) Write(p []byte) (n int, err error) {
	d.Dumps = append(d.Dumps, p)
	return len(p), nil
}

func NewOutMultiDump() *OutMultiDump {
	return &OutMultiDump{
		Dumps: make([][]byte, 0, 4),
	}
}
