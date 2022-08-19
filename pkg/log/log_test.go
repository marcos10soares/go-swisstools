package log_test

import (
	"testing"

	"github.com/marcos10soares/go-swisstools/pkg/log"

	"github.com/pkg/errors"
)

func BenchmarkInfoLog(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		log.Info("informational message")
	}
}

func BenchmarkErrorLog(b *testing.B) {
	err := errors.New("this is an error")

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		log.Error(err, "test")
	}
}
