package postgres

var (
	queryCreateUser = `
		INSERT INTO users (id, name, email, password, role_id, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	queryListUsers = `
		SELECT id, name, email, role_id, created_by, created_at
		FROM users
	`
	queryGetUserByID = `
		SELECT id, name, email, role_id, created_by, created_at
		FROM users
		WHERE id = $1
	`
	queryGetUserByEmail = `
		SELECT id, name, email, role_id, created_by, created_at
		FROM users
		WHERE email = $1
	`
	queryCreateRole = `
		INSERT INTO roles (id, name, created_by, created_at)
		VALUES ($1, $2, $3, $4)
	`
	queryListRoles = `
		SELECT id, name, created_by, created_at
		FROM roles
	`
	queryGetRoleByID = `
		SELECT id, name, created_by, created_at
		FROM roles
		WHERE id = $1
	`
	queryGetRoleByName = `
		SELECT id, name, created_by, created_at
		FROM roles
		WHERE name = $1
	`
	queryCreatePermission = `
		INSERT INTO permissions (id, name, created_by, created_at)
		VALUES ($1, $2, $3, $4)
	`
	queryListPermissions = `
		SELECT id, name, description, created_by, created_at
		FROM permissions
	`
	queryGetPermissionByID = `
		SELECT id, name, description, created_by, created_at
		FROM permissions
		WHERE id = $1
	`
	queryGetPermissionByName = `
		SELECT id, name, description, created_by, created_at
		FROM permissions
		WHERE name = $1
	`
	queryCreateRolePermission = `
		INSERT INTO role_permissions (role_id, permission_id, created_by, created_at)
		VALUES ($1, $2, $3, $4)
	`
	queryListRolePermissions = `
		SELECT role_id, permission_id, created_by, created_at
		FROM role_permissions
	`
	queryGetRolePermissionByRoleID = `
		SELECT role_id, permission_id, created_by, created_at
		FROM role_permissions
		WHERE role_id = $1
	`
	queryGetRolePermissionByPermissionID = `
		SELECT role_id, permission_id, created_by, created_at
		FROM role_permissions
		WHERE role_id = $1 AND permission_id = $2
	`
)
