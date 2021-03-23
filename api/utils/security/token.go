package security

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"github.com/jsdaniell/whats_api/api/models"
	"time"
)

const layout = "2006-01-02T15:04:05.000Z"

// CreateToken Creates the token stored on each user, receives a generic string to generate token based on string.
func CreateToken(s string) models.Token {
	h := sha1.New()
	h.Write([]byte(s))

	hashes := hex.EncodeToString(h.Sum(nil))
	actual := time.Now()
	date := actual.Format(time.RFC3339)

	var t = models.Token{
		CreatedAt: date,
		Token:     hashes,
	}

	return t
}

func ValidateToken(tk models.Token) error {

	t, err := time.Parse(time.RFC3339, tk.CreatedAt)
	if err != nil {
		return err
	}

	dayLater := t.Add(24 * time.Hour)

	if !t.Before(dayLater) {
		return errors.New("token expired")
	}

	return nil
}
