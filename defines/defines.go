package defines

var StructTemplateText = map[string]string{
	"title": "pub struct STRUCT_NAME {",
	"line": "	pub	FIELD_NAME	:	FIELD_TYPE,",
	"end": "}",
}

// DBTypeToStructType 数据库数据类型到 go 结构体数据类型的转换规则
var DBTypeToStructType = map[string]string{
	"int":                "i32",
	"integer":            "i32",
	"tinyint":            "i8",
	"smallint":           "i16",
	"mediumint":          "i32",
	"bigint":             "i64",
	"int unsigned":       "u32",
	"integer unsigned":   "u32",
	"tinyint unsigned":   "u8",
	"smallint unsigned":  "u16",
	"mediumint unsigned": "u32",
	"bigint unsigned":    "u64",
	"bit":                "byte",
	"bool":               "bool",
	"enum":               "String",
	"set":                "String",
	"varchar":            "String",
	"char":               "String",
	"tinytext":           "String",
	"mediumtext":         "String",
	"text":               "String",
	"longtext":           "String",
	"blob":               "String",
	"tinyblob":           "String",
	"mediumblob":         "String",
	"longblob":           "String",
	"date":               "std::time::SystemTime",
	"datetime":           "std::time::SystemTime",
	"timestamp":          "std::time::SystemTime",
	"time":               "std::time::SystemTime",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "String",
	"varbinary":          "String",
}

// TableColumn 数据库中字段信息
type TableColumn struct {
	ColumnName    string
	DataType      string
	ColumnKey     string
	IsNullable    string
	ColumnType    string
	ColumnComment string
}

// StructColumn go 结构体字段信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}
