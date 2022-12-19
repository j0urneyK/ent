// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/json/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// TIsNil applies the IsNil predicate on the "t" field.
func TIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldT))
}

// TNotNil applies the NotNil predicate on the "t" field.
func TNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldT))
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldURL))
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldURL))
}

// URLsIsNil applies the IsNil predicate on the "URLs" field.
func URLsIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldURLs))
}

// URLsNotNil applies the NotNil predicate on the "URLs" field.
func URLsNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldURLs))
}

// RawIsNil applies the IsNil predicate on the "raw" field.
func RawIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldRaw))
}

// RawNotNil applies the NotNil predicate on the "raw" field.
func RawNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldRaw))
}

// IntsIsNil applies the IsNil predicate on the "ints" field.
func IntsIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldInts))
}

// IntsNotNil applies the NotNil predicate on the "ints" field.
func IntsNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldInts))
}

// FloatsIsNil applies the IsNil predicate on the "floats" field.
func FloatsIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldFloats))
}

// FloatsNotNil applies the NotNil predicate on the "floats" field.
func FloatsNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldFloats))
}

// StringsIsNil applies the IsNil predicate on the "strings" field.
func StringsIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldStrings))
}

// StringsNotNil applies the NotNil predicate on the "strings" field.
func StringsNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldStrings))
}

// AddrIsNil applies the IsNil predicate on the "addr" field.
func AddrIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAddr))
}

// AddrNotNil applies the NotNil predicate on the "addr" field.
func AddrNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAddr))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
