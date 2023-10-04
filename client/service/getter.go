package service

import (
	"context"
	"fmt"
)

type Getter interface {
	GetUniqueKeys(ctx context.Context, userID int) (map[string][]string, error)
}

func (g *GetService) GetUniqueKeys(ctx context.Context, userID int) (map[string][]string, error) {
	const op = "service.getter.GetUniqueKeys"
	uniqueKeys, err := g.get.GetUniqueKeys(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return uniqueKeys, nil
}
