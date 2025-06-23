package utils


import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"encoding/json"
	"time"
	"fmt"
)



type Job struct {
	StudentEmail string `json:"student_mail"`
	CourseID     string `json:"course_id"`
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		DB:   0,
	})
	return rdb;
}


func StartWorker(rdb *redis.Client) {
	fmt.Println("Mail worker started in the background....")
	for {
		ctx := context.Background();
		result, err := rdb.BRPop(ctx, 5*time.Second, "mailing_queue").Result()
		if err == redis.Nil || len(result) < 2 {
			continue
		} else if err != nil {
			fmt.Println("Redis error:", err)
			continue
		}
		var job Job
		if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
			fmt.Println("JSON error:", err)
			continue
		}
		if err := SendMail(job.StudentEmail, job.CourseID); err != nil {
			fmt.Println("Email send failed:", err)
		}
	}
}


func PushtoQ(rdb *redis.Client, email, courseID string) error {
	job := map[string]string{
		"student_mail": email,
		"course_id":     courseID,
	}
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return rdb.LPush(context.Background(), "mailing_queue", data).Err()
}


