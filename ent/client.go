// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/leorcvargas/bgeraser/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/leorcvargas/bgeraser/ent/image"
	"github.com/leorcvargas/bgeraser/ent/imageprocess"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Image is the client for interacting with the Image builders.
	Image *ImageClient
	// ImageProcess is the client for interacting with the ImageProcess builders.
	ImageProcess *ImageProcessClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Image = NewImageClient(c.config)
	c.ImageProcess = NewImageProcessClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Image:        NewImageClient(cfg),
		ImageProcess: NewImageProcessClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:          ctx,
		config:       cfg,
		Image:        NewImageClient(cfg),
		ImageProcess: NewImageProcessClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Image.
//		Query().
//		Count(ctx)
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
	c.Image.Use(hooks...)
	c.ImageProcess.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Image.Intercept(interceptors...)
	c.ImageProcess.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ImageMutation:
		return c.Image.mutate(ctx, m)
	case *ImageProcessMutation:
		return c.ImageProcess.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ImageClient is a client for the Image schema.
type ImageClient struct {
	config
}

// NewImageClient returns a client for the Image from the given config.
func NewImageClient(c config) *ImageClient {
	return &ImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `image.Hooks(f(g(h())))`.
func (c *ImageClient) Use(hooks ...Hook) {
	c.hooks.Image = append(c.hooks.Image, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `image.Intercept(f(g(h())))`.
func (c *ImageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Image = append(c.inters.Image, interceptors...)
}

// Create returns a builder for creating a Image entity.
func (c *ImageClient) Create() *ImageCreate {
	mutation := newImageMutation(c.config, OpCreate)
	return &ImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Image entities.
func (c *ImageClient) CreateBulk(builders ...*ImageCreate) *ImageCreateBulk {
	return &ImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Image.
func (c *ImageClient) Update() *ImageUpdate {
	mutation := newImageMutation(c.config, OpUpdate)
	return &ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageClient) UpdateOne(i *Image) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImage(i))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageClient) UpdateOneID(id uuid.UUID) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImageID(id))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Image.
func (c *ImageClient) Delete() *ImageDelete {
	mutation := newImageMutation(c.config, OpDelete)
	return &ImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ImageClient) DeleteOne(i *Image) *ImageDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ImageClient) DeleteOneID(id uuid.UUID) *ImageDeleteOne {
	builder := c.Delete().Where(image.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageDeleteOne{builder}
}

// Query returns a query builder for Image.
func (c *ImageClient) Query() *ImageQuery {
	return &ImageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeImage},
		inters: c.Interceptors(),
	}
}

// Get returns a Image entity by its id.
func (c *ImageClient) Get(ctx context.Context, id uuid.UUID) (*Image, error) {
	return c.Query().Where(image.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageClient) GetX(ctx context.Context, id uuid.UUID) *Image {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryImages queries the images edge of a Image.
func (c *ImageClient) QueryImages(i *Image) *ImageQuery {
	query := (&ImageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, image.ImagesTable, image.ImagesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryImageProcesses queries the image_processes edge of a Image.
func (c *ImageClient) QueryImageProcesses(i *Image) *ImageProcessQuery {
	query := (&ImageProcessClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, id),
			sqlgraph.To(imageprocess.Table, imageprocess.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, image.ImageProcessesTable, image.ImageProcessesColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImageClient) Hooks() []Hook {
	return c.hooks.Image
}

// Interceptors returns the client interceptors.
func (c *ImageClient) Interceptors() []Interceptor {
	return c.inters.Image
}

func (c *ImageClient) mutate(ctx context.Context, m *ImageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ImageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ImageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Image mutation op: %q", m.Op())
	}
}

// ImageProcessClient is a client for the ImageProcess schema.
type ImageProcessClient struct {
	config
}

// NewImageProcessClient returns a client for the ImageProcess from the given config.
func NewImageProcessClient(c config) *ImageProcessClient {
	return &ImageProcessClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `imageprocess.Hooks(f(g(h())))`.
func (c *ImageProcessClient) Use(hooks ...Hook) {
	c.hooks.ImageProcess = append(c.hooks.ImageProcess, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `imageprocess.Intercept(f(g(h())))`.
func (c *ImageProcessClient) Intercept(interceptors ...Interceptor) {
	c.inters.ImageProcess = append(c.inters.ImageProcess, interceptors...)
}

// Create returns a builder for creating a ImageProcess entity.
func (c *ImageProcessClient) Create() *ImageProcessCreate {
	mutation := newImageProcessMutation(c.config, OpCreate)
	return &ImageProcessCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ImageProcess entities.
func (c *ImageProcessClient) CreateBulk(builders ...*ImageProcessCreate) *ImageProcessCreateBulk {
	return &ImageProcessCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ImageProcess.
func (c *ImageProcessClient) Update() *ImageProcessUpdate {
	mutation := newImageProcessMutation(c.config, OpUpdate)
	return &ImageProcessUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageProcessClient) UpdateOne(ip *ImageProcess) *ImageProcessUpdateOne {
	mutation := newImageProcessMutation(c.config, OpUpdateOne, withImageProcess(ip))
	return &ImageProcessUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageProcessClient) UpdateOneID(id uuid.UUID) *ImageProcessUpdateOne {
	mutation := newImageProcessMutation(c.config, OpUpdateOne, withImageProcessID(id))
	return &ImageProcessUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ImageProcess.
func (c *ImageProcessClient) Delete() *ImageProcessDelete {
	mutation := newImageProcessMutation(c.config, OpDelete)
	return &ImageProcessDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ImageProcessClient) DeleteOne(ip *ImageProcess) *ImageProcessDeleteOne {
	return c.DeleteOneID(ip.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ImageProcessClient) DeleteOneID(id uuid.UUID) *ImageProcessDeleteOne {
	builder := c.Delete().Where(imageprocess.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageProcessDeleteOne{builder}
}

// Query returns a query builder for ImageProcess.
func (c *ImageProcessClient) Query() *ImageProcessQuery {
	return &ImageProcessQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeImageProcess},
		inters: c.Interceptors(),
	}
}

// Get returns a ImageProcess entity by its id.
func (c *ImageProcessClient) Get(ctx context.Context, id uuid.UUID) (*ImageProcess, error) {
	return c.Query().Where(imageprocess.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageProcessClient) GetX(ctx context.Context, id uuid.UUID) *ImageProcess {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrigin queries the origin edge of a ImageProcess.
func (c *ImageProcessClient) QueryOrigin(ip *ImageProcess) *ImageQuery {
	query := (&ImageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ip.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(imageprocess.Table, imageprocess.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, imageprocess.OriginTable, imageprocess.OriginColumn),
		)
		fromV = sqlgraph.Neighbors(ip.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryResult queries the result edge of a ImageProcess.
func (c *ImageProcessClient) QueryResult(ip *ImageProcess) *ImageQuery {
	query := (&ImageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ip.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(imageprocess.Table, imageprocess.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, imageprocess.ResultTable, imageprocess.ResultColumn),
		)
		fromV = sqlgraph.Neighbors(ip.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImageProcessClient) Hooks() []Hook {
	return c.hooks.ImageProcess
}

// Interceptors returns the client interceptors.
func (c *ImageProcessClient) Interceptors() []Interceptor {
	return c.inters.ImageProcess
}

func (c *ImageProcessClient) mutate(ctx context.Context, m *ImageProcessMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ImageProcessCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ImageProcessUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ImageProcessUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ImageProcessDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ImageProcess mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Image, ImageProcess []ent.Hook
	}
	inters struct {
		Image, ImageProcess []ent.Interceptor
	}
)
