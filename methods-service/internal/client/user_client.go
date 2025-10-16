package client

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

// Types matching user-service
type User struct {
    Id        string
    Name      string
    Email     string
    CreatedAt string
    UpdatedAt string
}

type ListUsersRequest struct{}

type ListUsersResponse struct {
    Users []*User
}

type UserClient struct {
    conn *grpc.ClientConn
}

func NewUserClient(addr string) (*UserClient, error) {
    conn, err := grpc.Dial(addr, 
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithBlock())
    if err != nil {
        return nil, err
    }
    return &UserClient{conn: conn}, nil
}

func (c *UserClient) GetUserNames(ctx context.Context) ([]string, error) {
    // Use reflection to call ListUsers method
    out := new(ListUsersResponse)
    err := c.conn.Invoke(ctx, "/UserServiceServer/ListUsers", &ListUsersRequest{}, out)
    if err != nil {
        return nil, err
    }
    
    var names []string
    for _, user := range out.Users {
        names = append(names, user.Name)
    }
    return names, nil
}

func (c *UserClient) Close() {
    if c.conn != nil {
        c.conn.Close()
    }
}
