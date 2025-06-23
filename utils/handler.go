package utils

import (
	"github.com/redis/go-redis/v9"
	"gofr.dev/pkg/gofr"
)

func HandleSend(rdb *redis.Client) func(c *gofr.Context) (interface{}, error) {
	return func(c *gofr.Context) (interface{}, error) {
		var req struct {
			StudentMail string `json:"student_mail"`
			CourseID    string `json:"course_id"`
		}

		if err := c.Bind(&req); err != nil {
			return nil, err
		}

		if err := PushtoQ(rdb, req.StudentMail, req.CourseID); err != nil {
			return nil, err
		}

		return map[string]string{"message": "Mail enqueued successfully!"}, nil
	}
}
