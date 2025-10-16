package postgres

import (
    "context"
    "database/sql"
    "fmt"
    "user-service/internal/models"
    _ "github.com/lib/pq"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository() *UserRepository {
    connStr := "user=postgres password=password dbname=userdb sslmode=disable host=postgres port=5432"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }
    
    // Test connection
    if err := db.Ping(); err != nil {
        panic(fmt.Sprintf("Failed to ping database: %v", err))
    }
    
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
    query := "INSERT INTO users (name, email) VALUES (, ) RETURNING id, created_at, updated_at"
    err := r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return fmt.Errorf("failed to create user: %v", err)
    }
    return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
    user := &models.User{}
    query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = "
    err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("failed to get user: %v", err)
    }
    return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
    query := "UPDATE users SET name = , email = , updated_at = CURRENT_TIMESTAMP WHERE id = "
    result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
    if err != nil {
        return fmt.Errorf("failed to update user: %v", err)
    }
    
    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %v", err)
    }
    if rows == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
    query := "DELETE FROM users WHERE id = "
    result, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("failed to delete user: %v", err)
    }
    
    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %v", err)
    }
    if rows == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]*models.User, error) {
    query := "SELECT id, name, email, created_at, updated_at FROM users ORDER BY created_at DESC"
    rows, err := r.db.QueryContext(ctx, query)
    if err != nil {
        return nil, fmt.Errorf("failed to list users: %v", err)
    }
    defer rows.Close()
    
    var users []*models.User
    for rows.Next() {
        user := &models.User{}
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            return nil, fmt.Errorf("failed to scan user: %v", err)
        }
        users = append(users, user)
    }
    
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("rows error: %v", err)
    }
    
    return users, nil
}
