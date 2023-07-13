package types

const (
	// ModuleName defines the module name
	ModuleName = "od"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_od"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ProposalKey      = "Proposal/value/"
	ProposalCountKey = "Proposal/count/"
)

const (
	ApplicantStake int64 = 1
	ExpiryBlocks   int64 = 60
)

var MPs = []string{
	"opendao1t79pyys2afsr9z7e0td6sc30txz7psacjxfgyy",
	"opendao1cvmajkjw4x5rjpag28eex27c9m396vheekt9h0",
	"opendao1e6g36xjk6urjk79t35q64tuxehsnhuvmht5fdl",
}
