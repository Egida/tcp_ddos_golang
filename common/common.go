package common

import (
	"log"
	"os"
	"strconv"

	"github.com/catalinc/hashcash"
	"github.com/pkg/errors"
)

func GetenvIntDefault(env string, defaultVal int) (int, error) {
	e := os.Getenv(env)
	if len(e) == 0 {
		return defaultVal, nil
	}
	i, err := strconv.Atoi(e)
	if err != nil {
		return defaultVal, errors.WithStack(err)
	}
	return i, nil
}

func HashcashObjFromEnv() (*hashcash.Hash, error) {
	bits, err := GetenvIntDefault("HASHCASH_BITS", 20)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	saltLength, err := GetenvIntDefault("HASHCASH_SALT_LENGTH", 8)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	extra := os.Getenv("HASH_CASH_EXTRA")
	log.Printf("bits: %+v, saltLength: %+v, extra: %+v", bits, saltLength, extra)
	h := hashcash.New(uint(bits), uint(saltLength), extra)
	return h, nil
}
