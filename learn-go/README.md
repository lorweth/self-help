# Learn Go Language

## Learn Go concurrency

### Goroutine là 1 lightweight thread được quản lý bởi Go runtime.

`go f(x, y, z)` bắt đầu 1 goroutine running `f(x, y, z)`

### Channels là 1 typed conduit mà qua đó có thể gửi và nhân giá trị với channel operator `<-`

``` go
ch <- x // gửi v sang channel ch.
v := <-ch // nhận giá trị từ ch.
```

Luồng dữ liệu theo hướng mũi tên.

Kênh phải được tạo trước khi sử dụng, `ch := make(chan int)`

Theo mặc định, gửi và nhận block cho đển khi phía bên kia sẳn sàng. Điều này cho phép các goroutine đồng bộ mà không có các khóa rõ ràng hoặc các biến có điều kiện.

```Go
func sum(s []int, c chan int) {
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum // send sum to c
}

func main() {
  s := []int{7, 2, 8, -9, 4, 0}

  c := make(chan int)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)
  x, y := <-c, <-c // receive from c

  fmt.Println(x, y, x+y)
}
```

### Buggered Channels: Các kênh có thể được lưu vào bộ nhớ đệm

```Go
ch := make(chan int, 100) // cung cấp độ dài của bộ đệm vào tham số thứ 2
```

Chỉ gửi đến 1 khối kênh có bộ đệm đầy. Nhận khối khi bộ đệm trống.

### Range và Close

Sender có thể `close` 1 kênh để biểu thị rằng sẻ không gửi thêm giá trị nào nữa. Recivers có thể test xem 1 kênh đã bị đóng hay chưa bằng cách thêm tham số thứ 2 trong lúc nhận giá trị. 

```Go
v, isOpen := <-ch
```

`isOpen` sẽ failse nếu không có thêm giá trị nào để nhận và kênh đã bị đóng.

Vòng lặp `for i := range c` nhận giá trị từ kênh và lặp lại cho tới khi kênh đóng.

> Chỉ nên để sender đóng kênh. Gửi trên 1 kênh đã đóng sẽ gây ra panic.
> Kênh không giống files; Không nên thường xuyên đóng kênh. Đóng kênh chỉ cần thiết khi người nhân được thông báo là không còn giá trị nhận nào nữa, như là xử lý 1 range loop.

```Go
func fibonacci(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  close(c)
}

func main() {
  c := make(chan int, 10)
  go fibonacci(cap(c), c)
  for i := range c {
    fmt.Println(i)
  }
}
```

### Select cho phép 1 goroutine chờ trên nhiều thao tác giao tiếp.

Một khối được chọn cho đến khi một trong số các trường hợp của nó có thể chạy. Nếu nhiều cái sẵn sàng nó sẽ chọn ngẫu nhiên 1 cái.

```Go
func fibonacci(c, quit chan int) {
  x, y := 0, 1
  for {
    select {
    case c <- x:
      x, y = y, x+y
    case <-quit:
      fmt.Println("quit")
      return
    }
  }
}

func main() {
  c := make(chan int)
  quit := make(chan int)
  go func() {
    for i := 0; i < 10; i++ {
      fmt.Println(<-c)
    }
    quit <- 0
  }()
  fibonacci(c, quit)
}
```

`default` case trong 1 `select` chạy nếu không case nào sẵn sàng.

Sử dụng `default` case đển thử gửi hoặc nhận mà không bị chặn

```Go
func main() {
  tick := time.Tick(100 * time.Millisecond)
  boom := time.After(500 * time.Millisecond)
  for {
    select {
    case <-tick:
      fmt.Println("tick.")
    case <-boom:
      fmt.Println("BOOM!")
      return
    default:
      fmt.Println("    .")
      time.Sleep(50 * time.Millisecond)
    }
  }
}
```

### Bài tập cây nhị phân tương đương

```Go

// đệ quy duyệt cây Right Node Left
func RNL(t *tree.Tree, ch chan int) {
  if t != nil {
    RNL(t.Right, ch)
    ch <- t.Value
    RNL(t.Left, ch)
  }
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  RNL(t, ch)
  close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1, ch2 := make(chan int), make(chan int)

  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for {
    v1, ok1 := <-ch1
    v2, ok2 := <-ch2

    if ok1 && ok2 && v1 != v2 {
      return false
    }
    if !ok1 || !ok2 {
      break
    }
  }

  return true
}

func main() {
  s := Same(tree.New(1), tree.New(2))
  fmt.Println(s)
  ch := make(chan int)
  go Walk(tree.New(2), ch)

  for {
    v, ok := <-ch
    if !ok {
      break
    } else {
      fmt.Println(v)
    }
  }
}
```

### sync.Mutex là một cấu trúc dữ liệu được sử dụng để đảm bảo chỉ một goroutine có thể truy cập giá trị ở 1 thời điểm để tranh xung đột - mutual exclusion.

Thư viện tiêu chuẩn của Go cung cấp tính năng loại trừ lẫn nhau - mutual exclusion với `sys.Mutex` và 2 phương thức: `Lock`, `Unlock`.

```Go
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
  mu sync.Mutex
  v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
  c.mu.Lock()
  // Lock so only one goroutine at a time can access the map c.v.
  c.v[key]++
  c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
  c.mu.Lock()
  // Lock so only one goroutine at a time can access the map c.v.
  defer c.mu.Unlock() // để đảm bảo mutex sẽ được mở khóa
  return c.v[key]
}

func main() {
  c := SafeCounter{v: make(map[string]int)}
  for i := 0; i < 1000; i++ {
    go c.Inc("somekey")
  }

  time.Sleep(time.Second)
  fmt.Println(c.Value("somekey"))
}
```
