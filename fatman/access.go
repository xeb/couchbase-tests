// Fatman is something that stress tests Couchbase
package fatman

import (
	"fmt"
	"os"
	"github.com/couchbaselabs/go-couchbase"
	"time"
)

func Connect() (bucket *couchbase.Bucket) {
	bucket, err := couchbase.GetBucket("http://localhost:8091/", "default", "default")
	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	}
	return bucket
}

func WriteDocs(numWrites int, bucket *couchbase.Bucket) {
	for i := 0; i < numWrites; i++ {
		WriteDoc(i, bucket)
	}
}

func WriteDocsAsync(numWrites int, bucket *couchbase.Bucket) {
	if(numWrites == 0) { return }

	ops := make(chan int, numWrites)

	for i := 0; i < numWrites; i++ {
		go func(j int) {
			WriteDoc(i, bucket)
			ops <- j
		}(i)
	}

	// is there a better way to do this?
	for i := 0; i < numWrites; i++ {
		_ = <-ops
	}
}

func WriteDoc(i int, bucket *couchbase.Bucket) {
	t1 := time.Now()
	key := fmt.Sprintf("account-%d", i)
	account := &Account{i, fmt.Sprintf("test%d@test.com", i), t1}
	bucket.Set(key, -1, account)
}

// TODO: ReadDocsAsync and WriteDocsAsync look the same -- make into an interface?
func ReadDocsAsync(numReads int, bucket *couchbase.Bucket) {
	if(numReads == 0) { return }

	ops := make(chan int, numReads)

	for i := 0; i < numReads; i++ {
		go func(j int) {
			ReadDoc(i, bucket)
			ops <- j
		}(i)
	}

	// is there a better way to do this?
	for i := 0; i < numReads; i++ {
		_ = <-ops
	}
}

func ReadDocs(numReads int, bucket *couchbase.Bucket) {
	for i := numReads - 1; i >= 0; i-- {
		ReadDoc(i, bucket)
	}
}

func ReadDoc(i int, bucket *couchbase.Bucket) {
	key := fmt.Sprintf("account-%d", i)
	var account Account
	_ = bucket.Get(key, &account)
}
