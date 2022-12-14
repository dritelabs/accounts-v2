// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/dritelabs/accounts/internal/database/models"
)

func newProfile(db *gorm.DB, opts ...gen.DOOption) profile {
	_profile := profile{}

	_profile.profileDo.UseDB(db, opts...)
	_profile.profileDo.UseModel(&models.Profile{})

	tableName := _profile.profileDo.TableName()
	_profile.ALL = field.NewAsterisk(tableName)
	_profile.ID = field.NewField(tableName, "id")
	_profile.CreatedAt = field.NewTime(tableName, "created_at")
	_profile.UpdatedAt = field.NewTime(tableName, "updated_at")
	_profile.DeletedAt = field.NewField(tableName, "deleted_at")
	_profile.UserID = field.NewField(tableName, "user_id")
	_profile.BirthDate = field.NewTime(tableName, "birth_date")
	_profile.FirstName = field.NewField(tableName, "first_name")
	_profile.Gender = field.NewField(tableName, "gender")
	_profile.Locale = field.NewField(tableName, "locale")
	_profile.LastName = field.NewField(tableName, "last_name")
	_profile.MiddleName = field.NewField(tableName, "middle_name")
	_profile.Nickname = field.NewField(tableName, "nickname")
	_profile.Profile = field.NewField(tableName, "profile")
	_profile.Picture = field.NewField(tableName, "picture")
	_profile.Website = field.NewField(tableName, "website")
	_profile.ZoneInfo = field.NewField(tableName, "zone_info")
	_profile.User = profileBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "models.User"),
		DefaultShippingAddress: struct {
			field.RelationField
			User struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("User.DefaultShippingAddress", "models.Address"),
			User: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.DefaultShippingAddress.User", "models.User"),
			},
		},
		DefaultBillingAddress: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.DefaultBillingAddress", "models.Address"),
		},
		Profile: struct {
			field.RelationField
			User struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("User.Profile", "models.Profile"),
			User: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.Profile.User", "models.User"),
			},
		},
		Addresses: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.Addresses", "models.Address"),
		},
		Clients: struct {
			field.RelationField
			User struct {
				field.RelationField
			}
			Approvals struct {
				field.RelationField
				Scopes struct {
					field.RelationField
					Approvals struct {
						field.RelationField
					}
					Clients struct {
						field.RelationField
					}
				}
				Clients struct {
					field.RelationField
				}
				Users struct {
					field.RelationField
				}
			}
			Scopes struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("User.Clients", "models.Client"),
			User: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.Clients.User", "models.User"),
			},
			Approvals: struct {
				field.RelationField
				Scopes struct {
					field.RelationField
					Approvals struct {
						field.RelationField
					}
					Clients struct {
						field.RelationField
					}
				}
				Clients struct {
					field.RelationField
				}
				Users struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("User.Clients.Approvals", "models.Approval"),
				Scopes: struct {
					field.RelationField
					Approvals struct {
						field.RelationField
					}
					Clients struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("User.Clients.Approvals.Scopes", "models.Scope"),
					Approvals: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("User.Clients.Approvals.Scopes.Approvals", "models.Approval"),
					},
					Clients: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("User.Clients.Approvals.Scopes.Clients", "models.Client"),
					},
				},
				Clients: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("User.Clients.Approvals.Clients", "models.Client"),
				},
				Users: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("User.Clients.Approvals.Users", "models.User"),
				},
			},
			Scopes: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.Clients.Scopes", "models.Scope"),
			},
		},
		Approvals: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.Approvals", "models.Approval"),
		},
	}

	_profile.fillFieldMap()

	return _profile
}

type profile struct {
	profileDo

	ALL        field.Asterisk
	ID         field.Field
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	UserID     field.Field
	BirthDate  field.Time
	FirstName  field.Field
	Gender     field.Field
	Locale     field.Field
	LastName   field.Field
	MiddleName field.Field
	Nickname   field.Field
	Profile    field.Field
	Picture    field.Field
	Website    field.Field
	ZoneInfo   field.Field
	User       profileBelongsToUser

	fieldMap map[string]field.Expr
}

