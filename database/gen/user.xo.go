package gen

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

// User represents a row from 'public.users'.
type User struct {
	ID           uuid.UUID    `json:"id"`            // id
	Name         string       `json:"name"`          // name
	Email        string       `json:"email"`         // email
	PasswordHash string       `json:"password_hash"` // password_hash
	CreatedAt    sql.NullTime `json:"created_at"`    // created_at
	UpdatedAt    sql.NullTime `json:"updated_at"`    // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [User] exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the [User] has been marked for deletion
// from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the [User] to the database.
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.users (` +
		`id, name, email, password_hash, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`
	// run
	logf(sqlstr, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Update updates a [User] in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.users SET ` +
		`name = $1, email = $2, password_hash = $3, created_at = $4, updated_at = $5 ` +
		`WHERE id = $6`
	// run
	logf(sqlstr, u.Name, u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.Name, u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt, u.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [User] to the database.
func (u *User) Save(ctx context.Context, db DB) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for [User].
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.users (` +
		`id, name, email, password_hash, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, email = EXCLUDED.email, password_hash = EXCLUDED.password_hash, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the [User] from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.users ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UsersByEmail retrieves a row from 'public.users' as a [User].
//
// Generated from index 'idx_users_email'.
func UsersByEmail(ctx context.Context, db DB, email string) ([]*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, email, password_hash, created_at, updated_at ` +
		`FROM public.users ` +
		`WHERE email = $1`
	// run
	logf(sqlstr, email)
	rows, err := db.QueryContext(ctx, sqlstr, email)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*User
	for rows.Next() {
		u := User{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UserByEmail retrieves a row from 'public.users' as a [User].
//
// Generated from index 'users_email_key'.
func UserByEmail(ctx context.Context, db DB, email string) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, email, password_hash, created_at, updated_at ` +
		`FROM public.users ` +
		`WHERE email = $1`
	// run
	logf(sqlstr, email)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// UserByID retrieves a row from 'public.users' as a [User].
//
// Generated from index 'users_pkey'.
func UserByID(ctx context.Context, db DB, id uuid.UUID) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, email, password_hash, created_at, updated_at ` +
		`FROM public.users ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}
