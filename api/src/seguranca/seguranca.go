package seguranca

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func CompararHashComSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
