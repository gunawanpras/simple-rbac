# Simple RBAC Service

## Overview

The **Simple CRUD RBAC Service** is a lightweight role-based access control (RBAC) system built with Go. It allows managing users, roles, and permissions efficiently, ensuring that only authorized users can access specific resources. This service is useful for applications that require fine-grained access control without relying on heavyweight frameworks.

## How It Works

1. Users are assigned specific roles.
2. Roles define access permissions for different resources.
3. Requests to protected resources are validated based on the user's assigned roles.

## Database Design
https://dbdiagram.io/d/rbac-610241892ecb310fc3b4c107

## Getting Started

### Requirements

- Go 1.23 or later.
- `make` installed on your system.
- Docker for containerization

### How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/gunawanpras/simple-rbac.git
   cd simple-rbac
   ```

2. Build and run the service:
   ```bash
   make all
   ```
   The service will start on port `8080` by default.

## Example API Usage

### Create new user with new role
```bash
curl --location 'localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Jono",
    "email": "jono@example.id",
    "password": "12345678",    
    "role_id": "550e8400-e29b-41d4-a716-446655440001"
}'
```

## Documentations
- [redoc](http://localhost:8080/docs) (localhost)
- [postman](https://documenter.getpostman.com/view/819887/2sAYQiBnzx) (online)
- [swaggerhub](https://app.swaggerhub.com/apis-docs/gunawanpras/role-based_access_control_api/1.0.0) (online)

## License

GNU Public License 3.0 - The GPL 3.0 license requires that any software that uses or modifies the licensed code must also be made available under the same GPL 3.0 license, ensuring that the software remains free and open-source, with users having the freedom to use, modify, and distribute it.
