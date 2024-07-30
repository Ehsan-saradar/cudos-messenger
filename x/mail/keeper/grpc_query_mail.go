package keeper

import (
	"context"

	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MailAll(c context.Context, req *types.QueryAllMailRequest) (*types.QueryAllMailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mails []types.Mail
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	mailStore := prefix.NewStore(store, types.KeyPrefix(types.MailKey))

	pageRes, err := query.Paginate(mailStore, req.Pagination, func(key []byte, value []byte) error {
		var mail types.Mail
		if err := k.cdc.Unmarshal(value, &mail); err != nil {
			return err
		}

		mails = append(mails, mail)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMailResponse{Mail: mails, Pagination: pageRes}, nil
}

func (k Keeper) Mail(c context.Context, req *types.QueryGetMailRequest) (*types.QueryGetMailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	mail, found := k.GetMail(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMailResponse{Mail: mail}, nil
}
