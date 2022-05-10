package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mohammadreza-torkaman/blog/x/blog/types"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrID,
			"Post Blog Id does not exist for which comment with Blog Id %d was made", msg.PostID)
	}

	comment := types.Comment{
		Creator:   msg.Creator,
		Title:     msg.Title,
		Body:      msg.Body,
		PostID:    msg.PostID,
		CreatedAt: ctx.BlockHeight(),
	}

	// Check if the comment is older than the Post. If more than 100 blocks, then return error.
	if comment.CreatedAt > post.CreatedAt+100 {
		return nil, sdkerrors.Wrapf(types.ErrCommentOld,
			"Comment created at %d is older than post created at %d", comment.CreatedAt, post.CreatedAt)
	}

	commentID := k.AppendComment(ctx, comment)

	return &types.MsgCreateCommentResponse{Id: commentID}, nil
}
