package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DynamicSchema struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec DynamicSchemaSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status DynamicSchemaStatus `json:"status"`
}

type DynamicSchemaSpec struct {
	Embed             bool              `json:"embed,omitempty"`
	EmbedType         string            `json:"embedType,omitempty"`
	PluralName        string            `json:"pluralName,omitempty"`
	ResourceMethods   []string          `json:"resourceMethods,omitempty"`
	ResourceFields    map[string]Field  `json:"resourceFields,omitempty"`
	ResourceActions   map[string]Action `json:"resourceActions,omitempty"`
	CollectionMethods []string          `json:"collectionMethods,omitempty"`
	CollectionFields  map[string]Field  `json:"collectionFields,omitempty"`
	CollectionActions map[string]Action `json:"collectionActions,omitempty"`
	CollectionFilters map[string]Filter `json:"collectionFilters,omitempty"`
	IncludeableLinks  []string          `json:"includeableLinks,omitempty"`
}

type DynamicSchemaStatus struct {
	Fake string `json:"fake,omitempty"`
}

type Field struct {
	Type         string   `json:"type,omitempty"`
	Default      Values   `json:"default,omitempty"`
	Unique       bool     `json:"unique,omitempty"`
	Nullable     bool     `json:"nullable,omitempty"`
	Create       bool     `json:"create,omitempty"`
	Required     bool     `json:"required,omitempty"`
	Update       bool     `json:"update,omitempty"`
	MinLength    int64    `json:"minLength,omitempty"`
	MaxLength    int64    `json:"maxLength,omitempty"`
	Min          int64    `json:"min,omitempty"`
	Max          int64    `json:"max,omitempty"`
	Options      []string `json:"options,omitempty"`
	ValidChars   string   `json:"validChars,omitempty"`
	InvalidChars string   `json:"invalidChars,omitempty"`
	Description  string   `json:"description,omitempty"`
}

type Values struct {
	StringValue      string   `json:"stringValue,omitempty"`
	IntValue         int      `json:"intValue,omitempty"`
	BoolValue        bool     `json:"boolValue,omitempty"`
	StringSliceValue []string `json:"stringSliceValue,omitempty"`
}

type Action struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}

type Filter struct {
	Modifiers []string `json:"modifiers,omitempty"`
}

type ListOpts struct {
	Filters map[string]string `json:"filters,omitempty"`
}
