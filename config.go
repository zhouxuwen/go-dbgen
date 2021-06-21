package main

import "os"

const (
	ProjectName = "go-dbgen"
	Version     = "v1.0.0"
	Copyright   = "2021.06"
	Author      = ""
	AuthorEmail = ""
)

const (
	DS              = string(os.PathSeparator) // 通用/
	DbNullPrefix    = "Null"                   // 处理数据为空时结构的前缀定义
	TablePrefix     = "TABLE_"                 // 表前缀
	DefaultSavePath = "output"                 // 默认生成目录名称
)

const (
	TPL_MARKDOWN = "assets/tpl/markdown.tpl" // markdown模板
	TPL_MODELS   = "assets/tpl/models.tpl"   // models模板
	TPL_GORM     = "assets/tpl/gorm.tpl"     // gorm模板
)

const (
	Unknown = iota
	Darwin
	Window
	Linux
)

// generate file name
const (
	GODIR_MODELS = "models" // model file
	GODIR_Config = "config" // config file
	GODIR_Entity = "entity" // entity file

	GOFILE_ENTITY    = "db_entity.go"    // entity table file
	GoFile_TableList = "table_list.go"   // table file
	GoFile_Init      = "init.go"         // init file
	GoFile_Error     = "e.go"            // error file
	GoFile_Example   = "example_test.go" // example file
)

const (
	PkgDbModels = "mysql"  // db_models package name
	PkgEntity   = "entity" // entity package name
	PkgModels   = "models" // models package name
	PkgTable    = "config" // table package name
)

// help list
var CmdHelp = []CmdEntity{
	{"0", "Set build directory"},
	{"1", "Generate the table markdown document"},
	{"2", "Generate table structure entities"},
	{"3", "Generate CURD insert, delete, update and select"},
	{"4", "Sets the struct mapping name"},
	{"5", "Find or set the table name"},
	{"7, c, clear", "Clear the screen"},
	{"8, h, help", "Show help list"},
	{"9, q, quit", "Quit"},
}

//MysqlTypeToGoType mysql类型 <=> golang类型
var MysqlTypeToGoType = map[string]string{
	"tinyint":    "int32",
	"smallint":   "int32",
	"mediumint":  "int32",
	"int":        "int32",
	"integer":    "int64",
	"bigint":     "int64",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}

// PostgresTypeToGoType postgres数据类型转成go数据类型
var PostgresTypeToGoType = map[string]string{
	"int2":        "int",
	"int4":        "int",
	"int8":        "int64",
	"_int2":       "pq.Int64Array",
	"_int4":       "pq.Int64Array",
	"_int8":       "pq.Int64Array",
	"float4":      "float32",
	"float8":      "float64",
	"money":       "float64",
	"_float4":     "pq.Float64Array",
	"_float8":     "pq.Float64Array",
	"numeric":     "float64",
	"date":        "string",
	"text":        "string",
	"varchar":     "string",
	"char":        "string",
	"_varchar":    "pq.StringArray",
	"_char":       "pq.StringArray",
	"timestamptz": "time.Time",
	"time":        "time.Time",
	"bool":        "bool",
}
