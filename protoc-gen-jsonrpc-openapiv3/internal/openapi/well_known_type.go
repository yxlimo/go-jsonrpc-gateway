package openapi

var wktSchemas = map[string]*openapiSchemaObject{
	".google.protobuf.FieldMask": {
		Type: "string",
	},
	".google.protobuf.Timestamp": {
		Type:   "string",
		Format: "date-time",
	},
	".google.protobuf.Duration": {
		Type: "string",
	},
	".google.protobuf.StringValue": {
		Type: "string",
	},
	".google.protobuf.BytesValue": {
		Type:   "string",
		Format: "byte",
	},
	".google.protobuf.Int32Value": {
		Type:   "integer",
		Format: "int32",
	},
	".google.protobuf.UInt32Value": {
		Type:   "integer",
		Format: "int64",
	},
	".google.protobuf.Int64Value": {
		Type:   "string",
		Format: "int64",
	},
	".google.protobuf.UInt64Value": {
		Type:   "string",
		Format: "uint64",
	},
	".google.protobuf.FloatValue": {
		Type:   "number",
		Format: "float",
	},
	".google.protobuf.DoubleValue": {
		Type:   "number",
		Format: "double",
	},
	".google.protobuf.BoolValue": {
		Type: "boolean",
	},
	".google.protobuf.Empty": {
		Type: "object",
	},
	".google.protobuf.Struct": {
		Type: "object",
	},
	".google.protobuf.Value": {
		Type: "object",
	},
	".google.protobuf.ListValue": {
		Type: "array",
		Items: &openapiSchemaObject{
			Type: "object",
		},
	},
	".google.protobuf.NullValue": {
		Type: "string",
	},
}

func (s *Openapi) jsonrpcRequestSchema(method string, req *openapiSchemaObject) *openapiRequestBodyObject {
	return &openapiRequestBodyObject{
		Content: map[string]*openapiMediaTypeObject{
			"application/json": {Schema: &openapiSchemaObject{
				Type: "object",
				Properties: map[string]*openapiSchemaObject{
					"jsonrpc": {Type: "string", Enum: []string{"2.0"}},
					"method": {
						Type:    "string",
						Pattern: "^" + method + "$",
					},
					"params": req,
					"id":     {Type: "string"},
				},
				Required: []string{"jsonrpc", "method", "id"},
			}},
		},
	}
}

func (s *Openapi) jsonrpcResponseSchema(method string, res *openapiSchemaObject) *openapiResponseObject {
	return &openapiResponseObject{
		Content: map[string]*openapiMediaTypeObject{
			"application/json": {
				Schema: &openapiSchemaObject{
					Type: "object",
					Properties: map[string]*openapiSchemaObject{
						"jsonrpc": {Type: "string", Enum: []string{"2.0"}},
						"method":  {Type: "string", Pattern: "^" + method + "$"},
						"result":  res,
						"id":      {Type: "string"},
					},
					Required: []string{"result"},
				},
			},
		},
	}
}
