package keeper

import (
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"
)


type Row struct {
    TableName string       `json:"table_name"`
    Id uint                `json:"id"`
    Fields types.RowFields `json:"columns"`
}

func NewRow(tableName string, id uint, fields types.RowFields) Row {
    return Row {
    TableName: tableName,
    Id: id,
    Fields: fields,
    }
}

