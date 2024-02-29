// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marutaku/epub-index-creator/indexer/ent/keyword"
	"github.com/marutaku/epub-index-creator/indexer/ent/page"
)

// KeywordCreate is the builder for creating a Keyword entity.
type KeywordCreate struct {
	config
	mutation *KeywordMutation
	hooks    []Hook
}

// SetKeyword sets the "keyword" field.
func (kc *KeywordCreate) SetKeyword(s string) *KeywordCreate {
	kc.mutation.SetKeyword(s)
	return kc
}

// SetPageID sets the "page" edge to the Page entity by ID.
func (kc *KeywordCreate) SetPageID(id int) *KeywordCreate {
	kc.mutation.SetPageID(id)
	return kc
}

// SetNillablePageID sets the "page" edge to the Page entity by ID if the given value is not nil.
func (kc *KeywordCreate) SetNillablePageID(id *int) *KeywordCreate {
	if id != nil {
		kc = kc.SetPageID(*id)
	}
	return kc
}

// SetPage sets the "page" edge to the Page entity.
func (kc *KeywordCreate) SetPage(p *Page) *KeywordCreate {
	return kc.SetPageID(p.ID)
}

// Mutation returns the KeywordMutation object of the builder.
func (kc *KeywordCreate) Mutation() *KeywordMutation {
	return kc.mutation
}

// Save creates the Keyword in the database.
func (kc *KeywordCreate) Save(ctx context.Context) (*Keyword, error) {
	return withHooks(ctx, kc.sqlSave, kc.mutation, kc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KeywordCreate) SaveX(ctx context.Context) *Keyword {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KeywordCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KeywordCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KeywordCreate) check() error {
	if _, ok := kc.mutation.Keyword(); !ok {
		return &ValidationError{Name: "keyword", err: errors.New(`ent: missing required field "Keyword.keyword"`)}
	}
	if v, ok := kc.mutation.Keyword(); ok {
		if err := keyword.KeywordValidator(v); err != nil {
			return &ValidationError{Name: "keyword", err: fmt.Errorf(`ent: validator failed for field "Keyword.keyword": %w`, err)}
		}
	}
	return nil
}

func (kc *KeywordCreate) sqlSave(ctx context.Context) (*Keyword, error) {
	if err := kc.check(); err != nil {
		return nil, err
	}
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	kc.mutation.id = &_node.ID
	kc.mutation.done = true
	return _node, nil
}

func (kc *KeywordCreate) createSpec() (*Keyword, *sqlgraph.CreateSpec) {
	var (
		_node = &Keyword{config: kc.config}
		_spec = sqlgraph.NewCreateSpec(keyword.Table, sqlgraph.NewFieldSpec(keyword.FieldID, field.TypeInt))
	)
	if value, ok := kc.mutation.Keyword(); ok {
		_spec.SetField(keyword.FieldKeyword, field.TypeString, value)
		_node.Keyword = value
	}
	if nodes := kc.mutation.PageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   keyword.PageTable,
			Columns: []string{keyword.PageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.page_keywords = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KeywordCreateBulk is the builder for creating many Keyword entities in bulk.
type KeywordCreateBulk struct {
	config
	err      error
	builders []*KeywordCreate
}

// Save creates the Keyword entities in the database.
func (kcb *KeywordCreateBulk) Save(ctx context.Context) ([]*Keyword, error) {
	if kcb.err != nil {
		return nil, kcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Keyword, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KeywordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KeywordCreateBulk) SaveX(ctx context.Context) []*Keyword {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KeywordCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KeywordCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}
