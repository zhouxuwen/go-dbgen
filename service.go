package main

import (
	"sort"
	"strings"
)

//GetTableNameAndComment 将表名赋值给结构对象, 供其它方法使用
func (m *ModelS) GetTableNameAndComment(tableName string) (err error) {
	//读取所有表列表
	if len(m.Tables) == 0 {
		m.Tables, err = m.findPostgresTables(tableName)
		//m.Tables, err = m.findDbTables()
		if err != nil {
			return
		}
	}
	//如果正在处理表数据为空,则将所有的表赋值给它
	if len(m.DoTables) == 0 {
		m.DoTables = m.Tables
	}
	return
}

// findPostgresTables 从设计
func (m *ModelS) findPostgresTables(tableName string) (NameAndComment []TableNameAndComment, err error) {
	m.l.Lock()
	defer m.l.Unlock()
	var result []map[string]interface{}
	if strings.ToLower(tableName) == "all" {
		result, err = m.Find(`select relname as table_name,cast(obj_description(relfilenode,'pg_class') as varchar) as table_comment 
		from pg_class c`)
	} else {
		result, err = m.Find(`select relname as table_name,cast(obj_description(relfilenode,'pg_class') as varchar) as table_comment 
		from pg_class c where relname = $1`, tableName)
	}
	if err != nil {
		return nil, err
	}
	//获取表名 与 表注释
	NameAndComment = make([]TableNameAndComment, 0)
	//获取库里所有的表名
	for idx, info := range result {
		idx++
		name, _ := info["table_name"].(string)
		comment, _ := info["table_comment"].(string)
		NameAndComment = append(NameAndComment, TableNameAndComment{
			Index:   idx,
			Name:    name,
			Comment: comment,
		})
	} //排序, 采用升序
	sort.Slice(NameAndComment, func(i, j int) bool {
		return strings.ToLower(NameAndComment[i].Name) < strings.ToLower(NameAndComment[j].Name)
	})
	return
}

func (m *ModelS) GetPostgresTableDesc(tableName string) (reply []*TableDesc, err error) {
	m.l.Lock()
	defer m.l.Unlock()
	result, err := m.Find(`select column_name, udt_name, is_nullable, column_default
	from information_schema.columns where table_name = $1 and table_catalog = $2`, tableName, m.DBName)
	if err != nil {
		return nil, err
	}
	result2, err := m.Find(`select distinct a.attname AS column_name,d.description AS column_comment
	from pg_class c, pg_attribute a, pg_description d
	where  c.relname = $1 and a.attnum>0 and a.attrelid = c.oid and  d.objoid=a.attrelid and d.objsubid=a.attnum`, tableName)
	if err != nil {
		return nil, err
	}
	ColumnDescriptionMap := make(map[string]string)
	for _, row2 := range result2 {
		ColumnDescriptionMap[row2["column_name"].(string)] = row2["column_comment"].(string)
	}

	reply = make([]*TableDesc, 0)
	i := 0
	for _, row := range result {
		oriType := row["udt_name"].(string)
		var columnDefault string
		val, ok := row["column_default"].(string)
		if ok {
			columnDefault = val
		}
		reply = append(reply, &TableDesc{
			Index:          i,
			ColumnName:     row["column_name"].(string),
			GoColumnName:   m.T.Capitalize(row["column_name"].(string)),
			OriMysqlType:   oriType,
			UpperMysqlType: strings.ToUpper(oriType),
			GolangType:     PostgresTypeToGoType[oriType],
			ColumnComment:  ColumnDescriptionMap[row["column_name"].(string)],
			IsNull:         row["is_nullable"].(string),
			DefaultValue:   columnDefault,
		})
		i++
	}
	return
}
