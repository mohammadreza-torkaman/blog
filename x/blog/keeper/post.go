package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mohammadreza-torkaman/blog/x/blog/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)
	post.Id = count
	//setting position to the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)
	value := k.cdc.MustMarshal(&post)
	store.Set(byteKey, value)
	k.SetPostCount(ctx, count+1)
	//post id
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetPost(ctx sdk.Context, postID uint64) (types.Post, bool) {
	//getting store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))
	//Making post ID as a key
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, postID)
	// Get the post bytes using post ID as a key
	fetchedPost := store.Get(key)
	var post types.Post
	if fetchedPost == nil {
		return post, false
	}
	//Unmarshalling post bytes into object
	k.cdc.MustUnmarshal(fetchedPost, &post)
	return post, true
}
