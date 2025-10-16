# Machine Test Level 1 - Two Microservices Application

## Project Overview
This project implements two microservices in Go with PostgreSQL, Redis, gRPC communication, Docker containerization, and Kubernetes deployment.

## Timebox & Submission
- **Allotted Time**: 2 days (48 hours) from assignment receipt
- **Submission Date**: 2025-10-17
- **Submission Format**: Single ZIP file with complete code and instructions

## Technology Versions Used
- **Go**: 1.25
- **Docker**: 20.10+
- **Kubernetes**: 1.28+
- **kubectl**: 1.28+
- **Minikube**: 1.28+ (recommended for local testing)

## Assumptions & Notes
1. **Database**: Uses ephemeral PostgreSQL instance for development (data not persisted)
2. **Redis**: Uses in-memory Redis cache for development
3. **Secrets**: No hard-coded secrets - all credentials use environment variables
4. **Networking**: Services use default ports (50051 for gRPC, 8080 for HTTP)
5. **Development**: Focus on architecture and core functionality over production hardening
6. **Original Work**: This is original work created specifically for this machine test

## Architecture
- **User Service**: gRPC service for user CRUD operations with PostgreSQL and Redis caching
- **Methods Service**: HTTP service with sequential/parallel processing that calls User Service via gRPC

## Technology Stack
- **Language**: Go 1.25
- **Database**: PostgreSQL 13
- **Cache**: Redis 6
- **Communication**: gRPC
- **Containerization**: Docker with multi-stage builds
- **Orchestration**: Kubernetes

## Project Structure
