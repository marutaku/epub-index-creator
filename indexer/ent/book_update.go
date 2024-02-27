// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marutaku/epub-index-creator/indexer/ent/book"
	"github.com/marutaku/epub-index-creator/indexer/ent/page"
	"github.com/marutaku/epub-index-creator/indexer/ent/predicate"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where appends a list predicates to the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetIsbn sets the "isbn" field.
func (bu *BookUpdate) SetIsbn(s string) *BookUpdate {
	bu.mutation.SetIsbn(s)
	return bu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bu *BookUpdate) SetNillableIsbn(s *string) *BookUpdate {
	if s != nil {
		bu.SetIsbn(*s)
	}
	return bu
}

// SetTitle sets the "title" field.
func (bu *BookUpdate) SetTitle(s string) *BookUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bu *BookUpdate) SetNillableTitle(s *string) *BookUpdate {
	if s != nil {
		bu.SetTitle(*s)
	}
	return bu
}

// SetLanguage sets the "language" field.
func (bu *BookUpdate) SetLanguage(s string) *BookUpdate {
	bu.mutation.SetLanguage(s)
	return bu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (bu *BookUpdate) SetNillableLanguage(s *string) *BookUpdate {
	if s != nil {
		bu.SetLanguage(*s)
	}
	return bu
}

// SetAuthor sets the "author" field.
func (bu *BookUpdate) SetAuthor(s string) *BookUpdate {
	bu.mutation.SetAuthor(s)
	return bu
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (bu *BookUpdate) SetNillableAuthor(s *string) *BookUpdate {
	if s != nil {
		bu.SetAuthor(*s)
	}
	return bu
}

// SetPublisher sets the "publisher" field.
func (bu *BookUpdate) SetPublisher(s string) *BookUpdate {
	bu.mutation.SetPublisher(s)
	return bu
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePublisher(s *string) *BookUpdate {
	if s != nil {
		bu.SetPublisher(*s)
	}
	return bu
}

// AddPageIDs adds the "pages" edge to the Page entity by IDs.
func (bu *BookUpdate) AddPageIDs(ids ...int) *BookUpdate {
	bu.mutation.AddPageIDs(ids...)
	return bu
}

// AddPages adds the "pages" edges to the Page entity.
func (bu *BookUpdate) AddPages(p ...*Page) *BookUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddPageIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// ClearPages clears all "pages" edges to the Page entity.
func (bu *BookUpdate) ClearPages() *BookUpdate {
	bu.mutation.ClearPages()
	return bu
}

// RemovePageIDs removes the "pages" edge to Page entities by IDs.
func (bu *BookUpdate) RemovePageIDs(ids ...int) *BookUpdate {
	bu.mutation.RemovePageIDs(ids...)
	return bu
}

// RemovePages removes "pages" edges to Page entities.
func (bu *BookUpdate) RemovePages(p ...*Page) *BookUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemovePageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookUpdate) check() error {
	if v, ok := bu.mutation.Isbn(); ok {
		if err := book.IsbnValidator(v); err != nil {
			return &ValidationError{Name: "isbn", err: fmt.Errorf(`ent: validator failed for field "Book.isbn": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Language(); ok {
		if err := book.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "Book.language": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Author(); ok {
		if err := book.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "author", err: fmt.Errorf(`ent: validator failed for field "Book.author": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Publisher(); ok {
		if err := book.PublisherValidator(v); err != nil {
			return &ValidationError{Name: "publisher", err: fmt.Errorf(`ent: validator failed for field "Book.publisher": %w`, err)}
		}
	}
	return nil
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Isbn(); ok {
		_spec.SetField(book.FieldIsbn, field.TypeString, value)
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Language(); ok {
		_spec.SetField(book.FieldLanguage, field.TypeString, value)
	}
	if value, ok := bu.mutation.Author(); ok {
		_spec.SetField(book.FieldAuthor, field.TypeString, value)
	}
	if value, ok := bu.mutation.Publisher(); ok {
		_spec.SetField(book.FieldPublisher, field.TypeString, value)
	}
	if bu.mutation.PagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.PagesTable,
			Columns: []string{book.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedPagesIDs(); len(nodes) > 0 && !bu.mutation.PagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.PagesTable,
			Columns: []string{book.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.PagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.PagesTable,
			Columns: []string{book.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookMutation
}

// SetIsbn sets the "isbn" field.
func (buo *BookUpdateOne) SetIsbn(s string) *BookUpdateOne {
	buo.mutation.SetIsbn(s)
	return buo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableIsbn(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetIsbn(*s)
	}
	return buo
}

// SetTitle sets the "title" field.
func (buo *BookUpdateOne) SetTitle(s string) *BookUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableTitle(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetTitle(*s)
	}
	return buo
}

// SetLanguage sets the "language" field.
func (buo *BookUpdateOne) SetLanguage(s string) *BookUpdateOne {
	buo.mutation.SetLanguage(s)
	return buo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableLanguage(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetLanguage(*s)
	}
	return buo
}

// SetAuthor sets the "author" field.
func (buo *BookUpdateOne) SetAuthor(s string) *BookUpdateOne {
	buo.mutation.SetAuthor(s)
	return buo
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableAuthor(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetAuthor(*s)
	}
	return buo
}

// SetPublisher sets the "publisher" field.
func (buo *BookUpdateOne) SetPublisher(s string) *BookUpdateOne {
	buo.mutation.SetPublisher(s)
	return buo
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePublisher(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetPublisher(*s)
	}
	return buo
}

// AddPageIDs adds the "pages" edge to the Page entity by IDs.
func (buo *BookUpdateOne) AddPageIDs(ids ...int) *BookUpdateOne {
	buo.mutation.AddPageIDs(ids...)
	return buo
}

// AddPages adds the "pages" edges to the Page entity.
func (buo *BookUpdateOne) AddPages(p ...*Page) *BookUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddPageIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// ClearPages clears all "pages" edges to the Page entity.
func (buo *BookUpdateOne) ClearPages() *BookUpdateOne {
	buo.mutation.ClearPages()
	return buo
}

// RemovePageIDs removes the "pages" edge to Page entities by IDs.
func (buo *BookUpdateOne) RemovePageIDs(ids ...int) *BookUpdateOne {
	buo.mutation.RemovePageIDs(ids...)
	return buo
}

// RemovePages removes "pages" edges to Page entities.
func (buo *BookUpdateOne) RemovePages(p ...*Page) *BookUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemovePageIDs(ids...)
}

// Where appends a list predicates to the BookUpdate builder.
func (buo *BookUpdateOne) Where(ps ...predicate.Book) *BookUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookUpdateOne) Select(field string, fields ...string) *BookUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookUpdateOne) check() error {
	if v, ok := buo.mutation.Isbn(); ok {
		if err := book.IsbnValidator(v); err != nil {
			return &ValidationError{Name: "isbn", err: fmt.Errorf(`ent: validator failed for field "Book.isbn": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Language(); ok {
		if err := book.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "Book.language": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Author(); ok {
		if err := book.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "author", err: fmt.Errorf(`ent: validator failed for field "Book.author": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Publisher(); ok {
		if err := book.PublisherValidator(v); err != nil {
			return &ValidationError{Name: "publisher", err: fmt.Errorf(`ent: validator failed for field "Book.publisher": %w`, err)}
		}
	}
	return nil
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Book.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for _, f := range fields {
			if !book.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Isbn(); ok {
		_spec.SetField(book.FieldIsbn, field.TypeString, value)
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Language(); ok {
		_spec.SetField(book.FieldLanguage, field.TypeString, value)
	}
	if value, ok := buo.mutation.Author(); ok {
		_spec.SetField(book.FieldAuthor, field.TypeString, value)
	}
	if value, ok := buo.mutation.Publisher(); ok {
		_spec.SetField(book.FieldPublisher, field.TypeString, value)
	}
	if buo.mutation.PagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.PagesTable,
			Columns: []string{book.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedPagesIDs(); len(nodes) > 0 && !buo.mutation.PagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.PagesTable,
			Columns: []string{book.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.PagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.PagesTable,
			Columns: []string{book.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
