package openapi

type openapiObject struct {
	Version    string                        `json:"openapi"`
	Info       openapiInfoObject             `json:"info"`
	Paths      map[string]*openapiPathObject `json:"paths"`
	Components openapiComponentsObject       `json:"components,omitempty"`
}

type openapiInfoObject struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type openapiPathObject struct {
	Post *openapiOperationObject `json:"post"`
}

type openapiOperationObject struct {
	Summary     string                            `json:"summary,omitempty"`
	RequestBody *openapiRequestBodyObject         `json:"requestBody"`
	Responses   map[string]*openapiResponseObject `json:"responses"`
}

type openapiRequestBodyObject struct {
	Content map[string]*openapiMediaTypeObject `json:"content"`
}

type openapiResponseObject struct {
	Description string                             `json:"description"`
	Content     map[string]*openapiMediaTypeObject `json:"content"`
}

type openapiMediaTypeObject struct {
	Schema *openapiSchemaObject `json:"schema"`
}

type openapiSchemaObject struct {
	Ref                  string                          `json:"$ref,omitempty"`
	OneOf                []*openapiSchemaObject          `json:"oneOf,omitempty"`
	Type                 string                          `json:"type,omitempty"`
	Nullable             bool                            `json:"nullable,omitempty"`
	Enum                 []string                        `json:"enum,omitempty"`
	Items                *openapiSchemaObject            `json:"items,omitempty"`
	Properties           map[string]*openapiSchemaObject `json:"properties,omitempty"`
	AdditionalProperties *openapiSchemaObject            `json:"additionalProperties,omitempty"`
	Format               string                          `json:"format,omitempty"`
	Minimum              float64                         `json:"minimum,omitempty"`
	Required             []string                        `json:"required,omitempty"`
	Pattern              string                          `json:"pattern,omitempty"`
}

type openapiComponentsObject struct {
	Schemas map[string]*openapiSchemaObject `json:"schemas"`
}
