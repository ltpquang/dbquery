package dbquery

import (
	"database/sql"
	"database/sql/driver"
	"gorm.io/driver/mysql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type testingStruct struct{}

type ExecutorTestSuite struct {
	suite.Suite
	DB       *gorm.DB
	mock     sqlmock.Sqlmock
	executor *gormQueryExecutor

	query            Query
	expectedQueryStr string
	expectedArgs     []interface{}
}

func newSuit(query Query, expectedQueryStr string, expectedArgs []interface{}) *ExecutorTestSuite {
	return &ExecutorTestSuite{
		query:            query,
		expectedQueryStr: expectedQueryStr,
		expectedArgs:     expectedArgs,
	}
}

func (s *ExecutorTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(s.T(), err)

	db.Driver()
	s.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.executor = WithGorm(s.DB).(*gormQueryExecutor)
}

func (s *ExecutorTestSuite) TestExecute() {
	args := make([]driver.Value, len(s.expectedArgs))
	for i, arg := range s.expectedArgs {
		args[i] = driver.Value(arg)
	}

	s.mock.ExpectQuery(s.expectedQueryStr).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{})).
		WillReturnError(nil)

	_, err := s.executor.Execute(s.query)

	require.NoError(s.T(), err)
}

func (s *ExecutorTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestGormQueryExecutorFindAll_Execute(t *testing.T) {
	findResult := make([]*testingStruct, 0)
	query := NewBuilder().
		Model(&testingStruct{}).
		Preload("field_a").
		Preload("field_b").
		Preload("field_c").
		Find(&findResult).
		Sort("sorted_field", OrderDesc).
		Offset(10).
		Limit(20).
		InList("field_in_list", []interface{}{"this", "is", "a", "list"}).
		Equal("fieldEqual", 11).
		GreaterOrEqual("fieldGreaterOrEqual", 12).
		Greater("fieldGreater", 13).
		LessOrEqual("fieldLessOrEqual", 14).
		Less("fieldLess", 15).
		NotEqual("fieldNotEqual", 16).
		Not("fieldNot", 17).
		Build()

	expectedQueryStr := "SELECT * FROM `testing_structs` WHERE fieldEqual = ? AND (fieldGreaterOrEqual >= ?) AND fieldGreater > ? AND (fieldLessOrEqual <= ?) AND fieldLess < ? AND fieldNotEqual <> ? AND `fieldNot` <> ? AND field_in_list in (?,?,?,?) ORDER BY `sorted_field` DESC LIMIT 20 OFFSET 10"
	expectedArgs := []interface{}{11, 12, 13, 14, 15, 16, 17, "this", "is", "a", "list"}
	suite.Run(t, newSuit(query, expectedQueryStr, expectedArgs))

	var countResult int64
	query.Builder().Count(&countResult).Build()
	expectedQueryStr = "SELECT count(1) FROM `testing_structs` WHERE fieldEqual = ? AND (fieldGreaterOrEqual >= ?) AND fieldGreater > ? AND (fieldLessOrEqual <= ?) AND fieldLess < ? AND fieldNotEqual <> ? AND `fieldNot` <> ? AND field_in_list in (?,?,?,?)"
	suite.Run(t, newSuit(query, expectedQueryStr, expectedArgs))
}

func TestGormQueryExecutorFindFirst_Execute(t *testing.T) {
	firstResult := &testingStruct{}
	query := NewBuilder().
		Model(&testingStruct{}).
		Preload("field_a").
		Preload("field_b").
		Preload("field_c").
		First(firstResult).
		Sort("sorted_field", OrderDesc).
		Offset(10).
		Limit(20).
		InList("field_in_list", []interface{}{"this", "is", "a", "list"}).
		Equal("fieldEqual", 11).
		GreaterOrEqual("fieldGreaterOrEqual", 12).
		Greater("fieldGreater", 13).
		LessOrEqual("fieldLessOrEqual", 14).
		Less("fieldLess", 15).
		NotEqual("fieldNotEqual", 16).
		Not("fieldNot", 17).
		Build()

	expectedQueryStr := "SELECT * FROM `testing_structs` WHERE fieldEqual = ? AND (fieldGreaterOrEqual >= ?) AND fieldGreater > ? AND (fieldLessOrEqual <= ?) AND fieldLess < ? AND fieldNotEqual <> ? AND `fieldNot` <> ? AND field_in_list in (?,?,?,?) ORDER BY `sorted_field` DESC,`testing_structs`. LIMIT 1 OFFSET 10"
	expectedArgs := []interface{}{11, 12, 13, 14, 15, 16, 17, "this", "is", "a", "list"}
	suite.Run(t, newSuit(query, expectedQueryStr, expectedArgs))

	var countResult int64
	query.Builder().Count(&countResult).Build()
	expectedQueryStr = "SELECT count(1) FROM `testing_structs` WHERE fieldEqual = ? AND (fieldGreaterOrEqual >= ?) AND fieldGreater > ? AND (fieldLessOrEqual <= ?) AND fieldLess < ? AND fieldNotEqual <> ? AND `fieldNot` <> ? AND field_in_list in (?,?,?,?)"
	suite.Run(t, newSuit(query, expectedQueryStr, expectedArgs))
}

func TestGormQueryExecutorTableName_Execute(t *testing.T) {
	firstResult := &testingStruct{}
	query := NewBuilder().
		TableName("a_table_name").
		Model(&testingStruct{}).
		Preload("field_a").
		Preload("field_b").
		Preload("field_c").
		First(firstResult).
		Sort("sorted_field", OrderDesc).
		Offset(10).
		Limit(20).
		InList("field_in_list", []interface{}{"this", "is", "a", "list"}).
		Equal("fieldEqual", 11).
		GreaterOrEqual("fieldGreaterOrEqual", 12).
		Greater("fieldGreater", 13).
		LessOrEqual("fieldLessOrEqual", 14).
		Less("fieldLess", 15).
		NotEqual("fieldNotEqual", 16).
		Not("fieldNot", 17).
		Build()

	expectedQueryStr := "SELECT * FROM `a_table_name` WHERE fieldEqual = ? AND (fieldGreaterOrEqual >= ?) AND fieldGreater > ? AND (fieldLessOrEqual <= ?) AND fieldLess < ? AND fieldNotEqual <> ? AND `fieldNot` <> ? AND field_in_list in (?,?,?,?) ORDER BY `sorted_field` DESC,`a_table_name`. LIMIT 1 OFFSET 10"
	expectedArgs := []interface{}{11, 12, 13, 14, 15, 16, 17, "this", "is", "a", "list"}
	suite.Run(t, newSuit(query, expectedQueryStr, expectedArgs))

	var countResult int64
	query.Builder().Count(&countResult).Build()
	expectedQueryStr = "SELECT count(1) FROM `a_table_name` WHERE fieldEqual = ? AND (fieldGreaterOrEqual >= ?) AND fieldGreater > ? AND (fieldLessOrEqual <= ?) AND fieldLess < ? AND fieldNotEqual <> ? AND `fieldNot` <> ? AND field_in_list in (?,?,?,?)"
	suite.Run(t, newSuit(query, expectedQueryStr, expectedArgs))
}
