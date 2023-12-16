package datastore

import (
	"errors"

	"github.com/DevanshBatra20-PasswordManager/helpers"
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/gofr"
)

type password struct{}

func NewPassword() *password {
	return &password{}
}

func (p *password) Create(ctx *gofr.Context, password *models.Password, userId string) (string, error) {
	var firstName string
	_ = ctx.DB().QueryRowContext(ctx, "SELECT (First_Name) FROM users WHERE id = (?)", userId).Scan(&firstName)
	if firstName == "" {
		return "", errors.New("User with userId " + userId + " does not exist")
	}

	encryptedPassword, err := helpers.EncryptAES(password.Password_Value)
	if err != nil {
		return "Error in encrypting password", nil
	}
	_, err = ctx.DB().ExecContext(ctx, "INSERT INTO passwords (password_name, password_value, password_type, user_id) VALUES (?, ?, ?, ?)",
		password.Password_Name, encryptedPassword, password.Password_Type, userId)
	if err != nil {
		return "Error in creating password", err
	}

	return "Password created sucessfully", nil
}

func (p *password) Delete(ctx *gofr.Context, passwordId string) (string, error) {
	var passwordName string
	err := ctx.DB().QueryRowContext(ctx, "SELECT (password_name) FROM passwords WHERE password_id = (?)", passwordId).Scan(&passwordName)
	if err != nil {
		return "", errors.New("Password with passwordId " + passwordId + " does not exist")
	}

	_, err = ctx.DB().ExecContext(ctx, "DELETE FROM passwords WHERE password_id = (?)", passwordId)
	if err != nil {
		return "", err
	}

	return "Password Deleted Successfully", nil
}

func (p *password) Update(ctx *gofr.Context, password *models.Password, passwordId string) (string, error) {
	var passwordName string
	err := ctx.DB().QueryRowContext(ctx, "SELECT (password_name) FROM passwords WHERE password_id = (?)", passwordId).Scan(&passwordName)
	if err != nil {
		return "", errors.New("Password with passwordId " + passwordId + " does not exist")
	}

	encryptedPassword, err := helpers.EncryptAES(password.Password_Value)
	if err != nil {
		return "", nil
	}
	_, err = ctx.DB().ExecContext(ctx, "UPDATE passwords SET password_name = (?), password_value = (?), password_type = (?) WHERE password_id = (?)",
		password.Password_Name, encryptedPassword, password.Password_Type, passwordId)
	if err != nil {
		return "", err
	}

	return "Password Updated Successfully", nil
}

func (p *password) Get(ctx *gofr.Context, passwordId string) (*models.Password, error) {
	var password models.Password
	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM passwords WHERE password_id = (?)", passwordId).Scan(&password.Password_Id, &password.Password_Name, &password.Password_Value, &password.Password_Type, &password.User_Id)
	if err != nil {
		return &models.Password{}, errors.New("Password with passwordId " + passwordId + " does not exist")
	}

	password.Password_Value, err = helpers.DecryptAES(password.Password_Value)
	if err != nil {
		return &models.Password{}, nil
	}

	return &password, nil
}

func (p *password) GetByUserId(ctx *gofr.Context, userId string) ([]*models.Password, error) {
	var passwords []*models.Password
	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM passwords WHERE user_id = (?)", userId)
	if err != nil {
		return []*models.Password{}, errors.New("User with userId " + userId + " does not exist")
	}
	defer rows.Close()

	for rows.Next() {
		var password models.Password
		err := rows.Scan(&password.Password_Id, &password.Password_Name, &password.Password_Value, &password.Password_Type, &password.User_Id)
		if err != nil {
			return []*models.Password{}, err
		}

		password.Password_Value, err = helpers.DecryptAES(password.Password_Value)
		if err != nil {
			return []*models.Password{}, nil
		}

		passwords = append(passwords, &password)
	}

	return passwords, nil
}
