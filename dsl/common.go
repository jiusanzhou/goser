package dsl

import (
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

// Docs provides external documentation URLs. It is used by the generated
// OpenAPI specification.
//
// Docs must appear in an API, Service, Method or Attribute expr.
//
// Docs takes a single argument which is the defining DSL.
//
// Example:
//
//    var _ = API(
//        "cellar",
//        Docs(
//            Description("Additional documentation"),
//            URL("https://goa.design"),
//        ),
//    )
//
func Docs(opts ...Option) Option {
	docs := new(expr.DocsExpr)

	for _, o := range opts {
		o(docs)
	}
	
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.APIExpr:
			e.Docs = docs
		case *expr.ServiceExpr:
			e.Docs = docs
		case *expr.MethodExpr:
			e.Docs = docs
		case *expr.AttributeExpr:
			e.Docs = docs
		case *expr.HTTPFileServerExpr:
			e.Docs = docs
		default:
			// TODO: warning
		}
	}
}

// Description sets the expression description.
//
// Description may appear in API, Docs, Type or Attribute.
// Description may also appear in Response and FileServer.
//
// Description accepts one arguments: the description string.
//
// Example:
//
//    API(
//        "adder",
//        Description("Adder API"),
//    )
//
func Description(d string) Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.APIExpr:
			e.Description = d
		case *expr.ServerExpr:
			e.Description = d
		case *expr.HostExpr:
			e.Description = d
		case *expr.ServiceExpr:
			e.Description = d
		case *expr.ResultTypeExpr:
			e.Description = d
		case *expr.AttributeExpr:
			e.Description = d
		case *expr.DocsExpr:
			e.Description = d
		case *expr.MethodExpr:
			e.Description = d
		case *expr.ExampleExpr:
			e.Description = d
		case *expr.SchemeExpr:
			e.Description = d
		case *expr.HTTPResponseExpr:
			e.Description = d
		case *expr.HTTPFileServerExpr:
			e.Description = d
		case *expr.GRPCResponseExpr:
			e.Description = d
		default:
			// TODO: warning
		}
	}
}

// URL sets the contact, license or external documentation URL.
//
// URL must appear in Contact, License or Docs.
//
// URL accepts a single argument which is the URL.
//
// Example:
//
//    Docs(
//        URL("https://goa.design"),
//    )
//
func URL(url string) Option {
	return func(v eval.Expression) {
		switch def := v.(type) {
		case *expr.ContactExpr:
			def.URL = url
		case *expr.LicenseExpr:
			def.URL = url
		case *expr.DocsExpr:
			def.URL = url
		default:
			// TODO: warning
		}
	}
}

// Type finish docs
func Type(typ expr.DataType) Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.AttributeExpr:
			e.Type = typ
		default:
			// TODO: warning
		}
	}
}