func (p profile) Table(newTableName string) *profile {
	p.profileDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p profile) As(alias string) *profile {
	p.profileDo.DO = *(p.profileDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *profile) updateTableName(table string) *profile {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewField(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.UserID = field.NewField(table, "user_id")
	p.BirthDate = field.NewTime(table, "birth_date")
	p.FirstName = field.NewField(table, "first_name")
	p.Gender = field.NewField(table, "gender")
	p.Locale = field.NewField(table, "locale")
	p.LastName = field.NewField(table, "last_name")
	p.MiddleName = field.NewField(table, "middle_name")
	p.Nickname = field.NewField(table, "nickname")
	p.Profile = field.NewField(table, "profile")
	p.Picture = field.NewField(table, "picture")
	p.Website = field.NewField(table, "website")
	p.ZoneInfo = field.NewField(table, "zone_info")

	p.fillFieldMap()

	return p
}

func (p *profile) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *profile) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 17)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["birth_date"] = p.BirthDate
	p.fieldMap["first_name"] = p.FirstName
	p.fieldMap["gender"] = p.Gender
	p.fieldMap["locale"] = p.Locale
	p.fieldMap["last_name"] = p.LastName
	p.fieldMap["middle_name"] = p.MiddleName
	p.fieldMap["nickname"] = p.Nickname
	p.fieldMap["profile"] = p.Profile
	p.fieldMap["picture"] = p.Picture
	p.fieldMap["website"] = p.Website
	p.fieldMap["zone_info"] = p.ZoneInfo

}

func (p profile) clone(db *gorm.DB) profile {
	p.profileDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p profile) replaceDB(db *gorm.DB) profile {
	p.profileDo.ReplaceDB(db)
	return p
}

type profileBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	DefaultShippingAddress struct {
		field.RelationField
		User struct {
			field.RelationField
		}
	}
	DefaultBillingAddress struct {
		field.RelationField
	}
	Profile struct {
		field.RelationField
		User struct {
			field.RelationField
		}
	}
	Addresses struct {
		field.RelationField
	}
	Clients struct {
		field.RelationField
		User struct {
			field.RelationField
		}
		Approvals struct {
			field.RelationField
			Scopes struct {
				field.RelationField
				Approvals struct {
					field.RelationField
				}
				Clients struct {
					field.RelationField
				}
			}
			Clients struct {
				field.RelationField
			}
			Users struct {
				field.RelationField
			}
		}
		Scopes struct {
			field.RelationField
		}
	}
	Approvals struct {
		field.RelationField
	}
}

func (a profileBelongsToUser) Where(conds ...field.Expr) *profileBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a profileBelongsToUser) WithContext(ctx context.Context) *profileBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a profileBelongsToUser) Model(m *models.Profile) *profileBelongsToUserTx {
	return &profileBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type profileBelongsToUserTx struct{ tx *gorm.Association }

func (a profileBelongsToUserTx) Find() (result *models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a profileBelongsToUserTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a profileBelongsToUserTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a profileBelongsToUserTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a profileBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a profileBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type profileDo struct{ gen.DO }

type IProfileDo interface {
	gen.SubQuery
	Debug() IProfileDo
	WithContext(ctx context.Context) IProfileDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProfileDo
	WriteDB() IProfileDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProfileDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProfileDo
	Not(conds ...gen.Condition) IProfileDo
	Or(conds ...gen.Condition) IProfileDo
	Select(conds ...field.Expr) IProfileDo
	Where(conds ...gen.Condition) IProfileDo
	Order(conds ...field.Expr) IProfileDo
	Distinct(cols ...field.Expr) IProfileDo
	Omit(cols ...field.Expr) IProfileDo
	Join(table schema.Tabler, on ...field.Expr) IProfileDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProfileDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProfileDo
	Group(cols ...field.Expr) IProfileDo
	Having(conds ...gen.Condition) IProfileDo
	Limit(limit int) IProfileDo
	Offset(offset int) IProfileDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProfileDo
	Unscoped() IProfileDo
	Create(values ...*models.Profile) error
	CreateInBatches(values []*models.Profile, batchSize int) error
	Save(values ...*models.Profile) error
	First() (*models.Profile, error)
	Take() (*models.Profile, error)
	Last() (*models.Profile, error)
	Find() ([]*models.Profile, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Profile, err error)
	FindInBatches(result *[]*models.Profile, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Profile) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProfileDo
	Assign(attrs ...field.AssignExpr) IProfileDo
	Joins(fields ...field.RelationField) IProfileDo
	Preload(fields ...field.RelationField) IProfileDo
	FirstOrInit() (*models.Profile, error)
	FirstOrCreate() (*models.Profile, error)
	FindByPage(offset int, limit int) (result []*models.Profile, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProfileDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []models.Profile, err error)
}

// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (p profileDo) FilterWithNameAndRole(name string, role string) (result []models.Profile, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM profiles WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB

	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result)
	err = executeSQL.Error
	return
}

