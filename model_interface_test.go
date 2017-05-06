package lore

import "testing"

const (
	_TEST_DB_TABLENAME          string = "tests"
	_TEST_DB_FIELDNAME_ID       string = "id"
	_TEST_DB_FIELDNAME_FIELDONE string = "field_one"
	_TEST_DB_FIELDNAME_FIELDTWO string = "field_two"
	_TEST_MODEL_ID              int    = 1
	_TEST_MODEL_FIELD_ONE       int64  = 1
	_TEST_MODEL_FIELD_TWO       string = "Two"
)

type testModel struct {
	Id       int    `db:"id"`
	FieldOne int64  `db:"field_one"`
	FieldTwo string `db:"field_two"`
}

/*
Enforce interface.
*/
var _ ModelInterface = (*testModel)(nil)

/*
DbTableName implementation for ModelInterface.
*/
func (*testModel) DbTableName() string {
	return _TEST_DB_TABLENAME
}

/*
DbFieldMap implementation for ModelInterface.
*/
func (tm *testModel) DbFieldMap() map[string]interface{} {
	return map[string]interface{}{
		_TEST_DB_FIELDNAME_FIELDONE: tm.FieldOne,
		_TEST_DB_FIELDNAME_FIELDTWO: tm.FieldTwo,
	}
}

/*
DbPrimaryFieldKey implementation for ModelInterface.
*/
func (*testModel) DbPrimaryFieldKey() string {
	return _TEST_DB_FIELDNAME_ID
}

/*
DbPimraryFieldValue implementation for ModelInterface.
*/
func (tm *testModel) DbPrimaryFieldValue() interface{} {
	return _TEST_MODEL_ID
}

/*
createTestModelInstance creates a new testModel instance with valid values.
*/
func createTestModelInstance() *testModel {
	return &testModel{
		Id:       _TEST_MODEL_ID,
		FieldOne: _TEST_MODEL_FIELD_ONE,
		FieldTwo: _TEST_MODEL_FIELD_TWO,
	}
}

/*
TestModelInterfaceInstance tests the ModelInterface interface by creating a new model instance. This
is a trivial test.
*/
func TestModelInterfaceInstance(t *testing.T) {
	tm := createTestModelInstance()
	var mi ModelInterface
	mi = tm

	// Test table name.
	dbTableName := mi.DbTableName()
	if dbTableName != _TEST_DB_TABLENAME {
		t.Fatalf("Invalid DbTableName: %s, expected %s", dbTableName, _TEST_DB_TABLENAME)
	}

	// Test primary key/value.
	primaryFieldKey := mi.DbPrimaryFieldKey()
	primaryFieldValue := mi.DbPrimaryFieldValue()
	if primaryFieldKey != _TEST_DB_FIELDNAME_ID || primaryFieldValue != _TEST_MODEL_ID {
		t.Fatalf("Invalid primary key/value combination: (%s, %+v), expected (%s, %+v)", primaryFieldKey, primaryFieldValue, _TEST_DB_FIELDNAME_ID, _TEST_MODEL_ID)
	}

	// Test field map.
	dbFieldMap := mi.DbFieldMap()
	if len(dbFieldMap) != 2 || dbFieldMap[_TEST_DB_FIELDNAME_FIELDONE] != tm.FieldOne || dbFieldMap[_TEST_DB_FIELDNAME_FIELDTWO] != tm.FieldTwo {
		t.Fatalf("Invalid DbFieldMap: %+v", dbFieldMap)
	}
}

/*
newTestModelEmpty returns a pointer to a new, empty instance of testModel.
*/
func newTestModelEmpty() *testModel {
	return &testModel{}
}

/*
newTestModelEmptyList returns a pointer to a new, empty list for testModels.
*/
func newTestModelEmptyList() *[]testModel {
	return &[]testModel{}
}