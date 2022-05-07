package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/mohammadreza-torkaman/blog/x/blog/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Posts(goCtx context.Context, req *types.QueryPostsRequest) (*types.QueryPostsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var posts []*types.Post
	store := ctx.KVStore(k.storeKey)

	postStore := prefix.NewStore(store, []byte(types.PostKey))

	paginate, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var post types.Post
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}
		posts = append(posts, &post)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPostsResponse{Post: posts, Pagination: paginate}, nil
}
