package token

import (
	"fmt"
	"log"

	"github.com/andrew-lawlor/cc-simple-auth-server/db"
)

var tokenCache = make(map[string]bool)

// LoadTokens loads tokens from the database into the cache.
func LoadTokens() error {
	db := db.GetDB()
	rows, err := db.Query("SELECT token FROM tokens")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var token string
		if err := rows.Scan(&token); err != nil {
			fmt.Println(err.Error())
			return err
		}
		tokenCache[token] = true
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	log.Println("Tokens successfully loaded into cache.")
	return nil
}

func IsTokenValid(token string) bool {
	return tokenCache[token]
}
