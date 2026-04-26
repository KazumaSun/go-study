# 009 正解例

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
)

type UserRepository struct{ db *sql.DB }

func (r *UserRepository) TransferPoint(ctx context.Context, fromID, toID, point int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "UPDATE users SET point = point - ? WHERE id = ?", point, fromID); err != nil {
		return fmt.Errorf("decrease point: %w", err)
	}
	if _, err := tx.ExecContext(ctx, "UPDATE users SET point = point + ? WHERE id = ?", point, toID); err != nil {
		return fmt.Errorf("increase point: %w", err)
	}
	return tx.Commit()
}
```
