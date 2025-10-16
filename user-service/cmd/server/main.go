package main

import (
    "context"
    "log"
    "net"
    "time"

    "user-service/internal/models"
    "user-service/internal/repository/postgres"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/reflection"
)

// Protobuf message types (temporary until generated)
type CreateUserRequest struct {
    Name  string
    Email string
}

type GetUserRequest struct {
    Id string
}

type UserResponse struct {
    Id    string
    Name  string
    Email string
}

// UserServiceServer implements the gRPC service
type UserServiceServer struct {
    repo *postgres.UserRepository
}

// CreateUser creates a new user
func (s *UserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
    log.Printf("Creating user: %s, %s", req.Name, req.Email)
    
    user := &models.User{
        Name:  req.Name,
        Email: req.Email,
    }
    
    // In real implementation, this would call the repository
    // err := s.repo.CreateUser(ctx, user)
    // if err != nil {
    //     return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
    // }
    
    // Simulate successful creation
    user.ID = "1"
    
    return &UserResponse{
        Id:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }, nil
}

// GetUser retrieves a user by ID
func (s *UserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
    log.Printf("Getting user: %s", req.Id)
    
    // In real implementation, this would call the repository
    // user, err := s.repo.GetUserByID(ctx, req.Id)
    // if err != nil {
    //     return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
    // }
    
    // Simulate user found
    user := &models.User{
        ID:    req.Id,
        Name:  "John Doe",
        Email: "john@example.com",
    }
    
    return &UserResponse{
        Id:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }, nil
}

func main() {
    log.Printf("Starting User Service...")
    
    // Initialize database repository
    // repo, err := postgres.NewUserRepository()
    // if err != nil {
    //     log.Fatalf("Failed to initialize repository: %v", err)
    // }
    
    // Create gRPC server
    server := grpc.NewServer()
    
    // Create service instance (using nil repo for now)
    userService := &UserServiceServer{repo: nil}
    
    // Register service (commented until gRPC code generation is complete)
    // gen.RegisterUserServiceServer(server, userService)
    
    // Enable reflection for testing
    reflection.Register(server)

    // Start server
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen on :50051: %v", err)
    }

    log.Printf("? User Service started successfully on :50051")
    log.Printf("? gRPC reflection enabled")
    log.Printf("? Ready to accept requests")
    
    if err := server.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
