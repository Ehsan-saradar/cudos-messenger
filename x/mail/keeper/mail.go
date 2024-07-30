package keeper

import (
	"encoding/binary"

	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetMailCount get the total number of mail
func (k Keeper) GetMailCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MailCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMailCount set the total number of mail
func (k Keeper) SetMailCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MailCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMail appends a mail in the store with a new id and update the count
func (k Keeper) AppendMail(
	ctx sdk.Context,
	mail types.Mail,
) uint64 {
	// Create the mail
	count := k.GetMailCount(ctx)

	// Set the ID of the appended value
	mail.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MailKey))
	appendedValue := k.cdc.MustMarshal(&mail)
	store.Set(GetMailIDBytes(mail.Id), appendedValue)

	// Update mail count
	k.SetMailCount(ctx, count+1)

	k.appendMailIndex(ctx, mail.From, mail.Id)
	k.appendMailIndex(ctx, mail.To, mail.Id)

	return count
}

func (k Keeper) appendMailIndex(
	ctx sdk.Context,
	address string,
	id uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressMailIndexKey))
	store.Set(types.AddressMailIndex(address, id), []byte{})
}

// SetMail set a specific mail in the store
func (k Keeper) SetMail(ctx sdk.Context, mail types.Mail) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MailKey))
	b := k.cdc.MustMarshal(&mail)
	store.Set(GetMailIDBytes(mail.Id), b)
}

// GetMail returns a mail from its id
func (k Keeper) GetMail(ctx sdk.Context, id uint64) (val types.Mail, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MailKey))
	b := store.Get(GetMailIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllMail returns all mail
func (k Keeper) GetAllMail(ctx sdk.Context) (list []types.Mail) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MailKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mail
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMailIDBytes returns the byte representation of the ID
func GetMailIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMailIDFromBytes returns ID in uint64 format from a byte array
func GetMailIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
