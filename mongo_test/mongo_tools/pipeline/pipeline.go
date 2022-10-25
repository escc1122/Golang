package pipeline

import "go.mongodb.org/mongo-driver/bson"

func GetMatchOpt() *matchOpt {
	return &matchOpt{
		baseOpt{
			conditions: &bson.M{},
		},
	}
}

func GetGroupOpt() *groupOpt {
	return &groupOpt{
		baseOpt{
			conditions: &bson.M{},
		},
	}
}

func GetSortOpt() *sortOpt {
	return &sortOpt{
		baseOpt{
			conditions: &bson.M{},
		},
	}
}

func GetSimpleBsonD(actionKey string, actionValue interface{}) *bson.D {
	return &bson.D{{actionKey, actionValue}}
}

type baseOpt struct {
	conditions *bson.M
}

func (b *baseOpt) setBsonM(conditionKey string, bsonMKey string, value interface{}) {
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

func (b *baseOpt) setPara(conditionKey string, value interface{}) {
	conditions := *b.conditions
	conditions[conditionKey] = value
}

type matchOpt struct {
	//conditions *bson.M
	baseOpt
}

func (m *matchOpt) Gen() bson.D {
	return bson.D{
		{"$match", m.conditions},
	}
}

func (m *matchOpt) In(key string, value ...interface{}) *matchOpt {
	a := bson.A{}
	for _, v := range value {
		a = append(a, v)
	}
	conditions := *m.conditions
	conditions[key] = bson.D{{"$in", &a}}
	return m
}

func (m *matchOpt) Eq(key string, value interface{}) *matchOpt {
	conditions := *m.conditions
	conditions[key] = GetSimpleBsonD("$eq", value)
	return m
}
func (m *matchOpt) Gt(key string, value interface{}) *matchOpt {
	m.setBsonM(key, "$gt", value)
	return m
}
func (m *matchOpt) Gte(key string, value interface{}) *matchOpt {
	m.setBsonM(key, "$gte", value)
	return m
}

func (m *matchOpt) Lt(key string, value interface{}) *matchOpt {
	m.setBsonM(key, "$lt", value)
	return m
}

func (m *matchOpt) Lte(key string, value interface{}) *matchOpt {
	m.setBsonM(key, "$lte", value)
	return m
}

func (m *matchOpt) GteLt(key string, gteValue interface{}, ltValue interface{}) *matchOpt {
	m.Gte(key, gteValue)
	m.Lt(key, ltValue)
	return m
}

type groupOpt struct {
	baseOpt
}

func (g *groupOpt) SetGroupPara(key string, value string) *groupOpt {
	g.setBsonM("_id", key, "$"+value)
	return g
}

func (g *groupOpt) SetSum(sumKey string, value string) *groupOpt {
	conditions := *g.conditions
	conditions[sumKey] = GetSimpleBsonD("$sum", "$"+value)
	return g
}

func (g *groupOpt) Gen() bson.D {
	return bson.D{
		{"$group", g.conditions},
	}
}

type sortOpt struct {
	baseOpt
}

func (s *sortOpt) Gen() bson.D {
	return bson.D{
		{"$sort", s.conditions},
	}
}

func (s *sortOpt) SetSort(para string, sortType int) *sortOpt {
	s.setPara(para, sortType)
	return s
}
