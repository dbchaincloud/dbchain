package keeper

import (
    "fmt"
    "strings"
)

const (
    KeyPrefixMeta  = "mt"
    KeyPrefixData  = "dt"
    KeyPrefixIndex = "ix"
    KeyPrefixGroup = "grp"
)

// to store name of all tables
func getTablesKey() string {
    return fmt.Sprintf("%s:tables", KeyPrefixMeta)
}

// to store the id for next new record of a table
func getNextIdKey(tableName string) string {
    return fmt.Sprintf("%s:nextId:%s", KeyPrefixMeta, tableName)
}

// to store the name of fields for a table
func getTableKey(tableName string) string {
    return fmt.Sprintf("%s:tn:%s", KeyPrefixMeta, tableName)
}

// to store table fields which have index on
func getMetaTableIndexKey(tableName string) string {
    return fmt.Sprintf("%s:idx:%s", KeyPrefixMeta, tableName)
}

// to store the id of a indexed field
func getIndexKey(tableName string, fieldName string, value string) string {
    return fmt.Sprintf("%s:%s:%s:%s", KeyPrefixIndex, tableName, fieldName, value)
}

//////////////////
//              //
// data related //
//              //
//////////////////

// to store the value of a fields on a record of a table.
func getDataKey(tableName string, id uint, fieldName string) string {
    return fmt.Sprintf("%s:%s:%d:%s", KeyPrefixData, tableName, id, fieldName)
}

// to get the start and end parameters of iterator which seeks certain value of a field
func getDataIteratorStartAndEndKey(tableName string) (string, string) {
    start := fmt.Sprintf("%s:%s:", KeyPrefixData, tableName)
    end   := fmt.Sprintf("%s:%s;", KeyPrefixData, tableName)
    return start, end
}

func getIdFromDataKey(key string) string {
    arr := strings.Split(key, ":")
    return arr[2]
}

func getFieldNameFromDataKey(key string) string {
    arr := strings.Split(key, ":")
    return arr[3]
}

///////////////////
//               //
// group related //
//               //
///////////////////

func getGroupKey(groupName string) string {
    return fmt.Sprintf("%s:%s", KeyPrefixGroup, groupName)
}

func getAdminGroupKey() string {
    return getGroupKey("admin")
}
