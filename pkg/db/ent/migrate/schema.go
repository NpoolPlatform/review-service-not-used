// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "object_type", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "object_id", Type: field.TypeUUID},
		{Name: "reviewer_id", Type: field.TypeUUID},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"wait", "approved", "rejected"}},
		{Name: "message", Type: field.TypeString},
		{Name: "trigger", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
	}
	// ReviewRulesColumns holds the columns for the "review_rules" table.
	ReviewRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "object_type", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
		{Name: "rules", Type: field.TypeString, Default: "{}"},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// ReviewRulesTable holds the schema information for the "review_rules" table.
	ReviewRulesTable = &schema.Table{
		Name:       "review_rules",
		Columns:    ReviewRulesColumns,
		PrimaryKey: []*schema.Column{ReviewRulesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ReviewsTable,
		ReviewRulesTable,
	}
)

func init() {
}
