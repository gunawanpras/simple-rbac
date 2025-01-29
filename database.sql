-- RBAC Database Schema Design

-- Table: Roles
CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(100) DEFAULT 'SYSTEM',
    updated_at TIMESTAMP NULL,
    updated_by VARCHAR(100) NULL
);
CREATE INDEX idx_roles_name ON roles(name);

-- Table: Permissions
CREATE TABLE permissions (
    id UUID PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(100) DEFAULT 'SYSTEM',
    updated_at TIMESTAMP NULL,
    updated_by VARCHAR(100) NULL
);
CREATE INDEX idx_permissions_name ON permissions(name);

-- Table: Users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(100) DEFAULT 'SYSTEM',
    updated_at TIMESTAMP NULL,
    updated_by VARCHAR(100) NULL,
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE SET NULL
);
CREATE INDEX idx_users_email ON users(email);

-- Table: Role_Permissions (Join Table)
CREATE TABLE role_permissions (
    role_id UUID NOT NULL,
    permission_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(100) DEFAULT 'SYSTEM',
    PRIMARY KEY (role_id, permission_id),
    CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
    CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE
);
CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);

-- Example Data Insertion

-- Roles
INSERT INTO roles (id, name) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'Admin'),
('550e8400-e29b-41d4-a716-446655440003', 'User');

-- Permissions
INSERT INTO permissions (id, name, description) VALUES
('550e8400-e29b-41d4-a716-446655440002', 'Manage Users', 'Permission to manage users'),
('550e8400-e29b-41d4-a716-446655440004', 'View Reports', 'Permission to view reports'),
('550e8400-e29b-41d4-a716-446655440005', 'Edit Profile', 'Permission to edit own profile'),
('550e8400-e29b-41d4-a716-446655440006', 'Delete Content', 'Permission to delete content');

-- Users
INSERT INTO users (id, name, email, password, role_id) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'John Doe', 'johndoe@example.com', 'hashed_password', '550e8400-e29b-41d4-a716-446655440001'),
('550e8400-e29b-41d4-a716-446655440007', 'Jane Smith', 'janesmith@example.com', 'secure_password', '550e8400-e29b-41d4-a716-446655440003');

-- Role_Permissions
INSERT INTO role_permissions (role_id, permission_id) VALUES
('550e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440002'),
('550e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440004'),
('550e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440005'),
('550e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440006');