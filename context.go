//需要看
//package main
//
//import (
//	"context"
//	"fmt"
//	"net/http"
//	"runtime"
//	"time"
//)
//
//func slow(ctx context.Context) {
//	c := make(chan int, 1)
//	go func() {
//		time.Sleep(3 * time.Second)
//		c <- 2
//	}()
//	select {
//
//	case <-ctx.Done():
//		fmt.Println("done")
//		return
//	case <-c:
//		fmt.Println("ccc")
//		return
//	}
//	fmt.Println("slow")
//	time.Sleep(5 * time.Second)
//	fmt.Println("slow2")
//}
//func main() {
//	runtime.GC()
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	fmt.Println(ctx.Value("ss"), <-ctx.Done())
//	cancel()
//	fmt.Println(ctx.Err())
//
//	defer cancel()
//	req, err := http.NewRequest("GET", "http://127.0.0.1:8200", nil)
//	res, err := http.DefaultClient.Do(req.WithContext(ctx))
//	fmt.Println(ctx.Err())
//
//	fmt.Println(res, err, "res")
//
//	slow(ctx)
//	ssss := make(chan int, 1)
//	ddd := make(map[int64]chan int, 0)
//	ddd[0] = ssss
//	ddd[0] <- 1
//
//	select {
//	case v, ok := <-ssss:
//
//		fmt.Println(v, ok, "ss")
//
//	}
//}

package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

// Done is provided for use in select statements:
//
//  // Stream generates values with DoSomething and sends them to out
//  // until DoSomething returns an error or ctx.Done is closed.
//  func Stream(ctx context.Context, out chan<- Value) error {
//  	for {
//  		v, err := DoSomething(ctx)
//  		if err != nil {
//  			return err
//  		}
//  		select {
//  		case <-ctx.Done():
//  			return ctx.Err()
//  		case out <- v:
//  		}
//  	}
//  }
//
func someHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //父ctx关闭子ctx也关闭
	ctx1, cancel1 := context.WithTimeout(ctx, 12*time.Second)
	go doStuff(ctx)
	go doStuff1(ctx1)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
	cancel1()
	time.Sleep(3 * time.Second) //避免main函数退出影响

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff1(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done1")
			return
		default:
			logg.Printf("work1")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}
