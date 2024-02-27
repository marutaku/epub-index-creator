// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marutaku/epub-index-creator/indexer/ent/book"
	"github.com/marutaku/epub-index-creator/indexer/ent/keyword"
	"github.com/marutaku/epub-index-creator/indexer/ent/predicate"
)

// BookQuery is the builder for querying Book entities.
type BookQuery struct {
	config
	ctx        *QueryContext
	order      []book.OrderOption
	inters     []Interceptor
	predicates []predicate.Book
	withCars   *KeywordQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookQuery builder.
func (bq *BookQuery) Where(ps ...predicate.Book) *BookQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BookQuery) Limit(limit int) *BookQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BookQuery) Offset(offset int) *BookQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BookQuery) Unique(unique bool) *BookQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BookQuery) Order(o ...book.OrderOption) *BookQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryCars chains the current query on the "cars" edge.
func (bq *BookQuery) QueryCars() *KeywordQuery {
	query := (&KeywordClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, selector),
			sqlgraph.To(keyword.Table, keyword.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, book.CarsTable, book.CarsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Book entity from the query.
// Returns a *NotFoundError when no Book was found.
func (bq *BookQuery) First(ctx context.Context) (*Book, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{book.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BookQuery) FirstX(ctx context.Context) *Book {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Book ID from the query.
// Returns a *NotFoundError when no Book ID was found.
func (bq *BookQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{book.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BookQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Book entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Book entity is found.
// Returns a *NotFoundError when no Book entities are found.
func (bq *BookQuery) Only(ctx context.Context) (*Book, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{book.Label}
	default:
		return nil, &NotSingularError{book.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BookQuery) OnlyX(ctx context.Context) *Book {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Book ID in the query.
// Returns a *NotSingularError when more than one Book ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BookQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = &NotSingularError{book.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BookQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Books.
func (bq *BookQuery) All(ctx context.Context) ([]*Book, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Book, *BookQuery]()
	return withInterceptors[[]*Book](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BookQuery) AllX(ctx context.Context) []*Book {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Book IDs.
func (bq *BookQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(book.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BookQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BookQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BookQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BookQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BookQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BookQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BookQuery) Clone() *BookQuery {
	if bq == nil {
		return nil
	}
	return &BookQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]book.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Book{}, bq.predicates...),
		withCars:   bq.withCars.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithCars tells the query-builder to eager-load the nodes that are connected to
// the "cars" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookQuery) WithCars(opts ...func(*KeywordQuery)) *BookQuery {
	query := (&KeywordClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withCars = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Isbn string `json:"isbn,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Book.Query().
//		GroupBy(book.FieldIsbn).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BookQuery) GroupBy(field string, fields ...string) *BookGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BookGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = book.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Isbn string `json:"isbn,omitempty"`
//	}
//
//	client.Book.Query().
//		Select(book.FieldIsbn).
//		Scan(ctx, &v)
func (bq *BookQuery) Select(fields ...string) *BookSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BookSelect{BookQuery: bq}
	sbuild.label = book.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BookSelect configured with the given aggregations.
func (bq *BookQuery) Aggregate(fns ...AggregateFunc) *BookSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BookQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !book.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BookQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Book, error) {
	var (
		nodes       = []*Book{}
		_spec       = bq.querySpec()
		loadedTypes = [1]bool{
			bq.withCars != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Book).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Book{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withCars; query != nil {
		if err := bq.loadCars(ctx, query, nodes,
			func(n *Book) { n.Edges.Cars = []*Keyword{} },
			func(n *Book, e *Keyword) { n.Edges.Cars = append(n.Edges.Cars, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BookQuery) loadCars(ctx context.Context, query *KeywordQuery, nodes []*Book, init func(*Book), assign func(*Book, *Keyword)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Book)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Keyword(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(book.CarsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.book_cars
		if fk == nil {
			return fmt.Errorf(`foreign-key "book_cars" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "book_cars" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BookQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BookQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for i := range fields {
			if fields[i] != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BookQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(book.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = book.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookGroupBy is the group-by builder for Book entities.
type BookGroupBy struct {
	selector
	build *BookQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BookGroupBy) Aggregate(fns ...AggregateFunc) *BookGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BookGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookQuery, *BookGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BookGroupBy) sqlScan(ctx context.Context, root *BookQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BookSelect is the builder for selecting fields of Book entities.
type BookSelect struct {
	*BookQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BookSelect) Aggregate(fns ...AggregateFunc) *BookSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BookSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookQuery, *BookSelect](ctx, bs.BookQuery, bs, bs.inters, v)
}

func (bs *BookSelect) sqlScan(ctx context.Context, root *BookQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}