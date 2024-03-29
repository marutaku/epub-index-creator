// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marutaku/epub-index-creator/ent/book"
	"github.com/marutaku/epub-index-creator/ent/keyword"
	"github.com/marutaku/epub-index-creator/ent/page"
)

// PageCreate is the builder for creating a Page entity.
type PageCreate struct {
	config
	mutation *PageMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (pc *PageCreate) SetTitle(s string) *PageCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetPath sets the "path" field.
func (pc *PageCreate) SetPath(s string) *PageCreate {
	pc.mutation.SetPath(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PageCreate) SetID(i int) *PageCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetBookID sets the "book" edge to the Book entity by ID.
func (pc *PageCreate) SetBookID(id int) *PageCreate {
	pc.mutation.SetBookID(id)
	return pc
}

// SetNillableBookID sets the "book" edge to the Book entity by ID if the given value is not nil.
func (pc *PageCreate) SetNillableBookID(id *int) *PageCreate {
	if id != nil {
		pc = pc.SetBookID(*id)
	}
	return pc
}

// SetBook sets the "book" edge to the Book entity.
func (pc *PageCreate) SetBook(b *Book) *PageCreate {
	return pc.SetBookID(b.ID)
}

// AddKeywordIDs adds the "keywords" edge to the Keyword entity by IDs.
func (pc *PageCreate) AddKeywordIDs(ids ...int) *PageCreate {
	pc.mutation.AddKeywordIDs(ids...)
	return pc
}

// AddKeywords adds the "keywords" edges to the Keyword entity.
func (pc *PageCreate) AddKeywords(k ...*Keyword) *PageCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return pc.AddKeywordIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (pc *PageCreate) Mutation() *PageMutation {
	return pc.mutation
}

// Save creates the Page in the database.
func (pc *PageCreate) Save(ctx context.Context) (*Page, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PageCreate) SaveX(ctx context.Context) *Page {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PageCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PageCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PageCreate) check() error {
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Page.title"`)}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := page.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Page.title": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Page.path"`)}
	}
	if v, ok := pc.mutation.Path(); ok {
		if err := page.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Page.path": %w`, err)}
		}
	}
	return nil
}

func (pc *PageCreate) sqlSave(ctx context.Context) (*Page, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PageCreate) createSpec() (*Page, *sqlgraph.CreateSpec) {
	var (
		_node = &Page{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(page.Table, sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Path(); ok {
		_spec.SetField(page.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if nodes := pc.mutation.BookIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.BookTable,
			Columns: []string{page.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.book_pages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.KeywordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.KeywordsTable,
			Columns: []string{page.KeywordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(keyword.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PageCreateBulk is the builder for creating many Page entities in bulk.
type PageCreateBulk struct {
	config
	err      error
	builders []*PageCreate
}

// Save creates the Page entities in the database.
func (pcb *PageCreateBulk) Save(ctx context.Context) ([]*Page, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Page, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PageMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PageCreateBulk) SaveX(ctx context.Context) []*Page {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PageCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PageCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
