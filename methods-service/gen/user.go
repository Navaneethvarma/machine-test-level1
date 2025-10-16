package gen

import (
    "context"
    "google.golang.org/grpc"
)

// Simple message types
type User struct {
    Id        string
    Name      string
    Email     string
    CreatedAt string
    UpdatedAt string
}

type CreateUserRequest struct {
    Name  string
    Email string
}

type GetUserRequest struct {
    Id string
}

type UpdateUserRequest struct {
    Id    string
    Name  string
    Email string
}

type DeleteUserRequest struct {
    Id string
}

type DeleteUserResponse struct {
    Success bool
}

type ListUsersRequest struct{}

type ListUsersResponse struct {
    Users []*User
}

type UserResponse struct {
    User  *User
    Error string
}

// Client interface
type UserServiceClient interface {
    CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
    GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
    UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
    DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
    ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type userServiceClient struct {
    cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
    return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
    out := new(UserResponse)
    err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
    return out, err
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
    out := new(UserResponse)
    err := c.cc.Invoke(ctx, "/user.UserService/GetUser", in, out, opts...)
    return out, err
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
    out := new(UserResponse)
    err := c.cc.Invoke(ctx, "/user.UserService/UpdateUser", in, out, opts...)
    return out, err
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
    out := new(DeleteUserResponse)
    err := c.cc.Invoke(ctx, "/user.UserService/DeleteUser", in, out, opts...)
    return out, err
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
    out := new(ListUsersResponse)
    err := c.cc.Invoke(ctx, "/user.UserService/ListUsers", in, out, opts...)
    return out, err
}

// Server interface
type UserServiceServer interface {
    CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error)
    GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
    UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error)
    DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
    ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
    mustEmbedUnimplementedUserServiceServer()
}

type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error) {
    return nil, nil
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*UserResponse, error) {
    return nil, nil
}

func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error) {
    return nil, nil
}

func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
    return nil, nil
}

func (UnimplementedUserServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
    return nil, nil
}

func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
    // Simple registration - in real scenario this would use proper gRPC registration
}
