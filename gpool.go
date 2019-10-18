package main

//连接池管理 go serverhttp
import (
	"io/ioutil"
	"net/http"
	"runtime"

	"fmt"
	"github.com/Jeffail/tunny"
	"time"
)

func main() {
	numCPUs := runtime.NumCPU()
	fmt.Println(numCPUs)
	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var result []byte

		// TODO: Something CPU heavy with payload
		fmt.Println("gid2", runtime.GetGoroutineId())
		time.Sleep(6 * time.Second)
		result = []byte("string")
		return result
	})

	defer pool.Close()
	fmt.Println("gid", runtime.GetGoroutineId())
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("gid3", runtime.GetGoroutineId())

		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		// Funnel this work into our pool. This call is synchronous and will
		// block until the job is completed.
		result, err := pool.ProcessTimed(input, time.Second*5)
		if err == tunny.ErrJobTimedOut {
			http.Error(w, "Request timed out", http.StatusRequestTimeout)
		}
		w.Write(result.([]byte))
	})

	if err := http.ListenAndServe(":8299", nil); err != nil {
		fmt.Println(err)
	}
}
