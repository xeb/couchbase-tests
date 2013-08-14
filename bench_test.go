package couchbenchmarks

import (
	"fmt"
	"testing"
	"github.com/xeb/couchbase-tests/fatman"
)

var numberIterations int = 1

func BenchmarkCouchbaseRead(b *testing.B) {
	bucket := fatman.Connect()
	defer bucket.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s something", "testing")
		fatman.ReadDocs(numberIterations, bucket)
	}
}

func BenchmarkCouchbaseWrite(b *testing.B) {
	bucket := fatman.Connect()
	defer bucket.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s something", "testing")
		fatman.WriteDocs(numberIterations, bucket)
	}
}