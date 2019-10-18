package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func main() {
	var s = "+" //[]byte 每个byte 由8个bit构成
	fmt.Println([]byte(s))
	fmt.Printf("%s\n", []byte(s))
	fmt.Printf("%d\n", []byte(s))
	fmt.Printf("%x\n", []byte(s))
	fmt.Println()

	var i int = 2
	bs, _ := json.Marshal(i)
	fmt.Println(bs)
	fmt.Printf("%x", bs)

}
