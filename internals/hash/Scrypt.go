package hash

import (
	internalContracts "github.com/mohammadMghi/moduleForGo/internals/contracts"
	"golang.org/x/crypto/scrypt"
)

func Generate(config internalContracts.IConfiguration, val string) ([]byte, error) {
	configSalt, err := config.Get("Salt")
	if err != nil {
		return nil,err
	}
	salt := []byte(configSalt.(string))

	dk, err := scrypt.Key([]byte(val), salt, 1<<15, 8, 1, 32)
	if err != nil {
		return nil,err
	}
	return dk, nil
}
