package actions

import (
	"context"
)

func SetPower(ctx context.Context, address string, status bool) error {
	return nil
}

func GetPower(ctx context.Context, address string) (bool, error) {
	return false, nil
}
