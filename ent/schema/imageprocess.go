package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ImageProcess holds the schema definition for the ImageProcess entity.
type ImageProcess struct {
	ent.Schema
}

// Fields of the ImageProcess.
func (ImageProcess) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("image_id", uuid.UUID{}),
		field.UUID("result_id", uuid.UUID{}).Optional(),
		field.String("kind").NotEmpty().Immutable(),
		field.Time("finished_at").Nillable().Optional(),
		field.Time("errored_at").Nillable().Optional(),
		field.String("error_reason").Nillable().Optional(),
	}
}

func (ImageProcess) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		// index.Fields("origin_id"),
		// index.Fields("result_id"),
	}
}

// // Edges of the ImageProcess.
func (ImageProcess) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("origin", Image.Type).Required().Unique().Field("image_id"),
		edge.To("result", Image.Type).Unique().Field("result_id"),
	}
}
