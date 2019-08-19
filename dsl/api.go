package dsl

import (
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

// API defines a network service API. It provides the API name, description and other global
// properties. There may only be one API declaration in a given design package.
//
// API is a top level DSL. API takes two arguments: the name of the API and the
// defining DSL.
//
// The API properties are leveraged by the OpenAPI specification. The server
// expressions are also used by the server and the client tool code generators.
//
// Example:
//
//    var _ = NewAPI(
//        "adder",
//        Title("title"),                // Title used in documentation
//        Description("description"),    // Description used in documentation
//        Version("2.0"),                // Version of API
//        TermsOfService("terms"),       // Terms of use
//        Contact(                       // Contact info
//            Name("contact name"),
//            Email("contact email"),
//            URL("contact URL"),
//        ),
//        License(                       // License
//            Name("license name"),
//            URL("license URL"),
//        ),
//        Docs(                          // Documentation links
//            Description("doc description"),
//            URL("doc URL"),
//        ),
//        Server("addersvr",
//            Host(
//                "development",
//                URI("http://localhost:80"),
//                URI("grpc://localhost:8080"),
//            ),
//        ),
//    )
//
func NewAPI(name string, opts ...Option) *expr.APIExpr {
	if name == "" {
		eval.ReportError("API first argument cannot be empty")
		return nil
	}
	
	apiexpr := expr.NewAPIExpr(name, nil)

	// TODO: we need to set all data to API
	
	for _, o := range opts {
		o(apiexpr)
	}

	return apiexpr
}


// Title sets the API title. It is used by the generated OpenAPI specification.
//
// Title must appear in a API expression.
//
// Title accepts a single string argument.
//
// Example:
//
//    var _ = API(
//        "divider",
//        Title("divider API"),
//    )
//
func Title(val string) Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.APIExpr:
			e.Title = val
		default:
			// TODO: warning
		}
	}
}

// Version specifies the API version. One design describes one version.
//
// Version must appear in a API expression.
//
// Version accepts a single string argument.
//
// Example:
//
//    var _ = API(
//        "divider",
//        Version("1.0"),
//    )
//
func Version(ver string) Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.APIExpr:
			e.Version = ver
		default:
			// TODO: warning
		}
	}
}

// Contact sets the API contact information.
//
// Contact must appear in a API expression.
//
// Contact takes a single argument which is the defining DSL.
//
// Example:
//
//    var _ = API(
//        "divider",
//        Contact(
//            Name("support"),
//            Email("support@goa.design"),
//            URL("https://goa.design"),
//        ),
//    )
//
func Contact(opts ...Option) Option {
	contact := new(expr.ContactExpr)
	
	for _, o := range opts {
		o(contact)
	}

	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.APIExpr:
			e.Contact = contact
		default:
			// TODO: warning
		}
	}
}

// License sets the API license information.
//
// License must appear in a API expression.
//
// License takes a single argument which is the defining DSL.
//
// Example:
//
//    var _ = API(
//        "divider",
//        License(
//            Name("MIT"),
//            URL("https://github.com/goadesign/goa/blob/master/LICENSE"),
//        ),
//    )
//
func License(opts ...Option) Option {
	license := new(expr.LicenseExpr)
	
	for _, o := range opts {
		o(license)
	}

	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.APIExpr:
			e.License = license
		default:
			// TODO: warning
		}
	}
}

// TermsOfService describes the API terms of services or links to them.
//
// TermsOfService must appear in a API expression.
//
// TermsOfService takes a single argument which is the TOS text or URL.
//
// Example:
//
//    var _ = API(
//        "github",
//        TermsOfService("https://help.github.com/articles/github-terms-of-API/"),
//    )
//
func TermsOfService(terms string) Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.APIExpr:
			e.TermsOfService = terms
		default:
			// TODO: warning
		}
	}
}

// Name sets the contact or license name.
//
// Name must appear in a Contact or License expression.
//
// Name takes a single argument which is the contact or license name.
//
// Example:
//
//    var _ = API(
//        "divider",
//        License(
//            Name("MIT"),
//            URL("https://github.com/goadesign/goa/blob/master/LICENSE"),
//        )
//    )
//
func Name(name string) Option {
	return func(v eval.Expression) {
		switch def := v.(type) {
		case *expr.ContactExpr:
			def.Name = name
		case *expr.LicenseExpr:
			def.Name = name
		default:
			// TODO: warning
		}
	}
}

// Email sets the contact email.
//
// Email must appear in a Contact expression.
//
// Email takes a single argument which is the email address.
//
// Example:
//
//    var _ = API(
//        "divider",
//        Contact(
//            Email("support@goa.design"),
//        ),
//    )
//
func Email(email string) Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.ContactExpr:
			e.Email = email
		default:
			// TODO: warning
		}
	}
}