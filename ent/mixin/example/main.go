package main

import (
	"fmt"
	"time"

	"github.com/yimoka/go/ent/mixin"
)

func main() {
	fmt.Println("生成 5 个 ID 并解析：")
	for i := 0; i < 5; i++ {
		// 生成 ID
		id := mixin.GenerateID()

		// 解析 ID
		timestamp, machineID, sequence, _ := mixin.ParseID(id)

		fmt.Printf("\nID %d:\n", i+1)
		fmt.Printf("  值: %s\n", id)
		fmt.Printf("  长度: %d\n", len(id))
		fmt.Printf("  生成时间: %v\n", timestamp.Format(time.RFC3339Nano))
		fmt.Printf("  机器ID: %d\n", machineID)
		fmt.Printf("  序列号: %d\n", sequence)

		// 短暂等待，让序列号变化更明显
		time.Sleep(time.Millisecond)
	}
}
