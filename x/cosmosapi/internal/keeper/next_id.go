package keeper

import (
    "fmt"
    "sync"
    "errors"
    sdk "github.com/cosmos/cosmos-sdk/types"

)

var mutex = &sync.Mutex{}
var NextIds = make(map[string]uint)

func getNextId(k Keeper, ctx sdk.Context, tableName string) (uint, error) {
    store := ctx.KVStore(k.storeKey)
    mutex.Lock()
    defer mutex.Unlock()

    var nextIdKey = getNextIdKey(tableName)
    var nextId uint
    var found bool
    if nextId, found = NextIds[tableName]; found {
    } else if bz := store.Get([]byte(nextIdKey)); bz != nil {
        k.cdc.MustUnmarshalBinaryBare(bz, &nextId)
    } else if bz = store.Get([]byte(getTableKey(tableName))); bz != nil {
        nextId = 1
    } else {
        return 0, errors.New(fmt.Sprintf("Invalid table name %s", tableName))
    }

    store.Set([]byte(nextIdKey), k.cdc.MustMarshalBinaryBare(nextId + 1))
    NextIds[tableName] = nextId + 1

    return nextId, nil
}

