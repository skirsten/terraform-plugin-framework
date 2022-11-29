package schema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Schema must satify the fwschema.Schema interface.
var _ fwschema.Schema = Schema{}

// Schema defines the structure and value types of resource data. This type
// is used as the resource.SchemaResponse type Schema field, which is
// implemented by the resource.DataSource type Schema method.
type Schema struct {
	// Attributes is the mapping of underlying attribute names to attribute
	// definitions.
	//
	// Names must only contain lowercase letters, numbers, and underscores.
	// Names must not collide with any Blocks names.
	Attributes map[string]Attribute

	// Blocks is the mapping of underlying block names to block definitions.
	//
	// Names must only contain lowercase letters, numbers, and underscores.
	// Names must not collide with any Attributes names.
	Blocks map[string]Block

	// Description is used in various tooling, like the language server, to
	// give practitioners more information about what this resource is,
	// what it's for, and how it should be used. It should be written as
	// plain text, with no special formatting.
	Description string

	// MarkdownDescription is used in various tooling, like the
	// documentation generator, to give practitioners more information
	// about what this resource is, what it's for, and how it should be
	// used. It should be formatted using Markdown.
	MarkdownDescription string

	// DeprecationMessage defines warning diagnostic details to display when
	// practitioner configurations use this resource. The warning diagnostic
	// summary is automatically set to "Resource Deprecated" along with
	// configuration source file and line information.
	//
	// Set this field to a practitioner actionable message such as:
	//
	//  - "Use examplecloud_other resource instead. This resource
	//    will be removed in the next major version of the provider."
	//  - "Remove this resource as it no longer is valid and
	//    will be removed in the next major version of the provider."
	//
	DeprecationMessage string

	// Version indicates the current version of the resource schema. Resource
	// schema versioning enables state upgrades in conjunction with the
	// [resource.ResourceWithStateUpgrades] interface. Versioning is only
	// required if there is a breaking change involving existing state data,
	// such as changing an attribute or block type in a manner that is
	// incompatible with the Terraform type.
	//
	// Versions are conventionally only incremented by one each release.
	Version int64
}

// ApplyTerraform5AttributePathStep applies the given AttributePathStep to the
// schema.
func (s Schema) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (any, error) {
	return fwschema.SchemaApplyTerraform5AttributePathStep(s, step)
}

// AttributeAtPath returns the Attribute at the passed path. If the path points
// to an element or attribute of a complex type, rather than to an Attribute,
// it will return an ErrPathInsideAtomicAttribute error.
func (s Schema) AttributeAtPath(ctx context.Context, p path.Path) (fwschema.Attribute, diag.Diagnostics) {
	return fwschema.SchemaAttributeAtPath(ctx, s, p)
}

// AttributeAtPath returns the Attribute at the passed path. If the path points
// to an element or attribute of a complex type, rather than to an Attribute,
// it will return an ErrPathInsideAtomicAttribute error.
func (s Schema) AttributeAtTerraformPath(ctx context.Context, p *tftypes.AttributePath) (fwschema.Attribute, error) {
	return fwschema.SchemaAttributeAtTerraformPath(ctx, s, p)
}

// GetAttributes returns the Attributes field value.
func (s Schema) GetAttributes() map[string]fwschema.Attribute {
	return schemaAttributes(s.Attributes)
}

// GetBlocks returns the Blocks field value.
func (s Schema) GetBlocks() map[string]fwschema.Block {
	return schemaBlocks(s.Blocks)
}

// GetDeprecationMessage returns the DeprecationMessage field value.
func (s Schema) GetDeprecationMessage() string {
	return s.DeprecationMessage
}

// GetDescription returns the Description field value.
func (s Schema) GetDescription() string {
	return s.Description
}

// GetMarkdownDescription returns the MarkdownDescription field value.
func (s Schema) GetMarkdownDescription() string {
	return s.MarkdownDescription
}

// GetVersion returns the Version field value.
func (s Schema) GetVersion() int64 {
	return s.Version
}

// Type returns the framework type of the schema.
func (s Schema) Type() attr.Type {
	return fwschema.SchemaType(s)
}

// TypeAtPath returns the framework type at the given schema path.
func (s Schema) TypeAtPath(ctx context.Context, p path.Path) (attr.Type, diag.Diagnostics) {
	return fwschema.SchemaTypeAtPath(ctx, s, p)
}

// TypeAtTerraformPath returns the framework type at the given tftypes path.
func (s Schema) TypeAtTerraformPath(ctx context.Context, p *tftypes.AttributePath) (attr.Type, error) {
	return fwschema.SchemaTypeAtTerraformPath(ctx, s, p)
}

// schemaAttributes is a resource to fwschema type conversion function.
func schemaAttributes(attributes map[string]Attribute) map[string]fwschema.Attribute {
	result := make(map[string]fwschema.Attribute, len(attributes))

	for name, attribute := range attributes {
		result[name] = attribute
	}

	return result
}

// schemaBlocks is a resource to fwschema type conversion function.
func schemaBlocks(blocks map[string]Block) map[string]fwschema.Block {
	result := make(map[string]fwschema.Block, len(blocks))

	for name, block := range blocks {
		result[name] = block
	}

	return result
}
