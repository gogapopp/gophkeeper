package sqlite

import (
	"context"
	"fmt"
)

func (r *Repository) GetUniqueKeys(ctx context.Context, userID int) (map[string][]string, error) {
	const op = "sqlite.get.GetUniqueKeys"
	tables := []string{"textdata", "binarydata", "carddata"}
	keys := make(map[string][]string)
	for _, table := range tables {
		rows, err := r.db.QueryContext(ctx, fmt.Sprintf("SELECT unique_key FROM %s WHERE user_id=?1", table), userID)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
		defer rows.Close()
		var uniqueKeys []string
		for rows.Next() {
			var uniqueKey string
			if err := rows.Scan(&uniqueKey); err != nil {
				return nil, fmt.Errorf("%s: %s", op, err)
			}
			uniqueKeys = append(uniqueKeys, uniqueKey)
		}
		keys[table] = uniqueKeys
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
	}
	return keys, nil
}
