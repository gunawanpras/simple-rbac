package postgres

var (
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
)
