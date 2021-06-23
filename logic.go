package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"
	"sync"
)

type Logic struct {
	T    *Tools
	DB   *ModelS
	Path string
	Once sync.Once
	l    sync.Mutex
}

// CreateModelsEntity 生成models层结构体
func (l *Logic) CreateModelsEntity(formatList []string) (err error) {
	// 将表结构写入文件
	for idx, table := range l.DB.DoTables {
		idx++
		// 表结构文件路径
		path := CreateDir(l.Path+GODIR_MODELS) + DS + table.Name + ".go"
		fmt.Printf("path :%v\n start generate enitity", path)
		// 查询表结构信息
		tableDesc, err := l.DB.GetPostgresTableDesc(table.Name)
		if err != nil {
			log.Fatal("CreateModelsEntityErr>>", err)
			continue
		}
		req := new(EntityReq)
		req.Index = idx
		req.EntityPath = path
		req.TableName = table.Name
		req.TableComment = table.Comment
		req.TableDesc = tableDesc
		req.FormatList = formatList
		req.EntityPkg = PkgModels
		// 生成基础信息
		err = l.GenerateModelsEntity(req)
		if err != nil {
			log.Fatal("CreateModelsEntityErr>>", err)
			continue
		}
	}
	return
}

// CreateGormCURD 生成gormCURD数据库操作
func (l *Logic) CreateGormCURD(formatList []string) (err error) {
	tableNameList := make([]*TableList, 0)
	// 表结构文件路径
	// 将表结构写入文件
	for _, table := range l.DB.DoTables {
		tableNameList = append(tableNameList, &TableList{
			UpperTableName: TablePrefix + l.T.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		// 查询表结构信息
		tableDesc, err := l.DB.GetPostgresTableDesc(table.Name)
		if err != nil {
			return err
		}
		req := new(EntityReq)
		req.TableName = table.Name
		req.TableComment = table.Comment
		req.TableDesc = tableDesc
		req.EntityPath = CreateDir(l.Path+GODIR_MODELS) + DS + table.Name + ".go"
		req.FormatList = formatList
		req.Pkg = PkgModels
		req.EntityPkg = PkgModels
		// 生成基础信息
		err = l.GenerateModelsEntity(req)
		if err != nil {
			return err
		}
		// 生成增,删,改,查文件
		err = l.GenerateGormCURDFile(table.Name, table.Comment, tableDesc)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateMarkdown 生成markdown文档
func (l *Logic) CreateMarkdown() (err error) {
	data := new(MarkDownData)
	// 将表结构写入文件
	i := 1
	for _, table := range l.DB.DoTables {
		fmt.Println("Doing table:" + table.Name)
		data.TableList = append(data.TableList, &TableList{
			Index:          i,
			UpperTableName: l.T.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		// 查询表结构信息
		desc := new(MarkDownDataChild)
		desc.List, err = l.DB.GetPostgresTableDesc(table.Name)
		//desc.List, err = l.DB.GetTableDesc(table.Name)
		if err != nil {
			log.Fatal("markdown>>", err)
			continue
		}
		desc.Index = i
		desc.TableName = table.Name
		desc.Comment = table.Comment
		data.DescList = append(data.DescList, desc)
		i++
	}
	// 生成所有表的文件
	err = l.GenerateMarkdown(data)
	if err != nil {
		return
	}
	return
}

// GetMysqlDir 创建和获取models目录
func (l *Logic) GetMysqlDir() string {
	return CreateDir(l.Path + GODIR_MODELS + DS)
}

// GetConfigDir 创建和获取MYSQL目录
func (l *Logic) GetConfigDir() string {
	return CreateDir(l.Path + GODIR_MODELS + DS + GODIR_Config + DS)
}

// GetEntityDir 创建和获取MYSQL目录
func (l *Logic) GetEntityDir() string {
	return CreateDir(l.Path + GODIR_MODELS + DS + GODIR_Entity + DS)
}

// GetRoot 获取根目录地址
func (l *Logic) GetRoot() string {
	return GetRootPath(l.Path) + DS
}

// GenerateModelsEntity 生成models层entity结构体
func (l *Logic) GenerateModelsEntity(req *EntityReq) (err error) {
	l.l.Lock()
	defer l.l.Unlock()
	fmt.Printf(">>> Generate entity start\n")
	var s string
	s = fmt.Sprintf(`
package %s
import (
	"errors"
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)
`, req.EntityPkg)
	// 判断import是否加载过
	check := fmt.Sprintf("package %s", req.EntityPkg)
	if l.T.CheckFileContainsChar(req.EntityPath, check) == false {
		l.T.WriteFile(req.EntityPath, s)
	}
	// 声明表结构变量
	TableData := new(TableInfo)
	TableData.Table = l.T.Capitalize(req.TableName)
	TableData.TableComment = AddToTableNameComment(TableData.Table, req.TableComment, "")
	// 判断表结构是否加载过
	if l.T.CheckFileContainsChar(req.EntityPath, "type "+TableData.Table+" struct") == true {
		log.Println(req.EntityPath + "It already exists. Please delete it and regenerate it")
		return
	}
	// 加载模板文件
	tplByte, err := Asset(TPL_MODELS)
	if err != nil {
		return
	}
	tpl, err := template.New("models").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 装载表字段信息
	for _, val := range req.TableDesc {
		TableData.Fields = append(TableData.Fields, &FieldsInfo{
			Name:         l.T.Capitalize(val.ColumnName),
			Type:         val.GolangType,
			DbOriField:   val.ColumnName,
			FormatFields: FormatField(val.ColumnName, req.FormatList),
			Remark:       AddToComment(val.ColumnComment, ""),
		})
	}
	content := bytes.NewBuffer([]byte{})
	tpl.Execute(content, TableData)
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteAppendFile(req.EntityPath, con)
	if err != nil {
		return
	}
	return
}

// GenerateGormCURDFile 生成gormCURD文件
func (l *Logic) GenerateGormCURDFile(tableName, tableComment string, tableDesc []*TableDesc) (err error) {
	l.l.Lock()
	defer l.l.Unlock()
	fmt.Printf(">>> Generate curd file start\n")
	var (
		allFields     = make([]string, 0)
		fieldsList    = make([]*SqlFieldInfo, 0)
		nullFieldList = make([]*NullSqlFieldInfo, 0)
	)
	for _, item := range tableDesc {
		allFields = append(allFields, item.ColumnName)
		fieldsList = append(fieldsList, &SqlFieldInfo{
			HumpName: l.T.Capitalize(item.ColumnName),
			Comment:  item.ColumnComment,
		})
		nullFieldList = append(nullFieldList, &NullSqlFieldInfo{
			HumpName:     l.T.Capitalize(item.ColumnName),
			OriFieldType: item.OriMysqlType,
			GoType:       MysqlTypeToGoType[item.OriMysqlType],
			Comment:      item.ColumnComment,
		})
	}

	sqlInfo := &SqlInfo{
		TableName:       tableName,
		StructTableName: l.T.Capitalize(tableName),
		UpperTableName:  TablePrefix + l.T.ToUpper(tableName),
		AllFieldList:    strings.Join(allFields, ","),
	}
	// 写入表名
	goFile := CreateDir(l.Path+GODIR_MODELS) + DS + tableName + ".go"
	s := fmt.Sprintf(`
package %s
import (
	"errors"
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)
`, PkgModels)
	// 判断package是否加载过
	checkStr := fmt.Sprintf("package %s", PkgModels)
	if l.T.CheckFileContainsChar(goFile, checkStr) == false {
		l.T.WriteFile(goFile, s)
	}

	// 解析模板
	tplByte, err := Asset(TPL_GORM)
	if err != nil {
		return
	}
	tpl, err := template.New("gorm").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, sqlInfo)
	if err != nil {
		return
	}
	// 表信息写入文件
	if l.T.CheckFileContainsChar(goFile, sqlInfo.StructTableName+"Model") == false {
		err = WriteAppendFile(goFile, content.String())
		if err != nil {
			fmt.Printf("err:%v\n", err)
			return
		}
	}

	if err != nil {
		return
	}
	return
}

// GenerateMarkdown 生成表列表
func (l *Logic) GenerateMarkdown(data *MarkDownData) (err error) {
	// 写入markdown
	file := l.Path + fmt.Sprintf("%s.md", l.DB.DBName)
	tplByte, err := Asset(TPL_MARKDOWN)
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	tpl, err := template.New("markdown").Parse(string(tplByte))
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}
