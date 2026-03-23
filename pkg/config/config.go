package config

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	FormalAPIKeyField = field.StringField(
		"formal-api-key",
		field.WithDisplayName("Formal API Key"),
		field.WithDescription("Your Formal API KEY."),
		field.WithIsSecret(true),
		field.WithRequired(true),
	)

	ConfigurationFields = []field.SchemaField{
		FormalAPIKeyField,
	}

	FieldRelationships = []field.SchemaFieldRelationship{}
)

//go:generate go run ./gen
var Config = field.NewConfiguration(
	ConfigurationFields,
	field.WithConstraints(FieldRelationships...),
	field.WithConnectorDisplayName("Formal"),
	field.WithHelpUrl("/docs/baton/formal"),
	field.WithIconUrl("/static/app-icons/formal.svg"),
)
