package database

import (
	"go_worlder_system/str"
	"reflect"
	"strings"
)

// ReservedWord is reserved words of MySQL
type ReservedWord struct {
	Select      string
	From        string
	Where       string
	In          string
	Like        string
	GroupBy     string
	OrderBy     string
	Asc         string
	Desc        string
	Limit       string
	And         string
	Or          string
	InsertInto  string
	Values      string
	Update      string
	Set         string
	Delete      string
	Distinct    string
	Ifnull      string
	Case        string
	When        string
	Then        string
	Else        string
	End         string
	IsNull      string
	As          string
	InnerJoin   string
	OuterJoin   string
	LeftJoin    string
	RightJoin   string
	On          string
	Over        string
	PartitionBy string
	Union       string
	UnionAll    string
	Intersect   string
	Minus       string
}

// NewReservedWord ...
func NewReservedWord() *ReservedWord {
	rw := &ReservedWord{}
	rv := reflect.Indirect(reflect.ValueOf(rw))
	for i := 0; i < rv.NumField(); i++ {
		sf := rv.Type().Field(i)
		word := strings.ToUpper(str.Split(sf.Name))
		rv.Field(i).SetString(word)
	}
	return rw
}
