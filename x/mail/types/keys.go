package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "mail"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mail"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func AddressMailIndex(address string, id uint64) []byte {
	bz, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, len(bz)+8)
	copy(buf, bz)
	copy(buf[len(address):], sdk.Uint64ToBigEndian(id))
	return buf
}

const (
	MailKey             = "Mail-value-"
	MailCountKey        = "Mail-count-"
	AddressMailIndexKey = "AddressMailIndex-value-"
)
