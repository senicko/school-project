package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/senicko/school-project-backend/pkg/app"
)

type UserRepo struct {
	dbPool *pgxpool.Pool
}

// NewRepo creates a new
func NewUserRepo(dbPool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		dbPool: dbPool,
	}
}

// CreateUser creates a new user.
func (ur UserRepo) Create(ctx context.Context, c app.User) (*app.User, error) {
	row := ur.dbPool.QueryRow(ctx, "INSERT INTO users (name, email, password, jokes) VALUES ($1, $2, $3, '[]'::jsonb) RETURNING *", c.Name, c.Email, c.Password)
	return scanUser(row)
}

// FindByEmail finds user by email.
func (ur UserRepo) FindByEmail(ctx context.Context, e string) (*app.User, error) {
	row := ur.dbPool.QueryRow(ctx, "SELECT * FROM users WHERE email=$1", e)
	return scanUser(row)
}

// FindById finds user by id.
func (ur UserRepo) FindByID(ctx context.Context, ID int) (*app.User, error) {
	row := ur.dbPool.QueryRow(ctx, "SELECT * FROM users WHERE id=$1", ID)
	return scanUser(row)
}

// SaveJoke saves a joke in user's joke collection.
func (ur UserRepo) SaveJoke(ctx context.Context, userID int, joke app.Joke) error {
	_, err := ur.dbPool.Exec(ctx, "UPDATE users SET jokes = jokes || $1::jsonb WHERE id = $2", joke, userID)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}

	return nil
}

// scanUser scans query row into User struct.
func scanUser(r pgx.Row) (*app.User, error) {
	user := app.NewUser()

	if err := r.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Jokes); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return user, nil
}
