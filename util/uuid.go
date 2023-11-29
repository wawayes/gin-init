package util

import (
	"crypto/md5"
	"fmt"
	"strconv"

	uuid "github.com/nu7hatch/gouuid"
)

const (
	DefaultPrefixShortidString = "default-awsdaexy0yi1s02m"
	DefaultShortidString       = "awsdaexy0yi1s02m"
	ShortidDigits              = "abcdefghijkmnpqrstuvwxyz0123456789"
)

// NewShortIDString create a shortid string, return <prefix>-<shortid> if prefix is not empty, or just return shortid
func NewShortIDString(prefix string) string {
	needPrefix := prefix != ""
	shortidStr := DefaultShortidString
	if needPrefix {
		shortidStr = DefaultPrefixShortidString
	}

	newID, _ := uuid.NewV4()

	shortidStr = UUIDToShortID(newID.String())
	if needPrefix {
		shortidStr = prefix + "-" + shortidStr
	}
	return shortidStr
}

func UUIDToShortID(UUID string) string {
	// 32uuid -> 32md5 hex
	data := []byte(UUID)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)

	var result []byte
	for i := 0; i < 16; i++ {
		// parse 2bit char from 16base to 10base
		index, _ := strconv.ParseUint(md5str[2*i:2*i+2], 16, 32)
		result = append(result, ShortidDigits[index%34])
	}
	return string(result)
}
