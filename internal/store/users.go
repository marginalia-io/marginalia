package store

import (
	"context"
	"database/sql"
	"fmt"
)

// HasUsers reports whether at least one user account exists. It is used to
// determine whether first-run setup (onboarding) has been completed.
func HasUsers(ctx context.Context, db *sql.DB) (bool, error) {
	var exists bool
	err := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users)").Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("query users exist: %w", err)
	}
	return exists, nil
}
