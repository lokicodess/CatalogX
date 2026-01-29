package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lokicodess/CatalogX/internal/domain"
)

type PostgresUserRepository struct {
	DB *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (p *PostgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	stmt := `
	INSERT INTO USERS(name, email, password_hash)
	VALUES ($1,$2,$3)
	RETURNING id, created_at, updated_at
	`
	err := p.DB.QueryRow(ctx, stmt, user.Name, user.Email, user.PasswordHash).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	stmt := `
SELECT id, name, email, password_hash, created_at, updated_at
FROM users
WHERE email = $1;
`

	err := p.DB.QueryRow(ctx, stmt, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
