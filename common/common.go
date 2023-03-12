package common

import (
	"bytes"
	"github.com/catalinc/hashcash"
	"github.com/pkg/errors"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// GetenvIntDefault parses environment variable to integer.
// The default value is returned in case evn var is empty.
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

// HashcashObjFromEnv returns Hash object for generating and validating hashcash stamp,
// hashcash parameters are read from environment variables.
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

// ReqWisdom connects to tcpServer and sends hashcash stamp for POW verification.
// In case of success verification it gets a random joke from the server.
func ReqWisdom(tcpServAddr, hashcashStamp string) (string, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpServAddr)
	if err != nil {
		return "", errors.WithStack(err)
	}
	conn, errD := net.DialTCP("tcp", nil, tcpAddr)
	if errD != nil {
		return "", errors.WithStack(errD)
	}
	defer conn.Close()
	strEcho := hashcashStamp
	if _, errW := conn.Write([]byte(strEcho)); errW != nil {
		return "", errors.WithStack(errW)
	}
	//log.Println("write to server = ", strEcho)
	reply := make([]byte, 1024)
	_, errD = conn.Read(reply)
	if errD != nil {
		return "", errors.WithStack(errD)
	}
	return strings.TrimSpace(string(bytes.Trim(reply, "\x00"))), nil
}