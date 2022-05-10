package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mohammadreza-torkaman/blog/x/blog/types"
)

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//Get comment to delete
	commentToDelete, found := k.GetComment(ctx, msg.CommentID)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrID, "Comment doesnt exist")
	}
	//If this comment doesn't belong to this post
	if commentToDelete.PostID != msg.PostID {
		return nil, sdkerrors.Wrapf(types.ErrID,
			"Post Blog Id does not exist for which comment with Blog Id %d was made", msg.PostID)
	}
	k.RemoveComment(ctx, commentToDelete.Id)

	return &types.MsgDeleteCommentResponse{}, nil
}
