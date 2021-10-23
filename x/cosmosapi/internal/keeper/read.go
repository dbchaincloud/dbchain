package keeper

import (
    "fmt"
    "errors"
    "strconv"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"
)

func (k Keeper) Find(ctx sdk.Context, tableName string, id uint) (types.RowFields, error){
    store := ctx.KVStore(k.storeKey)

    fieldNames, err := k.getTableFields(ctx, tableName)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Failed to get fields for table %s", tableName))
    }

    if id == 0 {
        return nil, errors.New("Id cannot be 0")
    }

    var fields = make(types.RowFields)
    var value string

    for _, fieldName := range fieldNames {
        key := getDataKey(tableName, id, fieldName)
        bz := store.Get([]byte(key)) 
        if bz != nil {
            k.cdc.MustUnmarshalBinaryBare(bz, &value)
            fields[fieldName] = value
        }
    }

    return fields, nil
}

// Find by an attribute in the r.Fields
func (k Keeper) FindBy(ctx sdk.Context, tableName string, field string,  value string) []uint {
    store := ctx.KVStore(k.storeKey)

    var hasIndex bool
    indexFields, err := k.GetIndex(ctx, tableName)
    if err == nil {
        for _, item := range(indexFields) {
            if item == field {
                hasIndex = true
                break
            }
        }
    }

    if hasIndex {
        var mold []uint
        key := getIndexKey(tableName, field, value)
        bz := store.Get([]byte(key))
        if bz == nil {
            return []uint{}
        } else {
            k.cdc.MustUnmarshalBinaryBare(bz, &mold)
            return mold
        }
    }

    // so-called full table scanning
    var result []uint
    start, end := getDataIteratorStartAndEndKey(tableName)
    iter := store.Iterator([]byte(start), []byte(end))
    var mold string
    for ; iter.Valid(); iter.Next() {
        key := iter.Key()
        k.cdc.MustUnmarshalBinaryBare(key, &mold)
        keyString := fmt.Sprint(mold)
        fn := getFieldNameFromDataKey(keyString)
        if fn == field {
	    val := iter.Value()
            k.cdc.MustUnmarshalBinaryBare(val, &mold)
            if mold == value {
                id := getIdFromDataKey(keyString)
                u64, _ := strconv.ParseUint(id, 10, 64)
                result = append(result, uint(u64))
            }
        }
    }
    return result
}

