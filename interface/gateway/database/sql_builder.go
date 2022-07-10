package database

import (
	"reflect"
	"strings"
)

// SQLType ...
type SQLType uint

const (
	Select SQLType = iota + 1
	Insert
	Update
	Delete
)

// SQLBuilder ...
type SQLBuilder struct {
	sqlType    SQLType
	table      Table
	tableQuery string
	columns    []string
	conditions []string
	group      string
}

// NewSQLBuilder ...
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

// Statement ...
func (sb *SQLBuilder) Statement() string {
	var statement string
	switch sb.sqlType {
	case Select:
		columnQuery := strings.Join(sb.columns, ",")
		statementArray := []string{
			rw.Select, columnQuery,
			rw.From, sb.tableQuery,
		}
		if len(sb.conditions) != 0 {
			conditionQuery := strings.Join(sb.conditions, " "+rw.And+" ")
			statementArray = append(statementArray, rw.Where, conditionQuery)
		}
		if sb.group != "" {
			statementArray = append(statementArray, rw.GroupBy, sb.group)
		}
		statement = strings.Join(statementArray, " ")
	case Delete:
		statementArray := []string{
			rw.Delete, sb.table.NAME(),
			rw.From, sb.tableQuery,
		}
		if len(sb.conditions) != 0 {
			conditionQuery := strings.Join(sb.conditions, " "+rw.And+" ")
			statementArray = append(statementArray, rw.Where, conditionQuery)
		}
		statement = strings.Join(statementArray, " ")
	default:
	}
	return statement
}

// Insert ...
func (sb *SQLBuilder) Insert(table Table) string {
	sb.sqlType = Insert
	sb.setColumns(table)
	columnQuery := strings.Join(sb.columns, ",")
	valuesQuery := "(?" + strings.Repeat(",?", len(sb.columns)-1) + ")"
	statement := strings.Join([]string{
		rw.InsertInto, table.NAME(), "(", columnQuery, ")",
		rw.Values, valuesQuery,
	}, " ")
	return statement
}

// Update ...
func (sb *SQLBuilder) Update(table Table) string {
	sb.sqlType = Update
	sb.setColumns(table)
	key := sb.key(table)
	columnQuery := strings.Join(sb.columns, " = ?,")
	columnQuery += " = ?"
	statement := strings.Join([]string{
		rw.Update, table.NAME(),
		rw.Set, columnQuery,
		rw.Where, key, "= ?",
	}, " ")
	return statement
}

// Delete ...
func (sb *SQLBuilder) Delete(table Table) *SQLBuilder {
	sb.sqlType = Delete
	sb.table = table
	sb.tableQuery = table.NAME()
	return sb
}

// Select ...
func (sb *SQLBuilder) Select(table Table) *SQLBuilder {
	sb.sqlType = Select
	sb.table = table
	sb.tableQuery = table.NAME()
	sb.setColumns(table)
	return sb
}

// Column ...
func (sb *SQLBuilder) Column(column string) *SQLBuilder {
	sb.columns = append(sb.columns, column)
	return sb
}

// Columns ...
func (sb *SQLBuilder) Columns(table Table) *SQLBuilder {
	sb.setColumns(table)
	return sb
}

// LeftJoin ...
func (sb *SQLBuilder) LeftJoin(table Table, l, r string) *SQLBuilder {
	sb.tableQuery = strings.Join([]string{
		sb.tableQuery,
		rw.LeftJoin, table.NAME(),
		rw.On, l, "=", r,
	}, " ")
	return sb
}

// LeftJoinWithColumns ...
func (sb *SQLBuilder) LeftJoinWithColumns(table Table, l, r string) *SQLBuilder {
	sb.tableQuery = strings.Join([]string{
		sb.tableQuery,
		rw.LeftJoin, table.NAME(),
		rw.On, l, "=", r,
	}, " ")
	sb.setColumns(table)
	return sb
}

// RightJoin ...
func (sb *SQLBuilder) RightJoin(table Table, l, r string) *SQLBuilder {
	sb.tableQuery = strings.Join([]string{
		sb.tableQuery,
		rw.RightJoin, table.NAME(),
		rw.On, l, "=", r,
	}, " ")
	return sb
}

// RightJoinWithColumns ...
func (sb *SQLBuilder) RightJoinWithColumns(table Table, l, r string) *SQLBuilder {
	sb.tableQuery = strings.Join([]string{
		sb.tableQuery,
		rw.RightJoin, table.NAME(),
		rw.On, l, "=", r,
	}, " ")
	sb.setColumns(table)
	return sb
}

// Match ...
func (sb *SQLBuilder) Match(column string) *SQLBuilder {
	sb.conditions = append(sb.conditions, column+" = ?")
	return sb
}

// Like ...
func (sb *SQLBuilder) Like(column string) *SQLBuilder {
	sb.conditions = append(sb.conditions, strings.Join([]string{column, rw.Like, "?"}, " "))
	return sb
}

// GroupBy ...
func (sb *SQLBuilder) GroupBy(column string) *SQLBuilder {
	sb.group = column
	return sb
}

func (sb *SQLBuilder) setColumns(table Table) {
	tv := reflect.Indirect(reflect.ValueOf(table))
	for i := 0; i < tv.NumField(); i++ {
		f := tv.Field(i)
		if f.Type().Kind() != reflect.String {
			continue
		}
		if strings.Contains(f.String(), "created_at") || strings.Contains(f.String(), "updated_at") {
			continue
		}
		sb.columns = append(sb.columns, f.String())
	}
}

func (sb *SQLBuilder) key(table Table) string {
	tv := reflect.Indirect(reflect.ValueOf(table))
	return tv.Field(1).String()
}
