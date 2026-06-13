package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"time"

	"greenlight.alexarmenta.net/internal/validator"
)

const (
	ScopeActivation = "activation"
)

type Token struct {
	Plaintext string
	Hash      []byte
	UserID    int
	Expiry    time.Time
	Scope     string
}

func generateToken(userID int, ttl time.Duration, scope string) Token {

	token := Token{
		Plaintext: rand.Text(),
		UserID:    userID,
		Expiry:    time.Now().Add(ttl),
		Scope:     scope,
	}

	array := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = array[:]

	return token

}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {

	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")

}

type TokenModel struct {
	DB *sql.DB
}

func (m TokenModel) New(userID int, ttl time.Duration, scope string) (Token, error) {

	token := generateToken(userID, ttl, scope)

	err := m.insert(token)

	return token, err

}

func (m TokenModel) insert(token Token) error {
	query := `
	insert into tokens 	(hash, user_id, expiry, scope)
	values 				($1,$2,$3,$4)
	`

	//solo guardas el hash y el otro?

	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//es porque este es autogenerado no lo validas es diferente
	_, err := m.DB.ExecContext(ctx, query, args...)

	return err
}

func (m TokenModel) DeleteAllForUser(scope string, userID int) error {

	query := `
	delete from tokens 
	where scope = $1 and user_id = $2 
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, scope, userID)

	return err

}
