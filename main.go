package main

import (
	"flag"
	"fmt"
	"time"
	"github.com/xeb/couchbase-tests/fatman"
)

var verbose bool
var writes int
var reads int
var goroutine bool
var fullTest bool

func init() {
	flag.IntVar(&writes, "writes", 0, "Number of writes")
	flag.IntVar(&reads, "reads", 0, "Number of reads")
	flag.BoolVar(&verbose, "verbose", false, "Verbose on/off")
	flag.BoolVar(&goroutine, "goroutine", false, "Whether or not to use goroutine for writes/reads")
	flag.BoolVar(&fullTest, "fullTest", false, "Run Full Test (scaling up to 1M reads/writes)")
	flag.Parse()
}

func main() {

	fmt.Printf("Using: writes=%d, reads=%d, fullTest=%t, verbose=%t, goroutine=%t\n", writes, reads, fullTest, verbose, goroutine)

	t0 := time.Now()
	bucket := fatman.Connect()

	if fullTest {
		tests := []int{10, 100, 1000, 10000, 25000, 50000, 100000, 250000, 500000, 750000, 1000000}
		
		for _, test := range tests {

			t0 = time.Now()
			fatman.WriteDocs(test, bucket)
			writeTime := time.Now().Sub(t0)

			t0 = time.Now()
			fatman.ReadDocs(test, bucket)
			readTime := time.Now().Sub(t0)

			fmt.Printf("%d\treads=%d\twrites=%d\n", test, readTime.Nanoseconds() / 1000000, writeTime.Nanoseconds() / 1000000)
		}

	} else {
		t0 = time.Now()
		if goroutine {
			fatman.WriteDocsAsync(writes, bucket)
		} else {
			fatman.WriteDocs(writes, bucket)
		}
		fmt.Printf("Inserted %d documents in %s\n", writes, time.Now().Sub(t0))

		t0 = time.Now()
		if goroutine {
			fatman.ReadDocsAsync(reads, bucket)
		} else {
			fatman.ReadDocs(reads, bucket)
		}
		fmt.Printf("Read %d documents in %s\n", reads, time.Now().Sub(t0))
	}

	fmt.Println("Done!")
}
