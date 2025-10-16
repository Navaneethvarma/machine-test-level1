package service

import (
    "context"
    "user-service/internal/models"
    "user-service/internal/repository/postgres"
    "user-service/gen"
)

type UserService struct {
    gen.UnimplementedUserServiceServer
    repo *postgres.UserRepository
}

func NewUserService(repo *postgres.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.UserResponse, error) {
    user := &models.User{
        Name:  req.Name,
        Email: req.Email,
    }
    
    if err := s.repo.CreateUser(ctx, user); err != nil {
        return &gen.UserResponse{Error: err.Error()}, nil
    }
    
    return &gen.UserResponse{
        User: &gen.User{
            Id:        user.ID,
            Name:      user.Name,
            Email:     user.Email,
            CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
            UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
        },
    }, nil
}

func (s *UserService) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.UserResponse, error) {
    user, err := s.repo.GetUserByID(ctx, req.Id)
    if err != nil {
        return &gen.UserResponse{Error: err.Error()}, nil
    }
    if user == nil {
        return &gen.UserResponse{Error: "user not found"}, nil
    }
    
    return &gen.UserResponse{
        User: &gen.User{
            Id:        user.ID,
            Name:      user.Name,
            Email:     user.Email,
            CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
            UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
        },
    }, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UserResponse, error) {
    user := &models.User{
        ID:    req.Id,
        Name:  req.Name,
        Email: req.Email,
    }
    
    if err := s.repo.UpdateUser(ctx, user); err != nil {
        return &gen.UserResponse{Error: err.Error()}, nil
    }
    
    // Fetch updated user
    updatedUser, err := s.repo.GetUserByID(ctx, req.Id)
    if err != nil {
        return &gen.UserResponse{Error: err.Error()}, nil
    }
    
    return &gen.UserResponse{
        User: &gen.User{
            Id:        updatedUser.ID,
            Name:      updatedUser.Name,
            Email:     updatedUser.Email,
            CreatedAt: updatedUser.CreatedAt.Format("2006-01-02 15:04:05"),
            UpdatedAt: updatedUser.UpdatedAt.Format("2006-01-02 15:04:05"),
        },
    }, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {
    if err := s.repo.DeleteUser(ctx, req.Id); err != nil {
        return &gen.DeleteUserResponse{Success: false}, nil
    }
    return &gen.DeleteUserResponse{Success: true}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *gen.ListUsersRequest) (*gen.ListUsersResponse, error) {
    users, err := s.repo.ListUsers(ctx)
    if err != nil {
        return nil, err
    }
    
    var pbUsers []*gen.User
    for _, user := range users {
        pbUsers = append(pbUsers, &gen.User{
            Id:        user.ID,
            Name:      user.Name,
            Email:     user.Email,
            CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
            UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
        })
    }
    
    return &gen.ListUsersResponse{Users: pbUsers}, nil
}
