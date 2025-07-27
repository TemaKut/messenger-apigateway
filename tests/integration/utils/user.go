package utils

import (
	"fmt"
	"github.com/google/uuid"
)

func RandomUserEmail() string {
	return fmt.Sprintf("test.messenger.%s@email.com", uuid.NewString())
}
