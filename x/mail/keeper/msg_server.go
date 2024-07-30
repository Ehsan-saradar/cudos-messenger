package keeper

import (
	"context"

	"github.com/CudoVentures/cudos-node/x/mail/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Send defines a method that creates a new mail
func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create the mail
	var mail = types.Mail{
		From:            msg.Sender,
		To:              msg.To,
		Subject:         msg.Subject,
		Body:            msg.Body,
		CreatedAtHeight: ctx.BlockHeight(),
	}

	k.AppendMail(ctx, mail)

	return &types.MsgSendResponse{}, nil
}
