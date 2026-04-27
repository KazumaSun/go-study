# 009 正解例（課題別）

## 課題1: トランザクション境界

```go
func (r *UserRepository) TransferPoint(ctx context.Context, fromID, toID, point int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "UPDATE users SET point=point-? WHERE id=?", point, fromID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "UPDATE users SET point=point+? WHERE id=?", point, toID); err != nil {
		return err
	}
	return tx.Commit()
}
```

## 課題2: N+1改善の方向

```sql
-- bad: 1件ずつSELECT
SELECT * FROM users WHERE id = ?;

-- good: まとめて取得
SELECT * FROM users WHERE id IN (?, ?, ?);
```

## 課題3: migration運用例

```text
001_create_users.up.sql
001_create_users.down.sql
002_add_index_users_email.up.sql
002_add_index_users_email.down.sql
```

```sql
CREATE INDEX idx_users_email ON users(email);
```
