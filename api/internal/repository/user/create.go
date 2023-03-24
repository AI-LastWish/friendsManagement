package user

import (
	"context"
	"errors"
	"fmt"
)

// CreateRelationship: create a friend connection between two email addresses.
func (repo *UserRepository) CreateRelationship(email string, friend string, stmt string) error {
	if email == friend {
		return errors.New("2 input emails are the same")
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := repo.Get(email)
	if err != nil {
		return err
	}

	_, errFriend := repo.Get(friend)
	if errFriend != nil {
		return errFriend
	}

	_, err = repo.db.ExecContext(ctx, stmt,
		email,
		friend,
	)
	if err != nil {
		fmt.Println("er", err)
		return err
	}

	return nil
}
