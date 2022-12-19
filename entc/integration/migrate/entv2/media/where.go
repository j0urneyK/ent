// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package media

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldID, id))
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldSource, v))
}

// SourceURI applies equality check predicate on the "source_uri" field. It's identical to SourceURIEQ.
func SourceURI(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldSourceURI, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldText, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldSource, vs...))
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldSource, v))
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldSource, v))
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldSource, v))
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldSource, v))
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldSource, v))
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldSource, v))
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldSource, v))
}

// SourceIsNil applies the IsNil predicate on the "source" field.
func SourceIsNil() predicate.Media {
	return predicate.Media(sql.FieldIsNull(FieldSource))
}

// SourceNotNil applies the NotNil predicate on the "source" field.
func SourceNotNil() predicate.Media {
	return predicate.Media(sql.FieldNotNull(FieldSource))
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldSource, v))
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldSource, v))
}

// SourceURIEQ applies the EQ predicate on the "source_uri" field.
func SourceURIEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldSourceURI, v))
}

// SourceURINEQ applies the NEQ predicate on the "source_uri" field.
func SourceURINEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldSourceURI, v))
}

// SourceURIIn applies the In predicate on the "source_uri" field.
func SourceURIIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldSourceURI, vs...))
}

// SourceURINotIn applies the NotIn predicate on the "source_uri" field.
func SourceURINotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldSourceURI, vs...))
}

// SourceURIGT applies the GT predicate on the "source_uri" field.
func SourceURIGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldSourceURI, v))
}

// SourceURIGTE applies the GTE predicate on the "source_uri" field.
func SourceURIGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldSourceURI, v))
}

// SourceURILT applies the LT predicate on the "source_uri" field.
func SourceURILT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldSourceURI, v))
}

// SourceURILTE applies the LTE predicate on the "source_uri" field.
func SourceURILTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldSourceURI, v))
}

// SourceURIContains applies the Contains predicate on the "source_uri" field.
func SourceURIContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldSourceURI, v))
}

// SourceURIHasPrefix applies the HasPrefix predicate on the "source_uri" field.
func SourceURIHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldSourceURI, v))
}

// SourceURIHasSuffix applies the HasSuffix predicate on the "source_uri" field.
func SourceURIHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldSourceURI, v))
}

// SourceURIIsNil applies the IsNil predicate on the "source_uri" field.
func SourceURIIsNil() predicate.Media {
	return predicate.Media(sql.FieldIsNull(FieldSourceURI))
}

// SourceURINotNil applies the NotNil predicate on the "source_uri" field.
func SourceURINotNil() predicate.Media {
	return predicate.Media(sql.FieldNotNull(FieldSourceURI))
}

// SourceURIEqualFold applies the EqualFold predicate on the "source_uri" field.
func SourceURIEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldSourceURI, v))
}

// SourceURIContainsFold applies the ContainsFold predicate on the "source_uri" field.
func SourceURIContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldSourceURI, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldText, v))
}

// TextIsNil applies the IsNil predicate on the "text" field.
func TextIsNil() predicate.Media {
	return predicate.Media(sql.FieldIsNull(FieldText))
}

// TextNotNil applies the NotNil predicate on the "text" field.
func TextNotNil() predicate.Media {
	return predicate.Media(sql.FieldNotNull(FieldText))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldText, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
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
func Not(p predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		p(s.Not())
	})
}
