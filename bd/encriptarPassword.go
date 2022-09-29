package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword Función para encriptar una password con un costo de 8, es 2 elevado a la 8 */
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
