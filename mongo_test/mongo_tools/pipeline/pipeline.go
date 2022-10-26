package pipeline

import "go.mongodb.org/mongo-driver/bson"

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

func GetSimpleBsonD(actionKey string, actionValue interface{}) *bson.D {
	return &bson.D{{actionKey, actionValue}}
}

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
	conditions[column] = GetSimpleBsonD("$eq", value)
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

func (g *groupGenerate) Sum(aliases string, sumPara string) *groupGenerate {
	conditions := *g.conditions
	conditions[aliases] = GetSimpleBsonD("$sum", "$"+sumPara)
	return g
}

func (g *groupGenerate) Count(aliases string) *groupGenerate {
	conditions := *g.conditions
	conditions[aliases] = GetSimpleBsonD("$sum", 1)
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
