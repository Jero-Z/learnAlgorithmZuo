package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var ctx = context.Background()

// CheckInSystem 是签到系统
type CheckInSystem struct {
	redisClient *redis.Client
}

// NewCheckInSystem 创建一个签到系统实例
func NewCheckInSystem(redisClient *redis.Client) *CheckInSystem {
	return &CheckInSystem{redisClient: redisClient}
}

// CheckIn 签到
func (cs *CheckInSystem) CheckIn(userID string) error {
	today := time.Now()
	key := fmt.Sprintf("checkin:%s:%d", userID, today.Year())

	// 设置今天的位为 1，表示签到
	_, err := cs.redisClient.SetBit(ctx, key, int64(today.Day()-1), 1).Result()
	if err != nil {
		return err
	}

	// 记录签到历史（有序集合）
	score := float64(today.Unix())
	_, err = cs.redisClient.ZAdd(ctx, key+":history", &redis.Z{Member: today.Format("2006-01-02"), Score: score}).Result()
	return err
}

// HasCheckedIn 检查用户今天是否签到
func (cs *CheckInSystem) HasCheckedIn(userID string) (bool, error) {
	today := time.Now()
	key := fmt.Sprintf("checkin:%s:%d", userID, today.Year())

	// 检查今天的位是否为 1，表示已签到
	bit, err := cs.redisClient.GetBit(ctx, key, int64(today.Day()-1)).Result()
	return bit == 1, err
}

// IsContinuousCheckIn 检查用户是否连续签到
func (cs *CheckInSystem) IsContinuousCheckIn(userID string, days int) (bool, error) {
	historyKey := fmt.Sprintf("checkin:%s:%d:history", userID, time.Now().Year())

	// 获取用户签到历史（有序集合）
	history, err := cs.redisClient.ZRevRange(ctx, historyKey, 0, int64(days-1)).Result()
	if err != nil {
		return false, err
	}

	// 检查最近连续 days 天的签到情况
	for i := 0; i < days; i++ {
		expectedDate := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		if i < len(history) && history[i] == expectedDate {
			continue
		}
		return false, nil
	}

	return true, nil
}

func main() {
	// 两个协程打印奇数和偶数

	s := make(chan struct{})
	d := make(chan struct{})
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		defer func() {
			waitGroup.Done()
		}()
		for i := 0; i <= 100; i++ {
			time.Sleep(1 * time.Microsecond)
			select {
			case <-s:
				if i%2 == 0 {
					fmt.Println("g1", i)
				}
				d <- struct{}{}
			}
		}
	}()
	waitGroup.Add(1)
	go func() {

		defer func() {
			waitGroup.Done()
		}()
		time.Sleep(1 * time.Microsecond)

		for i := 0; i <= 100; i++ {
			select {
			case <-d:
				if i%2 != 0 {
					fmt.Println("g2", i)
				}
				if i != 100 {
					s <- struct{}{}
				}

			}
		}
	}()
	s <- struct {
	}{}
	waitGroup.Wait()
	fmt.Println("打印结束")
	return
	// 使用本地 Redis，端口默认为 6379
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	checkinSystem := NewCheckInSystem(redisClient)

	// 示例：用户1连续签到3天
	userID := "user1"
	for i := 0; i < 3; i++ {
		checkinSystem.CheckIn(userID)
	}

	// 检查用户1今天是否签到
	result, err := checkinSystem.HasCheckedIn(userID)
	if err != nil {
		fmt.Println("Error checking if user has checked in:", err)
		return
	}
	fmt.Printf("User %s has checked in today: %t\n", userID, result)

	// 模拟用户1断签
	// 用户1断签后再连续签到2天
	checkinSystem.CheckIn(userID)
	checkinSystem.CheckIn(userID)

	// 检查用户1最近是否连续签到3天
	continuousResult, err := checkinSystem.IsContinuousCheckIn(userID, 3)
	if err != nil {
		fmt.Println("Error checking if user has continuously checked in:", err)
		return
	}
	fmt.Printf("User %s has continuously checked in for the last 3 days: %t\n", userID, continuousResult)
}
