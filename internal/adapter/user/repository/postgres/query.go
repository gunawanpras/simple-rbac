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
)
