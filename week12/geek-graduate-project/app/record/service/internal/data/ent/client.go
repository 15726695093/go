// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"clock-in/app/record/service/internal/data/ent/migrate"

	"clock-in/app/record/service/internal/data/ent/record"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Record is the client for interacting with the Record builders.
	Record *RecordClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Record = NewRecordClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Record: NewRecordClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Record: NewRecordClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Record.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Record.Use(hooks...)
}

// RecordClient is a client for the Record schema.
type RecordClient struct {
	config
}

// NewRecordClient returns a client for the Record from the given config.
func NewRecordClient(c config) *RecordClient {
	return &RecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `record.Hooks(f(g(h())))`.
func (c *RecordClient) Use(hooks ...Hook) {
	c.hooks.Record = append(c.hooks.Record, hooks...)
}

// Create returns a create builder for Record.
func (c *RecordClient) Create() *RecordCreate {
	mutation := newRecordMutation(c.config, OpCreate)
	return &RecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Record entities.
func (c *RecordClient) CreateBulk(builders ...*RecordCreate) *RecordCreateBulk {
	return &RecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Record.
func (c *RecordClient) Update() *RecordUpdate {
	mutation := newRecordMutation(c.config, OpUpdate)
	return &RecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RecordClient) UpdateOne(r *Record) *RecordUpdateOne {
	mutation := newRecordMutation(c.config, OpUpdateOne, withRecord(r))
	return &RecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RecordClient) UpdateOneID(id int) *RecordUpdateOne {
	mutation := newRecordMutation(c.config, OpUpdateOne, withRecordID(id))
	return &RecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Record.
func (c *RecordClient) Delete() *RecordDelete {
	mutation := newRecordMutation(c.config, OpDelete)
	return &RecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RecordClient) DeleteOne(r *Record) *RecordDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RecordClient) DeleteOneID(id int) *RecordDeleteOne {
	builder := c.Delete().Where(record.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RecordDeleteOne{builder}
}

// Query returns a query builder for Record.
func (c *RecordClient) Query() *RecordQuery {
	return &RecordQuery{
		config: c.config,
	}
}

// Get returns a Record entity by its id.
func (c *RecordClient) Get(ctx context.Context, id int) (*Record, error) {
	return c.Query().Where(record.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RecordClient) GetX(ctx context.Context, id int) *Record {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RecordClient) Hooks() []Hook {
	return c.hooks.Record
}