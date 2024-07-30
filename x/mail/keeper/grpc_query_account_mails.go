package keeper

import (
	"context"

	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountMails(goCtx context.Context, req *types.QueryAccountMailsRequest) (*types.QueryAccountMailsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	addressBz, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	var mails []types.Mail
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	addressMailIndexStore := prefix.NewStore(prefix.NewStore(store, types.KeyPrefix(types.AddressMailIndexKey)), addressBz)

	pageRes, err := query.Paginate(addressMailIndexStore, req.Pagination, func(key []byte, value []byte) error {
		mail, _ := k.GetMail(ctx, sdk.BigEndianToUint64(key[len(key)-8:]))

		mails = append(mails, mail)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAccountMailsResponse{Mail: mails, Pagination: pageRes}, nil
}
