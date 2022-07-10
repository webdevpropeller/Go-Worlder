package database

import (
	"go_worlder_system/str"
	"reflect"
	"strings"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

const (
	// Values
	trueVal  string = "true"
	falseVal string = "false"
	// column
	stock string = "stock"
	// User category ids
	brandOwnerID int = 1
	partnerID    int = 2
	// InventoryType
	receivingType int = 1
	shippingType  int = 2
	faultyType    int = 3
	// keys
	msgKey = "Message"
	idKey  = "ID"
	// format
	timeFormat = "2006-01-02"
)

var (
	// DB configurations
	masterDB                       = NewMasterDB()
	userDB                         = NewUserDB()
	accountDB                      = NewAccountDB()
	brandDB                        = NewBrandDB()
	categoryDB                     = NewCategoryDB()
	inventoryDB                    = NewInventoryDB()
	orderDB                        = NewOrderDB()
	productDB                      = NewProductDB()
	projectDB                      = NewProjectDB()
	transactionDB                  = NewTransactionDB()
	chatDB                         = NewChatDB()
	customeUserTable               = NewCustomeUserTable()
	customeTransactionPartnerTable = NewCustomeTransactionPartnerTable()
	customeTransactionAccountTable = NewCustomeTransactionAccountTable()
	// reserved word
	rw = NewReservedWord()
)

func generateSelectColumnQuery(columnList []string) string {
	return strings.Join(columnList, ",")
}

func generateInsertColumnQuery(columnList []string) string {
	columnQuery := strings.Join(columnList, ",")
	return "(" + columnQuery + ")"
}

func generateUpdateColumnQuery(columnList []string) string {
	columnQuery := strings.Join(columnList, "= ?,")
	return strings.TrimRight(columnQuery, ",")
}

func generateValuesQuery(valuesNum int) (valuesQuery string) {
	valuesQueryArray := []string{"(?"}
	for i := 0; i < valuesNum-1; i++ {
		valuesQueryArray = append(valuesQueryArray, ",?")
	}
	valuesQueryArray = append(valuesQueryArray, ")")
	valuesQuery = strings.Join(valuesQueryArray, "")
	return
}

func generateColumnList(args ...interface{}) []string {
	columnList := []string{}
	for _, arg := range args {
		rv := reflect.Indirect(reflect.ValueOf(arg))
		if rv.Type().Kind() == reflect.String {
			columnList = append(columnList, rv.String())
			continue
		}
		for i := 1; i < rv.NumField(); i++ {
			f := rv.Field(i)
			if strings.Contains(f.String(), "created_at") || strings.Contains(f.String(), "updated_at") {
				continue
			}
			columnList = append(columnList, f.String())
		}
	}
	return columnList
}

func generateValueList(model interface{}) []interface{} {
	rv := reflect.Indirect(reflect.ValueOf(model))
	valueList := []interface{}{}
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		var value interface{}
		if fv.Type().Kind() == reflect.Struct || fv.Type().Kind() == reflect.Ptr {
			switch fv.Type().Elem().Name() {
			case "FileHeader":
				continue
			case "Time":
				method, ok := fv.Type().MethodByName("Format")
				if !ok {
					continue
				}
				rtn := method.Func.Call([]reflect.Value{fv, reflect.ValueOf(timeFormat)})
				value = rtn[0].Interface().(string)
			default:
				if fv.Elem().Kind() != reflect.Struct {
					continue
				}
				fv = fv.Elem().FieldByName("ID")
				if !fv.IsValid() {
					continue
				}
				value = fv.Interface()
			}
		} else {
			value = fv.Interface()
		}
		valueList = append(valueList, value)
	}

	return valueList
}

// Pass a pointer as an argument
func generateScanList(model interface{}) []interface{} {
	rv := reflect.ValueOf(model)
	if rv.Type().Kind() != reflect.Ptr {
		log.Panic("The argument is not a pointer")
		return nil
	}
	initializePointerField(rv)
	scanList := assignPointer(rv)
	return scanList
}

func initializePointerField(rv reflect.Value) {
	rv = reflect.Indirect(rv)
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		if fv.Kind() != reflect.Ptr {
			continue
		}
		if fv.CanSet() {
			new := reflect.New(fv.Type().Elem())
			fv.Set(new)
		}
		if fv.Elem().Kind() != reflect.Struct {
			continue
		}
		initializePointerField(fv)
	}
}

func assignPointer(rv reflect.Value) []interface{} {
	scanList := []interface{}{}
	afterList := []reflect.Value{}
	rv = reflect.Indirect(rv)
	if rv.Type().Kind() != reflect.Struct {
		pointer := rv.Addr().Interface()
		scanList = append(scanList, pointer)
		return scanList
	}
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		if fv.Type().Kind() == reflect.Slice {
			continue
		} else if fv.Type().Kind() == reflect.Ptr {
			switch fv.Elem().Type().Name() {
			case "FileHeader":
				continue
			default:
				if fv.Elem().Kind() == reflect.Struct {
					afterList = append(afterList, fv)
					fv = fv.Elem().FieldByName("ID")
					if !fv.IsValid() {
						continue
					}
				} else {
					afterList = append(afterList, fv)
					continue
				}
			}
		}
		pointer := fv.Addr().Interface()
		scanList = append(scanList, pointer)
	}
	for _, after := range afterList {
		afterScanList := assignPointer(after)
		scanList = append(scanList, afterScanList...)
	}
	return scanList
}

func initialize(db DB) {
	dbv := reflect.Indirect(reflect.ValueOf(db))
	dbnamefv := dbv.FieldByName("db").FieldByName("name")
	dbName := str.ToSnakeCase(strings.TrimSuffix(dbv.Type().Name(), "DB"))
	dbnamePtr := (*string)(unsafe.Pointer(dbnamefv.UnsafeAddr()))
	*dbnamePtr = dbName
	for i := 1; i < dbv.NumField(); i++ {
		tv := dbv.Field(i)
		tablenamefv := tv.FieldByName("table").FieldByName("name")
		tableName := str.ToSnakeCase(strings.TrimSuffix(tv.Type().Name(), "Table"))
		tablenamePtr := (*string)(unsafe.Pointer(tablenamefv.UnsafeAddr()))
		*tablenamePtr = strings.Join([]string{dbName, tableName}, ".")
		for i := 1; i < tv.NumField(); i++ {
			cv := tv.Type().Field(i)
			columnName := str.ToSnakeCase(cv.Name)
			tv.Field(i).SetString(strings.Join([]string{tableName, columnName}, "."))
		}
	}
}

func alias(table Table, alias string) Table {
	table.SetAlias(alias)
	tv := reflect.Indirect(reflect.ValueOf(table))
	for i := 0; i < tv.NumField(); i++ {
		cv := tv.Type().Field(i)
		if cv.Type.Kind() != reflect.String {
			continue
		}
		tv.Field(i).SetString(strings.Join([]string{alias, str.ToSnakeCase(cv.Name)}, "."))
	}
	return table
}
