// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marutaku/epub-index-creator/ent/book"
	"github.com/marutaku/epub-index-creator/ent/keyword"
	"github.com/marutaku/epub-index-creator/ent/page"
	"github.com/marutaku/epub-index-creator/ent/predicate"
)

// PageUpdate is the builder for updating Page entities.
type PageUpdate struct {
	config
	hooks    []Hook
	mutation *PageMutation
}

// Where appends a list predicates to the PageUpdate builder.
func (pu *PageUpdate) Where(ps ...predicate.Page) *PageUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PageUpdate) SetTitle(s string) *PageUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *PageUpdate) SetNillableTitle(s *string) *PageUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// SetBookID sets the "book" edge to the Book entity by ID.
func (pu *PageUpdate) SetBookID(id int) *PageUpdate {
	pu.mutation.SetBookID(id)
	return pu
}

// SetNillableBookID sets the "book" edge to the Book entity by ID if the given value is not nil.
func (pu *PageUpdate) SetNillableBookID(id *int) *PageUpdate {
	if id != nil {
		pu = pu.SetBookID(*id)
	}
	return pu
}

// SetBook sets the "book" edge to the Book entity.
func (pu *PageUpdate) SetBook(b *Book) *PageUpdate {
	return pu.SetBookID(b.ID)
}

// AddKeywordIDs adds the "keywords" edge to the Keyword entity by IDs.
func (pu *PageUpdate) AddKeywordIDs(ids ...int) *PageUpdate {
	pu.mutation.AddKeywordIDs(ids...)
	return pu
}

// AddKeywords adds the "keywords" edges to the Keyword entity.
func (pu *PageUpdate) AddKeywords(k ...*Keyword) *PageUpdate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return pu.AddKeywordIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (pu *PageUpdate) Mutation() *PageMutation {
	return pu.mutation
}

// ClearBook clears the "book" edge to the Book entity.
func (pu *PageUpdate) ClearBook() *PageUpdate {
	pu.mutation.ClearBook()
	return pu
}

// ClearKeywords clears all "keywords" edges to the Keyword entity.
func (pu *PageUpdate) ClearKeywords() *PageUpdate {
	pu.mutation.ClearKeywords()
	return pu
}

// RemoveKeywordIDs removes the "keywords" edge to Keyword entities by IDs.
func (pu *PageUpdate) RemoveKeywordIDs(ids ...int) *PageUpdate {
	pu.mutation.RemoveKeywordIDs(ids...)
	return pu
}

// RemoveKeywords removes "keywords" edges to Keyword entities.
func (pu *PageUpdate) RemoveKeywords(k ...*Keyword) *PageUpdate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return pu.RemoveKeywordIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PageUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PageUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PageUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PageUpdate) check() error {
	if v, ok := pu.mutation.Title(); ok {
		if err := page.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Page.title": %w`, err)}
		}
	}
	return nil
}

func (pu *PageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
	}
	if pu.mutation.BookCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BookIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.KeywordsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedKeywordsIDs(); len(nodes) > 0 && !pu.mutation.KeywordsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.KeywordsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PageUpdateOne is the builder for updating a single Page entity.
type PageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PageMutation
}

// SetTitle sets the "title" field.
func (puo *PageUpdateOne) SetTitle(s string) *PageUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableTitle(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// SetBookID sets the "book" edge to the Book entity by ID.
func (puo *PageUpdateOne) SetBookID(id int) *PageUpdateOne {
	puo.mutation.SetBookID(id)
	return puo
}

// SetNillableBookID sets the "book" edge to the Book entity by ID if the given value is not nil.
func (puo *PageUpdateOne) SetNillableBookID(id *int) *PageUpdateOne {
	if id != nil {
		puo = puo.SetBookID(*id)
	}
	return puo
}

// SetBook sets the "book" edge to the Book entity.
func (puo *PageUpdateOne) SetBook(b *Book) *PageUpdateOne {
	return puo.SetBookID(b.ID)
}

// AddKeywordIDs adds the "keywords" edge to the Keyword entity by IDs.
func (puo *PageUpdateOne) AddKeywordIDs(ids ...int) *PageUpdateOne {
	puo.mutation.AddKeywordIDs(ids...)
	return puo
}

// AddKeywords adds the "keywords" edges to the Keyword entity.
func (puo *PageUpdateOne) AddKeywords(k ...*Keyword) *PageUpdateOne {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return puo.AddKeywordIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (puo *PageUpdateOne) Mutation() *PageMutation {
	return puo.mutation
}

// ClearBook clears the "book" edge to the Book entity.
func (puo *PageUpdateOne) ClearBook() *PageUpdateOne {
	puo.mutation.ClearBook()
	return puo
}

// ClearKeywords clears all "keywords" edges to the Keyword entity.
func (puo *PageUpdateOne) ClearKeywords() *PageUpdateOne {
	puo.mutation.ClearKeywords()
	return puo
}

// RemoveKeywordIDs removes the "keywords" edge to Keyword entities by IDs.
func (puo *PageUpdateOne) RemoveKeywordIDs(ids ...int) *PageUpdateOne {
	puo.mutation.RemoveKeywordIDs(ids...)
	return puo
}

// RemoveKeywords removes "keywords" edges to Keyword entities.
func (puo *PageUpdateOne) RemoveKeywords(k ...*Keyword) *PageUpdateOne {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return puo.RemoveKeywordIDs(ids...)
}

// Where appends a list predicates to the PageUpdate builder.
func (puo *PageUpdateOne) Where(ps ...predicate.Page) *PageUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PageUpdateOne) Select(field string, fields ...string) *PageUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Page entity.
func (puo *PageUpdateOne) Save(ctx context.Context) (*Page, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PageUpdateOne) SaveX(ctx context.Context) *Page {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PageUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PageUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PageUpdateOne) check() error {
	if v, ok := puo.mutation.Title(); ok {
		if err := page.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Page.title": %w`, err)}
		}
	}
	return nil
}

func (puo *PageUpdateOne) sqlSave(ctx context.Context) (_node *Page, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Page.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, page.FieldID)
		for _, f := range fields {
			if !page.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != page.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
	}
	if puo.mutation.BookCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BookIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.KeywordsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedKeywordsIDs(); len(nodes) > 0 && !puo.mutation.KeywordsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.KeywordsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Page{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
