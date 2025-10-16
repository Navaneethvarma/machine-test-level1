package service

import (
    "context"
    "log"
    "methods-service/internal/client"
    "sync"
    "time"
)

type MethodsService struct {
    userClient     *client.UserClient
    method1Running bool
    method1Mutex   sync.Mutex
    method1Cond    *sync.Cond
}

func NewMethodsService(userClient *client.UserClient) *MethodsService {
    ms := &MethodsService{
        userClient: userClient,
    }
    ms.method1Cond = sync.NewCond(&ms.method1Mutex)
    return ms
}

func (ms *MethodsService) ProcessMethod1(ctx context.Context, waitTime int) ([]string, error) {
    ms.method1Mutex.Lock()
    for ms.method1Running {
        log.Println("Method 1: Waiting for previous task to complete...")
        ms.method1Cond.Wait()
    }
    ms.method1Running = true
    ms.method1Mutex.Unlock()

    defer func() {
        ms.method1Mutex.Lock()
        ms.method1Running = false
        ms.method1Cond.Broadcast()
        ms.method1Mutex.Unlock()
        log.Println("Method 1: Task completed, released lock")
    }()

    log.Println("Method 1: Starting sequential processing...")
    return ms.processCommon(ctx, waitTime)
}

func (ms *MethodsService) ProcessMethod2(ctx context.Context, waitTime int) ([]string, error) {
    log.Println("Method 2: Starting parallel processing...")
    return ms.processCommon(ctx, waitTime)
}

func (ms *MethodsService) processCommon(ctx context.Context, waitTime int) ([]string, error) {
    // Get real user names from user-service via gRPC
    names, err := ms.userClient.GetUserNames(ctx)
    if err != nil {
        log.Printf("Error getting user names: %v", err)
        // Fallback to mock data if gRPC fails
        names = []string{"John Doe", "Jane Smith", "Bob Johnson"}
    }
    
    log.Printf("Sleeping for %d seconds...", waitTime)
    select {
    case <-time.After(time.Duration(waitTime) * time.Second):
    case <-ctx.Done():
        return nil, ctx.Err()
    }

    log.Printf("Returning %d user names", len(names))
    return names, nil
}

func (ms *MethodsService) Close() {
    if ms.userClient != nil {
        ms.userClient.Close()
    }
}
