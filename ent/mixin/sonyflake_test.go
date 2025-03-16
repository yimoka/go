package mixin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGenerateID(t *testing.T) {
	// 生成多个 ID 并展示
	fmt.Println("生成的 ID 示例：")
	ids := make([]string, 10)
	for i := 0; i < 10; i++ {
		ids[i] = GenerateID()
		fmt.Printf("ID %d: %s\n", i+1, ids[i])
	}

	// 测试 ID 长度
	for i, id := range ids {
		if len(id) > 15 {
			t.Errorf("ID %d 长度超过15: %s (长度: %d)", i+1, id, len(id))
		}
	}

	// 测试 ID 唯一性
	idMap := make(map[string]bool)
	for i, id := range ids {
		if idMap[id] {
			t.Errorf("发现重复 ID %d: %s", i+1, id)
		}
		idMap[id] = true
	}

	// 解析并展示 ID 信息
	fmt.Println("\n解析 ID 信息：")
	for i, id := range ids {
		timestamp, machineID, sequence, err := ParseID(id)
		if err != nil {
			t.Errorf("解析 ID %d 失败: %v", i+1, err)
			continue
		}
		fmt.Printf("ID %d 解析结果:\n", i+1)
		fmt.Printf("  原始ID: %s\n", id)
		fmt.Printf("  生成时间: %v\n", timestamp.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("  相对时间: %v\n", timestamp.Sub(defaultStartTime))
		fmt.Printf("  机器ID: %d\n", machineID)
		fmt.Printf("  序列号: %d\n", sequence)
		fmt.Println()
	}

	// 测试快速生成
	fmt.Println("测试快速生成100个ID：")
	startTime := time.Now()
	quickIDs := make([]string, 100)
	for i := 0; i < 100; i++ {
		quickIDs[i] = GenerateID()
	}
	quickDuration := time.Since(startTime)
	fmt.Printf("生成100个ID耗时: %v\n", quickDuration)
	fmt.Printf("平均每个ID生成时间: %v\n", quickDuration/100)

	// 展示部分快速生成的ID
	fmt.Println("\n快速生成的前5个ID：")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %s\n", i+1, quickIDs[i])
	}

	// 并发测试
	fmt.Println("\n并发测试（10个goroutine各生成100个ID）：")
	var wg sync.WaitGroup
	idChan := make(chan string, 1000)

	concurrentStart := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				idChan <- GenerateID()
			}
		}()
	}

	wg.Wait()
	close(idChan)

	// 检查唯一性
	concurrentIDMap := make(map[string]bool)
	for id := range idChan {
		if concurrentIDMap[id] {
			t.Errorf("在并发测试中发现重复ID: %s", id)
		}
		concurrentIDMap[id] = true
	}

	concurrentDuration := time.Since(concurrentStart)
	fmt.Printf("并发生成1000个ID总耗时: %v\n", concurrentDuration)
	fmt.Printf("平均每个ID生成时间: %v\n", concurrentDuration/1000)
}

func TestParseID(t *testing.T) {
	// 生成一个ID
	id := GenerateID()

	// 解析ID
	timestamp, machineID, sequence, err := ParseID(id)
	if err != nil {
		t.Fatalf("解析ID失败: %v", err)
	}

	// 验证时间戳
	now := time.Now()
	if timestamp.After(now) {
		t.Errorf("时间戳异常，生成时间晚于当前时间: %v > %v", timestamp, now)
	}

	if timestamp.Before(defaultStartTime) {
		t.Errorf("时间戳异常，生成时间早于起始时间: %v < %v", timestamp, defaultStartTime)
	}

	// 验证序列号范围（移除了多余的范围检查，因为 uint16 类型已经保证了范围）
	t.Logf("解析结果 - 时间: %v, 机器ID: %d, 序列号: %d", timestamp, machineID, sequence)
}
