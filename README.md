# Simple RBAC Service

## Overview

The **Simple RBAC Service** is a lightweight role-based access control (RBAC) system built with Go. It allows managing users, roles, and permissions efficiently, ensuring that only authorized users can access specific resources. This service is useful for applications that require fine-grained access control without relying on heavyweight frameworks.

## How It Works

1. Users are assigned specific roles.
2. Roles define access permissions for different resources.
3. Requests to protected resources are validated based on the user's assigned roles.

## Database Design
https://dbdiagram.io/d/rbac-610241892ecb310fc3b4c107

### Key Components of the Schema:

1. **Users**<br>
    Fields:<br>
        - **id**: Primary key, uniquely identifies each user.<br>
        - **username**: User's chosen display name.<br>
        - **role_id**: Foreign key referencing roles.id, indicating the user's role.<br>
        - **created_at**: Timestamp of when the user account was created.<br>

2. **Roles**<br>
    Fields:<br>
        - **id**: Primary key, uniquely identifies each role.<br>
        - **name**: Name of the role (e.g., admin, user).<br>
        - **description**: Description of the role's responsibilities.<br>
        - **created_at**: Timestamp of when the role was created.

3. **Permissions**<br>
    Fields:<br>
        - **id**: Primary key, uniquely identifies each permission. <br>
        - **name**: Name of the permission (e.g., read, write).<br>
        - **description**: Description of what the permission allows.<br>
        - **created_at**: Timestamp of when the permission was created.<br>

4. **Role_Permissions**<br>
    Fields:<br>
        - **role_id**: Foreign key referencing roles.id, indicating the role.<br>
        - **permission_id**: Foreign key referencing permissions.id, indicating the permission.<br>
        - **created_at**: Timestamp of when the role-permission relationship was established.

### Relationships and Functionality:

- Users to Roles:
    A many-to-one relationship where each user is assigned a single role. This is represented by the foreign key role_id in the Users table referencing roles.id.

- Roles to Permissions:
    A many-to-many relationship facilitated by the Role_Permissions table, where each role can have multiple permissions, and each permission can be assigned to multiple roles.

### Advantages of This Design:

**Simplified Permission Management**: By assigning permissions to roles rather than individual users, the system streamlines the process of managing user access. Changes in permissions can be made at the role level, automatically propagating to all users assigned to that role.

**Enhanced Security**: This approach ensures that users have access only to the resources necessary for their job functions, adhering to the principle of least privilege.

**Scalability**: As organizations grow, this design allows for easy addition of new roles and permissions without extensive modifications to the user base.

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
