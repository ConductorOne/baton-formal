package config

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	FormalApiKeyField = field.StringField(
		"formal-api-key",
		field.WithRequired(true),
		field.WithDescription("Your Formal API KEY."),
	)

	ConfigurationFields = []field.SchemaField{
		FormalApiKeyField,
	}
)

//go:generate go run ./gen
var Config = field.NewConfiguration(ConfigurationFields)
