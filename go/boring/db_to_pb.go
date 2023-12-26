package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
)

type Table struct {
	Table string `gorm:"column:TABLE_NAME"`
}

func (*Table) TableName() string {
	return "TABLES"
}

type Column struct {
	ColumnName string `gorm:"column:COLUMN_NAME"`
	DataType   string `gorm:"column:DATA_TYPE"`
	Comment    string `gorm:"column:COLUMN_COMMENT"`
}

func (*Column) TableName() string {
	return "COLUMNS"
}

func main() {
	dbName := "dopai"
	dns := "lzj:123456@tcp(192.168.31.54:3306)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	tables := make([]*Table, 0)
	if err = db.Where("table_schema = ?", dbName).Find(&tables).Error; err != nil {
		panic(err)
	}
	message := "syntax=\"proto3\";\npackage pbdef;\n"
	for _, table := range tables {
		message += generateMessage(db, dbName, table.Table) + "\n"
	}
	err = os.WriteFile("out.proto", []byte(message), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to write protobuf to file:", err)
		return
	}
	fmt.Println("Protobuf structure has been saved to out.proto.")
}

func generateMessage(db *gorm.DB, dbName, table string) string {
	columns := make([]*Column, 0)
	err := db.Where("table_schema = ? and table_name = ?", dbName, table).Find(&columns).Error
	if err != nil {
		return ""
	}

	messageStructure := ""
	for index, col := range columns {
		if col.Comment != "" {
			messageStructure += fmt.Sprintf("    %s %s = %d; //%s\n", getProtobufType(col.DataType), camelCase(col.ColumnName), index+1, col.Comment)
		} else {
			messageStructure += fmt.Sprintf("    %s %s = %d;\n", getProtobufType(col.DataType), camelCase(col.ColumnName), index+1)
		}
	}
	protoMessage := fmt.Sprintf("message %s {\n%s}\n", tableNameToCamelCase(table), messageStructure)
	return protoMessage
}

// 将表名转为驼峰命名
func tableNameToCamelCase(tableName string) string {
	words := strings.Split(tableName, "_")
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToTitle(words[i])
	}
	return strings.Join(words, "")
}

// 将字段名转为驼峰命名
func camelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if i == 0 {
			parts[i] = strings.ToLower(part)
		} else {
			parts[i] = strings.ToTitle(part)
		}
	}
	return strings.Join(parts, "")
}

func getProtobufType(mysqlType string) string {
	switch mysqlType {
	case "varchar", "char", "text", "mediumtext", "longtext":
		return "string"
	case "timestamp", "datetime", "date", "time":
		return "google.protobuf.Timestamp"
	case "bigint":
		return "int64"
	case "int", "mediumint", "smallint", "tinyint":
		return "int32"
	case "double", "decimal":
		return "double"
	case "float":
		return "float"
	case "json":
		return "google.protobuf.Any"
	case "enum", "set":
		return "string"
	case "binary", "varbinary", "blob", "longblob", "mediumblob", "tinyblob":
		return "bytes"
	default:
		return "string"
	}
}
