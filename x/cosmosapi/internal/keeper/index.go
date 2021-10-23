package keeper

import (
    "fmt"
    "errors"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"
)

func (k Keeper) updateIndex(ctx sdk.Context, tableName string, id uint, fields types.RowFields) (uint, error){
    store := ctx.KVStore(k.storeKey)

    oldRecord, _ := k.Find(ctx, tableName, id)

    indexFields, err := k.GetIndex(ctx, tableName)
    if err != nil {
        return 0, errors.New(fmt.Sprintf("Failed to get index for table %s", tableName))
    }

    if id == 0 {
        return 0, errors.New(fmt.Sprintf("Id for table %s is invalid", tableName))
    }

    var mold []string

    for _, indexField := range indexFields {
        if value, ok := fields[indexField]; ok {
            if oldRecord != nil {
                if oldValue, oldOk := oldRecord[indexField]; oldOk {
                    oldKey := getIndexKey(tableName, indexField, oldValue)
                    bz := store.Get([]byte(oldKey))
                    if bz != nil {
                        k.cdc.MustUnmarshalBinaryBare(bz, &mold)
                        mold = removeItemFromSet(mold, fmt.Sprint(id))
                        if len(mold) > 0 {
                            store.Set([]byte(oldKey), k.cdc.MustMarshalBinaryBare(mold))
                        } else {
                            store.Delete([]byte(oldKey))
                        }
                    }
                }
            }
            key := getIndexKey(tableName, indexField, value)
            store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(id)) 
        }
    }

    return id, nil
}

//////////////////////
//                  //
// helper functions //
//                  //
//////////////////////

func removeItemFromSet(set []string, item string) []string {
    for i, v := range set {
        if v == item {
            set[i] = set[len(set)-1]
            set[len(set)-1] = ""    // probably keep from mem leaking
            return set[:len(set)-1]
        }
    }
    return set
}
