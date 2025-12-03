package chainstore

const (
	// Prefixes
	AccountPrefix = "account/"
)

func accountKey(address string) []byte {
	return append([]byte(AccountPrefix), []byte(address)...)
}
