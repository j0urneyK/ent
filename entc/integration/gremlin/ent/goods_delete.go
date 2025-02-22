// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/goods"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// GoodsDelete is the builder for deleting a Goods entity.
type GoodsDelete struct {
	config
	hooks    []Hook
	mutation *GoodsMutation
}

// Where appends a list predicates to the GoodsDelete builder.
func (gd *GoodsDelete) Where(ps ...predicate.Goods) *GoodsDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GoodsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GoodsMutation](ctx, gd.gremlinExec, gd.mutation, gd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GoodsDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GoodsDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gd.gremlin().Query()
	if err := gd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	gd.mutation.done = true
	return res.ReadInt()
}

func (gd *GoodsDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(goods.Label)
	for _, p := range gd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// GoodsDeleteOne is the builder for deleting a single Goods entity.
type GoodsDeleteOne struct {
	gd *GoodsDelete
}

// Exec executes the deletion query.
func (gdo *GoodsDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goods.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GoodsDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}
