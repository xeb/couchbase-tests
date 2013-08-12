package main

import (
	"flag"
	"fmt"
	"github.com/couchbaselabs/go-couchbase"
	"os"
	"time"
)

var verbose bool
var writes int
var reads int
var fullTest bool

type Account struct {
	Id        int
	Name      string
	LastLogin time.Time
}

func init() {
	flag.IntVar(&writes, "writes", 0, "Number of writes")
	flag.IntVar(&reads, "reads", 0, "Number of reads")
	flag.BoolVar(&verbose, "verbose", false, "Verbose on/off")
	flag.BoolVar(&fullTest, "fullTest", false, "Run Full Test (scaling up to 1M reads/writes)")
	flag.Parse()
}

func writeDocs(numWrites int, bucket *couchbase.Bucket) {

	// insert a bunch of documents
	for i := 0; i < numWrites; i++ {
		t1 := time.Now()
		key := fmt.Sprintf("account-%d", i)
		account := &Account{i, fmt.Sprintf("test%d@test.com", i), t1}
		bucket.Set(key, -1, account)
		if verbose {
			fmt.Printf("Saved account %s\n", account)
		}
	}
}

func readDocs(numReads int, bucket *couchbase.Bucket) {
	for i := numReads - 1; i >= 0; i-- {

		key := fmt.Sprintf("account-%d", i)
		var account Account
		_ = bucket.Get(key, &account)
		if verbose {
			fmt.Printf("Reading Doc %d, %s\n", i, account)
		}
	}
}

func main() {

	fmt.Printf("Using: writes=%d, reads=%d, fullTest=%t, verbose=%t\n", writes, reads, fullTest, verbose)

	t0 := time.Now()
	fmt.Print("Connecting to localhost...")
	bucket, err := couchbase.GetBucket("http://localhost:8091/", "default", "default")
	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", time.Now().Sub(t0))

	if fullTest {
		tests := []int{10, 100, 1000, 10000, 25000, 50000, 100000, 250000, 500000, 750000, 1000000}
		// tests := []int{10, 100, 1000}
		for _, test := range tests {

			t0 = time.Now()
			writeDocs(test, bucket)
			writeTime := time.Now().Sub(t0)

			t0 = time.Now()
			readDocs(test, bucket)
			readTime := time.Now().Sub(t0)

			fmt.Printf("%d\treads=%d\twrites=%d\n", test, readTime.Nanoseconds() / 1000000, writeTime.Nanoseconds() / 1000000)
		}
	} else {
		t0 = time.Now()
		writeDocs(writes, bucket)
		fmt.Printf("Inserted %d documents in %s\n", writes, time.Now().Sub(t0))

		t0 = time.Now()
		readDocs(reads, bucket)
		fmt.Printf("Read %d documents in %s\n", reads, time.Now().Sub(t0))
	}

	fmt.Println("Done!")
}
