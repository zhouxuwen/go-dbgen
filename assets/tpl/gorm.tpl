
type {{.StructTableName}}Model struct {
	DB *gorm.DB
	P  *{{.StructTableName}}
}

var (
	DBIsNilErr = errors.New("db is nil")
)

func (s *{{.StructTableName}}) TableName() string {
	return "{{.TableName}}"
}

func (s *{{.StructTableName}}Model) GetAllColumns() string {
	return ` {{.AllFieldList}} `
}

// New{{.StructTableName}}Model 新建Master2slaveModel对象
func New{{.StructTableName}}Model(db *gorm.DB) *{{.StructTableName}}Model {
	return &{{.StructTableName}}Model{
		DB: db,
		P:  &{{.StructTableName}}{},
	}
}

// Create 新增记录
func (s *{{.StructTableName}}Model) Create(value {{.StructTableName}}) error {
	if s.DB == nil {
		return DBIsNilErr
	}
	if err := s.DB.Debug().Table(s.P.TableName()).Save(&value).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除指定记录
func (s *{{.StructTableName}}Model) Delete(deleteWhereMap map[string]interface{}) error {
	if s.DB == nil {
		return DBIsNilErr
	}
	db := s.DB.Debug().Table(s.P.TableName())
	for key, value := range deleteWhereMap {
		whereStr := fmt.Sprintf("%v = ?", key)
		db = db.Where(whereStr, value)
	}
	if err := db.Delete({{.StructTableName}}{}).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新指定记录
func (s *{{.StructTableName}}Model) Update(updateMap map[string]interface{}, whereMap map[string]interface{}) error {
	if s.DB == nil {
		return DBIsNilErr
	}
	db := s.DB.Debug().Table(s.P.TableName())

	for key, value := range whereMap {
		whereStr := fmt.Sprintf("%v = ?", key)
		db = db.Where(whereStr, value)
	}

	if err := db.Update(updateMap).Error; err != nil {
		return err
	}
	return nil
}

// SelectAllColumns 查询所有记录
func (s *{{.StructTableName}}Model) SelectAllColumns() ([]{{.StructTableName}}, error) {
	if s.DB == nil {
		return nil, DBIsNilErr
	}
	objs := make([]{{.StructTableName}}, 0)
	if err := s.DB.Debug().Table(s.P.TableName()).
		Select(s.GetAllColumns()).
		Find(&objs).Error; err != nil {
			return nil, err
	}
	return objs, nil
}