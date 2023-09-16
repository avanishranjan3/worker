package Tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const TaskType = "task:user"

func HandleUserdtl(ctx context.Context, t *asynq.Task) error {
	var users []map[string]interface{}
	err := json.Unmarshal(t.Payload(), &users)
	if err != nil {
		fmt.Println("Unmarshaling Error! ", err)
		return err
	}
	for i := range users {
		fmt.Println("Users: ", users[i])
	}
	return nil
}
