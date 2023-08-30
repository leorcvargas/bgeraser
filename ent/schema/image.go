package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("format").NotEmpty().Immutable(),
		field.Int64("size").Positive().Immutable(),
		field.String("original_filename").NotEmpty().Immutable(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

func (Image) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Image.Type).
			Through("image_processes", ImageProcess.Type),
	}
}
