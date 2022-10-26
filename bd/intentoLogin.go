package bd

import (
	"github.com/coralheyman/twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	passBytes := []byte(password)
	passBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passBD, passBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
