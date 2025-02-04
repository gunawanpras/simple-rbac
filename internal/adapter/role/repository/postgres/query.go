package postgres

var (
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
)
