// Fatman is something that tests Couchbase
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
		t1 := time.Now()
		key := fmt.Sprintf("account-%d", i)
		account := &Account{i, fmt.Sprintf("test%d@test.com", i), t1}
		bucket.Set(key, -1, account)
	}
}

func ReadDocs(numReads int, bucket *couchbase.Bucket) {
	for i := numReads - 1; i >= 0; i-- {
		key := fmt.Sprintf("account-%d", i)
		var account Account
		_ = bucket.Get(key, &account)
	}
}