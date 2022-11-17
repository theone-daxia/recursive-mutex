
# RecursiveMutex

记录下极客时间课程的学习成果。

golang 的 Mutex 不可重入，这里用两种方法实现一个可重入锁。

1. 获取 goroutine id，来记录哪个 goroutine 获取了锁。
2. 由 goroutine 提供一个 token，来标识它自己。

## Usage/Examples

```go
import(
    "fmt"
    "github/theone-daxia/recursive-mutex"
)

func main() {
    var rmux recursive-mutex.RecursiveMutex
    rmux.Lock()
    rmux.Lock()
    fmt.Println(123)
    rmux.Unlock()
    rmux.Unlock()
}
```


## Running Tests

To run tests, run the following command

```bash
  go test
```


## License

[MIT](https://choosealicense.com/licenses/mit/)

