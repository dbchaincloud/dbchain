package keeper

import (
    "fmt"
    "errors"
    "strings"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"
)

/////////////////////////////
//                         //
// table related functions //
//                         //
/////////////////////////////

func (k Keeper) getTables(ctx sdk.Context) ([]string, error) {
    store := ctx.KVStore(k.storeKey)
    tablesKey := getTablesKey()
    bz := store.Get([]byte(tablesKey))
    if bz == nil {
        return nil, errors.New("No tables found")
    }
    var tableNames []string
    k.cdc.MustUnmarshalBinaryBare(bz, &tableNames)
    return tableNames, nil
}


// Check if the table is present in the store or not
func (k Keeper) IsTablePresent(ctx sdk.Context, name string) bool {
    store := ctx.KVStore(k.storeKey)
    return store.Has([]byte(getTableKey(name)))
}

func (k Keeper) IsFieldPresent(ctx sdk.Context, tableName string, field string) bool {
    table, err := k.GetTable(ctx, tableName)
    if err != nil {
        return false
    }
    for _, f := range(table.Fields) {
        if f == field {
            return true
        }
    }
    return false
}

// Create a new table
func (k Keeper) CreateTable(ctx sdk.Context, owner sdk.AccAddress, name string, fields []string) {
    store := ctx.KVStore(k.storeKey)
    table := types.NewTable()
    table.Owner = owner
    table.Name = name
    table.Fields = preProcessFields(fields)
    store.Set([]byte(getTableKey(table.Name)), k.cdc.MustMarshalBinaryBare(table))

    var tables []string
    bz :=store.Get([]byte(getTablesKey()))
    if bz == nil {
        tables = append(tables, table.Name)
    } else {
        k.cdc.MustUnmarshalBinaryBare(bz, &tables)
        tables = append(tables, table.Name)
    }
    store.Set([]byte(getTablesKey()), k.cdc.MustMarshalBinaryBare(tables))
}


// Get a table 
func (k Keeper) GetTable(ctx sdk.Context, name string) (types.Table, error) {
    store := ctx.KVStore(k.storeKey)
    bz := store.Get([]byte(getTableKey(name)))
    if bz == nil {
        return types.Table{}, errors.New(fmt.Sprintf("table %s not found", name))
    }
    var table types.Table
    k.cdc.MustUnmarshalBinaryBare(bz, &table)
    return table, nil
}

// Add a field
func (k Keeper) AddField(ctx sdk.Context, name string, field string) (bool, error) {
    table, err := k.GetTable(ctx, name)
    if err != nil {
        return false, err
    }

    for _, fld := range table.Fields {
        if field == fld {
            return false, errors.New(fmt.Sprintf("field %s existed already", field))
        }
    }

    table.Fields = append(table.Fields, field)

    store := ctx.KVStore(k.storeKey)
    store.Set([]byte(getTableKey(table.Name)), k.cdc.MustMarshalBinaryBare(table))
    return true, nil
}

// Remove a field
func (k Keeper) RemoveField(ctx sdk.Context, name string, field string) (bool, error){
    table, err := k.GetTable(ctx, name)
    if err != nil {
        return false, err
    }

    if field == "id" {
        return false, errors.New(fmt.Sprintf("cannot remove field id"))
    }

    var foundField = false 
    for i, fld := range table.Fields {
        if field == fld {
            foundField = true
            table.Fields = append(table.Fields[:i], table.Fields[i+1:]...)
            break
        }
    }
    if !foundField {
        return false, errors.New(fmt.Sprintf("field %s not existed", field))
    }

    store := ctx.KVStore(k.storeKey)
    store.Set([]byte(getTableKey(table.Name)), k.cdc.MustMarshalBinaryBare(table))
    return true, nil
}

// Rename a field
func (k Keeper) RenameField(ctx sdk.Context, name string, oldField string, newField string) (bool, error) {
    table, err := k.GetTable(ctx, name)
    if err != nil {
        return false, err
    }

    if oldField == "" || newField == "" {
        return false, errors.New(fmt.Sprintf("cannot have empty field name"))
    }

    oldField = strings.ToLower(oldField)
    newField = strings.ToLower(newField)

    if oldField == "id" || newField == "id" {
        return false, errors.New(fmt.Sprintf("cannot rename field id"))
    }


    var foundField = false
    var index = 0
    for i, fld := range table.Fields {
        if oldField == fld {
            foundField = true
            index = i
            break
        }
        if newField == fld {
            return false, errors.New(fmt.Sprintf("cannot rename to field %s", newField))
        }
    }
    if !foundField {
        return false, errors.New(fmt.Sprintf("field %s not found", oldField))
    }

    table.Fields[index] = newField

    store := ctx.KVStore(k.storeKey)
    store.Set([]byte(getTableKey(table.Name)), k.cdc.MustMarshalBinaryBare(table))
    return true, nil
}

/////////////////////////////
//                         //
// index related functions //
//                         //
/////////////////////////////

// for now we only support indexes on single field
func (k Keeper) CreateIndex(ctx sdk.Context, owner sdk.AccAddress, tableName string, field string) {
    store := ctx.KVStore(k.storeKey)
    key := getMetaTableIndexKey(tableName)
    var index_fields []string

    bz := store.Get([]byte(key))
    if bz != nil {
        k.cdc.MustUnmarshalBinaryBare(bz, &index_fields)
    }
    index_fields = append(index_fields, field)
    store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(index_fields))
    // TODO: to create index data for the existing records of the table
}

func (k Keeper) GetIndex(ctx sdk.Context, tableName string) ([]string, error) {
    store := ctx.KVStore(k.storeKey)
    key := getMetaTableIndexKey(tableName)
    bz := store.Get([]byte(key))
    if bz == nil {
        return []string{}, errors.New(fmt.Sprintf("Index of table %s not found", tableName))
    }
    var index_fields []string
    k.cdc.MustUnmarshalBinaryBare(bz, &index_fields)
    return index_fields, nil
}

////////////////////
//                //
// helper methods //
//                //
////////////////////

// to preprocess the new table field names
// to make sure the fields be lowercase
// to make sure field id be in place
func preProcessFields(fields []string) []string {
    var result = []string{"id"}
    m := make(map[string]bool)
    m["id"] = true

    for _, field := range fields {
        newName := strings.ToLower(field)
        if newName == "" {
            continue
        }
        if m[newName] {
            continue
        } else {
           m[newName] = true
           result = append(result, newName)
        }
    }
    return result
}
