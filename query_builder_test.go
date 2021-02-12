package dbquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type modelStruct struct{}

type findStruct struct{}

var interfaceSlice = make([]interface{}, 0)

var countResult int64

const offset = 3
const limit = 4

var inList = []interface{}{
	"this", "is", "a", "list",
}

const valueEqual = 10
const valueGreaterOrEqual = 11
const valueGreater = 12
const valueLessOrEqual = "thirteen"
const valueLess = "fourteen"
const valueNotEqual = 15

func TestQuery_Builder(t *testing.T) {
	builder := NewBuilder()
	builder = buildModel(builder)
	builder = buildPreload(builder)
	builder = buildFind(builder)
	builder = buildCount(builder)
	builder = buildSort(builder)
	builder = buildOffset(builder)
	builder = buildLimit(builder)
	builder = buildInList(builder)
	builder = buildMatches(builder)

	query := builder.Build()
	testModel(t, query)
	testPreload(t, query)
	testFind(t, query)
	testCount(t, query)
	testSort(t, query)
	testOffset(t, query)
	testLimit(t, query)
	testInList(t, query)
	testMatches(t, query)
}

// region Model
func buildModel(builder *Builder) *Builder {
	return builder.Model(&modelStruct{})
}

func testModel(t *testing.T, q Query) {
	qt := typeModel
	verifyAvailability(t, q, qt, 1)
	verifyItemType(t, q.data[qt], &modelStruct{})
}

//endregion

// region Preload
func buildPreload(builder *Builder) *Builder {
	return builder.
		Preload("preloadA").
		Preload("preloadB", "conditionB1").
		Preload("preloadC", "conditionC1", "conditionC2")
}

func testPreload(t *testing.T, q Query) {
	qt := typePreload
	verifyAvailability(t, q, qt, 3)
	verifyItemType(t, q.data[qt], &QueryKeyValue{})

	items := q.data[typePreload]

	preloadA := items[0].(*QueryKeyValue)
	assert.Equal(t, "preloadA", preloadA.key)
	assert.Nil(t, preloadA.value)

	preloadB := items[1].(*QueryKeyValue)
	assert.Equal(t, "preloadB", preloadB.key)
	assert.NotNil(t, preloadB.value)
	assert.IsType(t, interfaceSlice, preloadB.value)
	assert.Len(t, preloadB.value, 1)
	assert.Equal(t, "conditionB1", preloadB.value.([]interface{})[0])

	preloadC := items[2].(*QueryKeyValue)
	assert.Equal(t, "preloadC", preloadC.key)
	assert.NotNil(t, preloadC.value)
	assert.IsType(t, interfaceSlice, preloadC.value)
	assert.Len(t, preloadC.value, 2)
	assert.Equal(t, "conditionC1", preloadC.value.([]interface{})[0])
	assert.Equal(t, "conditionC2", preloadC.value.([]interface{})[1])
}

// endregion

// region Find
func buildFind(builder *Builder) *Builder {
	return builder.Find(&findStruct{})
}

func testFind(t *testing.T, q Query) {
	qt := typeFind
	verifyAvailability(t, q, qt, 1)
	verifyItemType(t, q.data[qt], &findStruct{})
}

// endregion

// region Count
func buildCount(builder *Builder) *Builder {
	return builder.Count(&countResult)
}

func testCount(t *testing.T, q Query) {
	qt := typeCount
	verifyAvailability(t, q, qt, 1)
	verifyItemType(t, q.data[qt], &countResult)

	*(q.data[qt][0].(*int64)) = 5
	assert.Equal(t, int64(5), countResult)
}

// endregion

// region Sort
func buildSort(builder *Builder) *Builder {
	return builder.Sort("sortedField", OrderDesc)
}

