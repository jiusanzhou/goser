// This test demonstrates all the possible usage of Type.
package dsl_test

var _ = API(
	"dsl_spec",

	// API title for docs
	Title("Optional API title"),

	// API description for docs
	Description("Optional API description"),

	// API version
	Version("1.0"),

	// API support information.
	Contact(
		Name("contact name"),
		Email("contact@goser.zoe.im"),
		URL("https://goser.zoe.im"),
	),

	// API Licensing information
	License(
		Name("License name"),
		URL("https://goser.zoe.im/license"),
	),

	// Docs allows linking to external documentation.
	Docs(
		Description("Optional description"),
		URL("https://goser.zoe.im/getting-started"),
	),

	// Server describes a single API host and may appear more than once.
	// URL must include the protocol and hostname and may include a port.
	// The hostname and port may use parameters to define possible
	// alternative values.
	Server(
		"calcsvr",

		Description("Optional description"),

		Host(
			"production",

            Description("Production host.")
            // URIs can be parameterized using {param} notation.
            URI("https://{version}.goser.zoe.im/calc")
            URI("grpcs://{version}.goser.zoe.im")

            // Variable describes a URI variable.
            Variable(
				"version", String,

				Description("API version"),
                // URI parameters must have a default value and/or an
                // enum validation.
            	Default("v1")
            ),
		),

		Host(
			"development",
			Description("Development hosts."),
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://localhost:80/calc"),
			URI("grpc://localhost:8080"),
		),
	),

	Meta("meta_key", "meta_value"),
)

// The Service expression defines a single service. There may be any number of
// Service declarations in one design.
var _ = Service(
	"service",
	// Service description for code comments and docs
	Description("Optional service description"),
	// Server definitions that appear in the Service DSL override all the API
	// level definitions.
	Server(
		"https://service.goser.zoe.im:443",
		Description("Service specific server description"),
	),
	// Docs allows linking to external documentation.
	Docs(
		Description("Optional description"),
		URL("https://goser.zoe.im/getting-started"),
	),

)