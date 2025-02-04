package postgres

var (
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
