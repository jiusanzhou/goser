// This test demonstrates all the possible usage of Type.
package dsl_test

// BasicType shows the basic usage for Type.
//
// Type is the DSL function used to describe user types. A user type is an
// object type that can be used to define request and response types or
// attributes of request and response types.
//
// Type takes the name of the type as first argument. This name must be unique
// across all types in a given design package.
var BasicType = Type(
	"name",

	// Optional description used in code comments and docs.
	Description("Optional description"),

	// Attribute defines a field of the type, see below for all possible
	// usage.
	Attribute("an_attribute", String),

	// Required lists the required attributes. Required takes one or more
	// attribute names and can appear one or more times.
	Required("an_attribute"),
)

// BasicResultType shows the basic usage for ResultType.
//
// ResultType is the DSL function used to describe result types. A result type is a
// special kind of type that adds the concept of views: A view defines a subset
// of the type attributes to be rendered. This is used to describe *response*
// types where an method may render different attributes depending on the
// request state or when different methods render the type differently (for
// example a list method may render less attribute than a method that retrieves
// a single value). All result types muse define a default view. The default view
// is the view named "default".
//
// Result type takes a result type identifier (as defined by RFC 6838) as first
// argument. This identifier must be unique across all result types in a given
// package.
var BasicResultType = ResultType(
	"application/vnd.goser.result",

	// Optional description used in code comments and docs.
	Description("Optional description"),

	// Attributes lists the result type attributes identically to Type.
	Attributes(
		Attribute("an_attribute", String),
		Required("an_attribute"),
	),

	// View defines a view. View may appear one or more times in a ResultType
	// expression.
	View(
		"default",
		// There is no need to repeat the attribute properties when
		// listing the view attributes, only the name is required.
		Attribute("an_attribute"),
	),
)

// AllTypes is a type definition with attributes using all possible data types.
var AllTypes = Type(
	"allTypes",
	Description("An object with attributes of all possible types"),

	Attribute("string", String),
	Attribute("bytes", Bytes),
	Attribute("boolean", Boolean),
	Attribute("int32", Int32),
	Attribute("int64", Int64),
	Attribute("float32", Float32),
	Attribute("float64", Float64),
	Attribute("any", Any),
	Attribute("array", ArrayOf(String)),
	Attribute("map", MapOf(String, String)),
	Attribute(
		"object",
		Description("Inner type"),
		Attribute("inner_attribute", String),
		Required("inner_attribute"),
	),
	Attribute("user", AUserType),
	Attribute("result", AResultType),
	Attribute("collection", CollectionOf(AResultType)),
)

// AUserType is a type used to define an attribute in AllTypes.
var AUserType = Type(
	"UserType",
	Description("Optional description"),
	Attribute("required", String),
	Attribute("optional", String),
	Required("required"),
)

// AResultType is a result type used to define an attribute in AllTypes.
var AResultType = ResultType(
	"ResultType",
	Description("Optional description"),
	Attributes(
		Attribute("optional", String),
		Attribute("required", String),
		Required("required"),
	),
	View("default",
		Attribute("optional", String),
		Attribute("required", String),
	),
	View(
		"tiny",
		Attribute("required", String),
	),
)

// ACollectionResult type shows all the possible DSL of CollectionOf.
var ACollectionResult = CollectionOf(
	AResultType,
	// View allows defining collection specific views.
	// The view is defined using the attributes of the element of the
	// collection.
	View(
		"collection",
		Attribute("optional"),
	),

	// View can also refer to existing views defined in the element result
	// type. If no View is specified (i.e. no DSL argument is provided to
	// CollectionOf) then all the element result type views are inherited.
	View("tiny"),
)

// AArrayType is a array type with a element validation.
var AArrayType = ArrayOf(
	String,
	Pattern("regexp"),
)

// AMapType is a map with element and key validations.
var AMapType = MapOf(
	// Key is used to define validations that apply to the keys of the map.
	Key(String, Pattern("keyregexp")),

	// Value is used to define validations that apply to the values of the map.
	Value(String, Pattern("valueregexp")),
)

// Attrs is a type definition which demonstrates the different ways attributes
// may be defined.
var Attrs = Type(
	"Attributes",
	// Attribute defined with a name and a type
	Attribute("name", String),
	// Attribute defined with a name, a type and a description
	Attribute("name_2", String, Description("description")),
	// Attribute defined with a name, a type and validations
	Attribute("name_3", String, MinLength(10)),
	// Attribute defined with a name, a type, a description, validations, a
	// default value and an example.
	Attribute(
		"name_4", String,
		MinLength(10),
		MaxLength(100),
		Default("default value"),
		Example("example value"),
		Example("another example value"),
		Example("title", "example text"),
	),
)

// Validations is a type definition with all possible validations.
var Validations = Type(
	"Validations",

	Description("An object with attributes with all possible validations"),
	Attribute(
		"string", String,
		// MinLength specifies the minimum number of characters in the
		// string.
		MinLength(5),
		// MaxLength specifies the maximum number of characters in the
		// string.
		MaxLength(100),
		// Pattern specifies a regular expression that the value must
		// validate.
		Pattern(`^A.*@goa\.design`),
		// Format specifies a format the string must comply to.
		// See ValidationFormat constants in design package for the list
		// of supported formats.
		Format(FormatEmail),
		// Enum specifies the list of possible values (a real design
		// would probably not use other validations together with Enum)
		Enum("support@goa.design", "info@goa.design"),
	),

	Attribute(
		"bytes", Bytes,
		MinLength(5),
		MaxLength(10),
		Enum([]byte{'1', '2'}, []byte{'3', '4'}),
	),

	Attribute(
		"number", Int64,
		// Minimum specifies the minimum value for the number
		// (inclusive)
		Minimum(1),
		// Minimum specifies the maximum value for the number
		// (inclusive)
		Maximum(10),
		Enum(1, 2, 3, 4),
	),

	Attribute(
		"array", ArrayOf(String),
		MinLength(2),
		MaxLength(10),
		Enum([]string{"a", "b"}, []string{"c", "d"}),
	),

	Attribute(
		"map", ArrayOf(String),
		MinLength(5),
		MaxLength(10),
	),
)

// Embedded is a type definition with embedded object attribute definitions.
var Embedded = Type(
	"Embedded",

	// Attribute accepts either a type or a DSL describing a type as second
	// argument.
	Attribute(
		"embedded",
		ObjectOf(
			Attribute("inner1", String),
			Attribute("inner2", Int32),
		),
		Description("Inner object description"),
		Required("inner"),
	),
)

// Recursive is a type definition with recursive attribute definitions.
var Recursive = Type(
	"Recursive",
	// Attribute allows specifying a type using its name rather than a Go
	// variable to make it possible to describe recursive data structures
	// without running into circular dependency compilation errors.
	Attribute("recursive", "Recursive"),
	Attribute("recursives", ArrayOf("Recursive")),
)