func testSort(t *testing.T, q Query) {
	qt := typeSort
	verifyAvailability(t, q, qt, 1)
	verifyItemType(t, q.data[qt], &QueryKeyValue{})

	kv := q.data[qt][0].(*QueryKeyValue)
	assert.Equal(t, "sortedField", kv.key)
	assert.Equal(t, OrderDesc, kv.value)
}

// endregion

// region Offset
func buildOffset(builder *Builder) *Builder {
	return builder.Offset(offset)
}

func testOffset(t *testing.T, q Query) {
	qt := typeOffset
	verifyAvailability(t, q, qt, 1)
	verifyItemType(t, q.data[qt], offset)
	assert.Equal(t, offset, q.data[qt][0].(int))
}

// endregion

// region Limit
func buildLimit(builder *Builder) *Builder {
	return builder.Limit(limit)
}

func testLimit(t *testing.T, q Query) {
	qt := typeLimit
	verifyAvailability(t, q, qt, 1)
	verifyItemType(t, q.data[qt], limit)
	assert.Equal(t, limit, q.data[qt][0].(int))
}

// endregion

// region In List
func buildInList(builder *Builder) *Builder {
	return builder.InList("fieldInList", inList)
}

func testInList(t *testing.T, q Query) {
	qt := typeInList
	verifyAvailability(t, q, qt, 1)
	verifyItemType(t, q.data[qt], &QueryKeyValue{})

	kv := q.data[qt][0].(*QueryKeyValue)
	assert.Equal(t, "fieldInList", kv.key)
	toCheckList := kv.value.([]interface{})
	assert.Len(t, toCheckList, 4)
	assert.Equal(t, toCheckList[0], "this")
	assert.Equal(t, toCheckList[1], "is")
	assert.Equal(t, toCheckList[2], "a")
	assert.Equal(t, toCheckList[3], "list")
}

// endregion

// region Matches
func buildMatches(builder *Builder) *Builder {
	return builder.
		Equal("fieldEqual", valueEqual).
		GreaterOrEqual("fieldGreaterOrEqual", valueGreaterOrEqual).
		Greater("fieldGreater", valueGreater).
		LessOrEqual("fieldLessOrEqual", valueLessOrEqual).
		Less("fieldLess", valueLess).
		NotEqual("fieldNotEqual", valueNotEqual)
}

func testMatches(t *testing.T, q Query) {
	qt := typeMatch
	verifyAvailability(t, q, qt, 6)
	verifyItemType(t, q.data[qt], &ComparingObject{})

	testMatch(t, q.data[qt][0], "fieldEqual", operatorEqual, valueEqual)
	testMatch(t, q.data[qt][1], "fieldGreaterOrEqual", operatorGreaterOrEqual, valueGreaterOrEqual)
	testMatch(t, q.data[qt][2], "fieldGreater", operatorGreater, valueGreater)
	testMatch(t, q.data[qt][3], "fieldLessOrEqual", operatorLessOrEqual, valueLessOrEqual)
	testMatch(t, q.data[qt][4], "fieldLess", operatorLess, valueLess)
	testMatch(t, q.data[qt][5], "fieldNotEqual", operatorNotEqual, valueNotEqual)
}

func testMatch(t *testing.T, data interface{}, expectedFieldName string, expectedOpe operator, expectedValue interface{}) {
	cprObj := data.(*ComparingObject)
	assert.Equal(t, expectedFieldName, cprObj.key)
	assert.Equal(t, expectedOpe, cprObj.ope)
	assert.Equal(t, expectedValue, cprObj.value)
}

// endregion

func verifyAvailability(t *testing.T, q Query, qType queryType, length int) {
	assert.NotNil(t, q.data[qType], "%s empty", qType)
	assert.IsType(t, interfaceSlice, q.data[qType], "%s wrong type", qType)
	assert.Len(t, q.data[qType], length, "%s wrong length", qType)
}

func verifyItemType(t *testing.T, data interface{}, expectedType interface{}) {
	items := data.([]interface{})
	for _, item := range items {
		assert.IsType(t, expectedType, item)
	}
}
