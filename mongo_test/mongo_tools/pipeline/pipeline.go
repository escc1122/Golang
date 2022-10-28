package pipeline

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

const (
	ASC  = 1
	DESC = -1
)

func GetMatchGenerate() *matchGenerate {
	return &matchGenerate{
		baseGenerate{
			conditions: &bson.M{},
		},
	}
}

func GetGroupGenerate() *groupGenerate {
	return &groupGenerate{
		baseGenerate{
			conditions: &bson.M{},
		},
	}
}

func GetSortGenerate() *sortGenerate {
	return &sortGenerate{
		baseGenerate{
			conditions: &bson.M{},
		},
	}
}

func GetFacetGenerate() *facetGenerate {
	return &facetGenerate{
		baseGenerate{
			conditions: &bson.M{},
		},
	}
}

//func GetSimpleBsonD(actionKey string, actionValue interface{}) *bson.D {
//	return &bson.D{{actionKey, actionValue}}
//}

func GetLimitBsonD(limit int) bson.D {
	return bson.D{{"$limit", limit}}
}

func GetSkipBsonD(skip int) bson.D {
	return bson.D{{"$skip", skip}}
}

type baseGenerate struct {
	conditions *bson.M
}

func (b *baseGenerate) setBsonM(conditionKey string, bsonMKey string, value interface{}) {
	conditions := *b.conditions
	var bsonM bson.M
	if conditions[conditionKey] == nil {
		bsonM = bson.M{}
		conditions[conditionKey] = &bsonM
	} else {
		bsonM = *conditions[conditionKey].(*bson.M)
	}
	bsonM[bsonMKey] = value
}

func (b *baseGenerate) setValue(conditionKey string, value interface{}) {
	conditions := *b.conditions
	conditions[conditionKey] = value
}

func (b *baseGenerate) genBsonD(actionKey string, actionValue interface{}) *bson.D {
	return &bson.D{{actionKey, actionValue}}
}

type matchGenerate struct {
	baseGenerate
}

func (m *matchGenerate) GenBsonD() bson.D {
	return bson.D{
		{"$match", m.conditions},
	}
}

func (m *matchGenerate) In(column string, value ...interface{}) *matchGenerate {
	a := bson.A{}
	for _, v := range value {
		a = append(a, v)
	}
	conditions := *m.conditions
	conditions[column] = bson.D{{"$in", &a}}
	return m
}

func (m *matchGenerate) Eq(column string, value interface{}) *matchGenerate {
	conditions := *m.conditions
	conditions[column] = m.genBsonD("$eq", value)
	return m
}
func (m *matchGenerate) Gt(column string, value interface{}) *matchGenerate {
	m.setBsonM(column, "$gt", value)
	return m
}
func (m *matchGenerate) Gte(column string, value interface{}) *matchGenerate {
	m.setBsonM(column, "$gte", value)
	return m
}

func (m *matchGenerate) Lt(column string, value interface{}) *matchGenerate {
	m.setBsonM(column, "$lt", value)
	return m
}

func (m *matchGenerate) Lte(column string, value interface{}) *matchGenerate {
	m.setBsonM(column, "$lte", value)
	return m
}

func (m *matchGenerate) GteLt(column string, gteValue interface{}, ltValue interface{}) *matchGenerate {
	m.Gte(column, gteValue)
	m.Lt(column, ltValue)
	return m
}

type groupGenerate struct {
	baseGenerate
}

func (g *groupGenerate) GroupBy(aliases string, column string) *groupGenerate {
	g.setBsonM("_id", aliases, "$"+column)
	return g
}

// GroupByAll https://www.mongodb.com/docs/manual/reference/operator/aggregation/group/#std-label-null-example
// Group by null
// The following aggregation operation specifies a group _id of null, calculating the total sale amount, average quantity, and count of all documents in the collection.
func (g *groupGenerate) GroupByAll() *groupGenerate {
	//g.setBsonM("_id", "all", nil)
	conditions := *g.conditions
	conditions["_id"] = nil
	return g
}

func (g *groupGenerate) Sum(aliases string, sumPara string) *groupGenerate {
	conditions := *g.conditions
	conditions[aliases] = g.genBsonD("$sum", "$"+sumPara)
	return g
}

func (g *groupGenerate) Count(aliases string) *groupGenerate {
	conditions := *g.conditions
	conditions[aliases] = g.genBsonD("$sum", 1)
	return g
}

// ShowData https://www.mongodb.com/docs/manual/reference/operator/aggregation/push/#mongodb-group-grp.-push
func (g *groupGenerate) ShowData(strut interface{}) *groupGenerate {
	m := bson.M{}
	val := reflect.ValueOf(strut).Elem()
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag
		key := tag.Get("bson")
		if key == "" {
			key = tag.Get("json")
		}
		if key == "" {
			key = val.Type().Field(i).Name
		}
		m[key] = "$" + key
	}
	g.setBsonM("data", "$push", m)
	return g
}

func (g *groupGenerate) GenBsonD() bson.D {
	return bson.D{
		{"$group", g.conditions},
	}
}

type sortGenerate struct {
	baseGenerate
}

func (s *sortGenerate) GenBsonD() bson.D {
	return bson.D{
		{"$sort", s.conditions},
	}
}

func (s *sortGenerate) Sort(column string, sortType int) *sortGenerate {
	s.setValue(column, sortType)
	return s
}

// facetGenerate https://www.mongodb.com/docs/v5.0/reference/operator/aggregation/facet/
type facetGenerate struct {
	baseGenerate
}

func (f *facetGenerate) AppendPipeline(aliases string, pipeline mongo.Pipeline) *facetGenerate {
	f.setValue(aliases, pipeline)
	return f
}
func (f *facetGenerate) GenBsonD() bson.D {
	return bson.D{
		{"$facet", f.conditions},
	}
}
