// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package enttask

import (
	"time"

	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/ent/schema/task"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.EQ(vc))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.EQ(v))
	})
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.EQ(vc))
	})
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.NEQ(vc))
	})
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...task.Priority) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.Within(v...))
	})
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...task.Priority) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.Without(v...))
	})
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.GT(vc))
	})
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.GTE(vc))
	})
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.LT(vc))
	})
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldPriority, p.LTE(vc))
	})
}

// PrioritiesIsNil applies the IsNil predicate on the "priorities" field.
func PrioritiesIsNil() predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldPriorities)
	})
}

// PrioritiesNotNil applies the NotNil predicate on the "priorities" field.
func PrioritiesNotNil() predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldPriorities)
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.EQ(v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.NEQ(v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.Within(vs...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.Without(vs...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.GT(v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.GTE(v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.LT(v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreatedAt, p.LTE(v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(tr *dsl.Traversal) {
		trs := make([]any, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(tr *dsl.Traversal) {
		trs := make([]any, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
