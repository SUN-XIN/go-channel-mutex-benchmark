# go-channel-mutex-benchmark

# Output with 2 goroutines parallel
    $>go test -test.bench Bench
    BenchmarkChan-4           	     100	  21430065 ns/op
    BenchmarkChanParallel-4   	     100	  21683900 ns/op
    BenchmarkMutex-4          	    1000	   1618203 ns/op

    
# Output with 8 goroutines parallel
    $> go test -test.bench Bench
    BenchmarkChan-4           	     100	  21267335 ns/op
    BenchmarkChanParallel-4   	      50	  29291662 ns/op
    BenchmarkMutex-4          	    1000	   1603088 ns/op

# Output with 16 goroutines parallel
    $> go test -test.bench Bench
    BenchmarkChan-4           	     100	  21163605 ns/op
    BenchmarkChanParallel-4   	      50	  32103690 ns/op
    BenchmarkMutex-4          	    1000	   1649793 ns/op
    
# Conclusion
  * channel is slower than Mutex
  * channel is slower as we add more consumers 
