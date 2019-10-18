package main

//帮助理解redis active idle idletimeout的意义 以及理解redis链接池的实现
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
	"time"
)

var __ConnPool *redis.Pool

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := __ConnPool.Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}
func init() {
	fmt.Println("came ing")
	__ConnPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", "172.16.50.69", 6379),
				redis.DialPassword(""),
				redis.DialDatabase(0),
				redis.DialConnectTimeout(time.Second*2),
				redis.DialReadTimeout(time.Second*2),
				redis.DialWriteTimeout(time.Second*2),
			)
		},
		//TestOnBorrow: func(c redis.Conn, t time.Time) error {
		//	_, err := c.Do("PING")
		//	return err
		//},
		MaxActive:   3,
		MaxIdle:     2,
		IdleTimeout: 1 * time.Second, // TODO 调整为稍大点的值?
		Wait:        false,
	}
	go func() {
		for {
			fmt.Println("monitor",
				__ConnPool.ActiveCount(),
				__ConnPool.IdleCount())
			time.Sleep(time.Second)

		}
	}()
	//if _, err := Do("PING"); err != nil {
	//	panic(err)
	//}
}

func TestHello(t *testing.T) {
	go func() {
		conn1 := __ConnPool.Get()
		conn1.Do("RPUSH", "testkey23", 2)
		time.Sleep(2 * time.Second)

		conn1.Close()

	}()
	go func() {
		conn1 := __ConnPool.Get()
		conn1.Do("RPUSH", "testkey23", 2)
		time.Sleep(2 * time.Second)

		conn1.Close()

	}()
	time.Sleep(3 * time.Second)
	conn := __ConnPool.Get()
	time.Sleep(10 * time.Second)
	fmt.Println(conn.Do("RPUSH", "testkey23", 2))

	time.Sleep(10 * time.Second)
	fmt.Println(conn.Do("RPUSH", "testkey23", 2))
	fmt.Println(conn.Do("RPUSH", "testkey23", 2))

	conn.Close()
	fmt.Println(conn.Do("RPUSH", "testkey23", 2))
	fmt.Println(conn.Do("RPUSH", "testkey23", 2))
	__ConnPool.Close()
	time.Sleep(4 * time.Second)
	return

}

func TimeStringToUnix(s string) int64 {
	//	const testAndSet = `
	//	local value = redis.call('get', KEYS[1])
	//	if value then
	//		return 1
	//	end
	//	redis.call('setex',KEYS[1], 86400, 1)
	//	return 0
	//`
	//	// testAndSet returns 1 if key already exists,
	//	// set key if not eixsts, then returns 0.
	//	replyint, err := redis.Int(Do("eval", testAndSet, 1, "go_weiai_call:outcall:phone:20181229"))
	loc, _ := time.LoadLocation("Local")                                 //获取时区
	unixTime, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		//dlog.Error("errmsg", err.Error())
	}
	return unixTime.Unix()
}