func (p profileDo) Debug() IProfileDo {
	return p.withDO(p.DO.Debug())
}

func (p profileDo) WithContext(ctx context.Context) IProfileDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p profileDo) ReadDB() IProfileDo {
	return p.Clauses(dbresolver.Read)
}

func (p profileDo) WriteDB() IProfileDo {
	return p.Clauses(dbresolver.Write)
}

func (p profileDo) Session(config *gorm.Session) IProfileDo {
	return p.withDO(p.DO.Session(config))
}

func (p profileDo) Clauses(conds ...clause.Expression) IProfileDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p profileDo) Returning(value interface{}, columns ...string) IProfileDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p profileDo) Not(conds ...gen.Condition) IProfileDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p profileDo) Or(conds ...gen.Condition) IProfileDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p profileDo) Select(conds ...field.Expr) IProfileDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p profileDo) Where(conds ...gen.Condition) IProfileDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p profileDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IProfileDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p profileDo) Order(conds ...field.Expr) IProfileDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p profileDo) Distinct(cols ...field.Expr) IProfileDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p profileDo) Omit(cols ...field.Expr) IProfileDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p profileDo) Join(table schema.Tabler, on ...field.Expr) IProfileDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p profileDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProfileDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p profileDo) RightJoin(table schema.Tabler, on ...field.Expr) IProfileDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p profileDo) Group(cols ...field.Expr) IProfileDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p profileDo) Having(conds ...gen.Condition) IProfileDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p profileDo) Limit(limit int) IProfileDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p profileDo) Offset(offset int) IProfileDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p profileDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProfileDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p profileDo) Unscoped() IProfileDo {
	return p.withDO(p.DO.Unscoped())
}

func (p profileDo) Create(values ...*models.Profile) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p profileDo) CreateInBatches(values []*models.Profile, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p profileDo) Save(values ...*models.Profile) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p profileDo) First() (*models.Profile, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Profile), nil
	}
}

func (p profileDo) Take() (*models.Profile, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Profile), nil
	}
}

func (p profileDo) Last() (*models.Profile, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Profile), nil
	}
}

func (p profileDo) Find() ([]*models.Profile, error) {
	result, err := p.DO.Find()
	return result.([]*models.Profile), err
}

func (p profileDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Profile, err error) {
	buf := make([]*models.Profile, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p profileDo) FindInBatches(result *[]*models.Profile, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p profileDo) Attrs(attrs ...field.AssignExpr) IProfileDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p profileDo) Assign(attrs ...field.AssignExpr) IProfileDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p profileDo) Joins(fields ...field.RelationField) IProfileDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p profileDo) Preload(fields ...field.RelationField) IProfileDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p profileDo) FirstOrInit() (*models.Profile, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Profile), nil
	}
}

func (p profileDo) FirstOrCreate() (*models.Profile, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Profile), nil
	}
}

func (p profileDo) FindByPage(offset int, limit int) (result []*models.Profile, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p profileDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p profileDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p profileDo) Delete(models ...*models.Profile) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *profileDo) withDO(do gen.Dao) *profileDo {
	p.DO = *do.(*gen.DO)
	return p
}
