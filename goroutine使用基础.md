# goroutine使用基础
简单介绍go语言中并发的使用
## 基础并发使用
通过go关键字实现并发
```Go
go your_func(a, b)
```

## 并发
我们在实际coding中，经常使用goroutine处理多路并发操作，比如并发抓取网络资源（url），
需要在for（获取其他）循环语句中处理：

```go

for i:=1; i<100; i++ {
   go getNetSources(i)
}
```

此时，主进程并不知道并发的协程何时结束，需要阻主进程直到所有协程都并发完成再退出。
实现方式主要有3中：

### 1、time.Sleep()
此种方式简单，粗暴，可能产生系列问题，最好不用。

### 2、通过channel、select
```go
func gofunc(ch chan int, qt chan bool){
	x := 100
    for {
        select {
            case ch<-x:
				x += 1
            case <-qt:
                fmt.Println("finish.")
                return
        }
    }
}

func main(){
    ch := make(chan int)
    qt := make(chan bool, 1)
	go func(){
		for i:=1;i<10;i++ {
			fmt.Println(i, <-ch)
		}
		qt<-true
	}()
    gofunc(ch, qt)
    close(ch)
    close(qt)
}

```

### 3、通过golang提供的sync库中的WaitGroup
```go
    func gofunc(i int, rs chan int){
        fmt.Println(i, "hw f")
        rs<-100 + i
    }

    func main(){
        var wg sync.WaitGroup
        rs := make(chan int)
        for i:=1; i<10; i++ {
            wg.Add(1)
            go func(x int){
                gofunc(x, rs)
                wg.Done()
            }(i)
            fmt.Println(<-rs)
        }

        close(rs)
        wg.Wait()
    }
```
* 每一个协程之前Add(1)
* 协程执行玩之后Done()
* 最后通过Wait()阻塞主进程，等待所有协程都完成后退出

**说明**
其中主进程中通过chan rs进行结果回传，
```go
gofunc(x, rs)
```
协程执行完之后必须处理channel中的数据，否则会报
```go
panic(0x4b8e60, 0xc82006c060)
	/usr/lib/go-1.6/src/runtime/panic.go:481 +0x3e6
```
错误

### 其他
也有通过select{}方法进行并发等待的方式，但会一直阻塞住进行等待，对需要并发执行操作的脚本不太合适。
如果有数据传输最好使用select case 方式
