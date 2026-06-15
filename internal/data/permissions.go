package data

import (
	"context"
	"database/sql"
	"slices"
	"time"
)

type Permissions []string

func (p Permissions) Include(code string) bool {
	return slices.Contains(p, code)
}

type PermissionModel struct {
	DB *sql.DB
}

func (m PermissionModel) GetAllForUser(userID int) (Permissions, error) {

	query := `

	select * 
	
	from permissions as p
	
	inner join users_permissions as up 	on p.id = up.permission_id
	inner join users as u 				on u.id = up.user_id

	where u.id = $1

	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var permissions Permissions

	for rows.Next() {

		var permission string

		//shadowing
		err := rows.Scan(&permission)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
