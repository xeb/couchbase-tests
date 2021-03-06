Go Couchbase Tests
================
A synthetic benchmark to test the performance of Couchbase.

### Setup
A little repository to test out Couchbase for work.  I ran these tests on:
* Late 2010 Macbook Air (1.86GHz Core 2 Duo, 2GB DD3) against 
* Couchebase 2.1.1-764-rel with 
* 512MB in the cluster 
* 256MB for the "default" bucket
* [go-couchbase](http://github.com/couchbaselabs/go-couchbase)

### Model Used
A very simple model for an Account is used in this test.

```
type Account struct {
	Id        int
	Name      string
	LastLogin time.Time
}
```

### Run Some Tests
The main package is easy to run.  Just specify the number of writes & reads you'd like to simulate.

```
go run main.go -writes=12 -reads=7777 -verbose
```

Or if you'd like to run the full tests below, you can via:
```
go run main.go -fullTest
```

Or you can use the golang [benchmarking library](http://golang.org/pkg/testing/)
```
go test bench_test.go -bench="Couchbase"
```

### Goroutine Tests
I just added some goroutine support.  You can turn it on with the "-goroutine" switch.  It will do all reads and writes in separate goroutines if you want.  The performance improvements of the goroutines seem to be more on writes.  Take this singular example:

```
xeb$ go build main.go
xeb$ ./main -reads=10000 -writes=10000 -goroutine=false
Using: writes=10000, reads=10000, fullTest=false, verbose=false, goroutine=false
Inserted 10000 documents in 4.101123344s
Read 10000 documents in 2.642911385s
Done!
```
And now with goroutines turned on...
```
xeb$ ./main -reads=10000 -writes=10000 -goroutine=true
Using: writes=10000, reads=10000, fullTest=false, verbose=false, goroutine=true
Inserted 10000 documents in 2.906022295s
Read 10000 documents in 2.460277874s
Done!
```
Although there are consequences at certain scales...
```
xeb$ ./main -reads=20000 -writes=20000 -goroutine=false
Using: writes=20000, reads=20000, fullTest=false, verbose=false, goroutine=false
Inserted 20000 documents in 7.01395426s
Read 20000 documents in 3.57084739s
Done!
xeb$ ./main -reads=20000 -writes=20000 -goroutine=true
Using: writes=20000, reads=20000, fullTest=false, verbose=false, goroutine=true
Inserted 20000 documents in 9.988769682s
Read 20000 documents in 8.863286616s
Done!
```

### Results
<table>
	<tr>
		<th>Count</th>
		<th>Duration (ms)<br/><em>(Read)</em></th>
		<th>Duration (ms)<br/><em>(Write)</em></th>
		<th>Ops per Second<br/><em>(Read)</em></th>
		<th>Ops per Second<br/><em>(Write)</em></th>
	</tr>
	<tr>
		<td>10</td>
		<td>2ms</td>
		<td>4ms</td>
		<td>5,000/sec</td>
		<td>2,500/sec</td>
	</tr>
	<tr>

		<td>100</td>
		<td>45</td>
		<td>49</td>
		<td>2,222</td>
		<td>2,040</td>
	</tr>
	<tr>

		<td>1,000</td>
		<td>179</td>
		<td>228</td>
		<td>5,586</td>
		<td>4,385</td>
	</tr>
	<tr>

		<td>10,000</td>
		<td>2,107</td>
		<td>2,114</td>
		<td>4,746</td>
		<td>4,730</td>
	</tr>
	<tr>

		<td>25,000</td>
		<td>4,596</td>
		<td>6,302</td>
		<td>5,439</td>
		<td>3,966</td>
	</tr>
	<tr>

		<td>50,000</td>
		<td>9,471</td>
		<td>13,718</td>
		<td>5,279</td>
		<td>3,644</td>
	</tr>
	<tr>

		<td>100,000</td>
		<td>17,940</td>
		<td>29,248</td>
		<td>5,574</td>
		<td>3,419</td>
	</tr>
	<tr>
		<td>250,000</td>
		<td>36,450</td>
		<td>60,444</td>
		<td>6,858</td>
		<td>4,136</td>
	</tr>
	<tr>

		<td>500,000</td>
		<td>77,706</td>
		<td>124,862</td>
		<td>6,434</td>
		<td>4,004</td>
	</tr>
	<tr>

		<td>750,000</td>
		<td>101,358</td>
		<td>201,432</td>
		<td>7,399</td>
		<td>3,723</td>
	</tr>
	<tr>

		<td>1,000,000</td>
		<td>134,733</td>
		<td>221,826</td>
		<td>7,422</td>
		<td>4,508</td>
	</tr>
	<tr>
		<td colspan="3"><em>Weighted Avg</em></td>
		<td>7,041/sec</td>
		<td>4,086/sec</td>
	</tr>
</table>


### From Couchbase Server

![8k ops / sec](https://raw.github.com/xeb/couchbase-tests/master/docs/performance.png "Performance")

![growth rates](https://raw.github.com/xeb/couchbase-tests/master/docs/performance2.png "Performance 2")

### Conclusion
Couchbase has tremendous performance.  On average, it appears that writes are only 2 times as slow as reads in this scenario.
