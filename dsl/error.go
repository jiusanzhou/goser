package dsl

import (
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

// Error describes a method error return value. The description includes a
// unique name (in the scope of the method), an optional type, description and
// DSL that further describes the type. If no type is specified then the
// built-in ErrorResult type is used. The DSL syntax is identical to the
// Attribute DSL.
//
// Error must appear in the Service (to define error responses that apply to all
// the service methods) or Method expressions.
//
// See Attribute for details on the Error arguments.
//
// Example:
//
//    var _ = Service(
//        "divider",
//        Error("invalid_arguments"), // Uses type ErrorResult
//
//        // Method which uses the default type for its response.
//        Method(
//            "divide",
//            Payload(DivideRequest),
//            Error("div_by_zero", DivByZero, "Division by zero"),
//        )
//    )
//
func Error(name string, opts ...Option) Option {
	att := new(expr.AttributeExpr)

	for _, o := range opts {
		o(att)
	}
	
	if att.Type == nil {
		att.Type = expr.ErrorResult
	}

	erro := &expr.ErrorExpr{AttributeExpr: att, Name: name}

	return func(v eval.Expression) {
		switch actual := v.(type) {
		case *expr.ServiceExpr:
			actual.Errors = append(actual.Errors, erro)
		case *expr.MethodExpr:
			actual.Errors = append(actual.Errors, erro)
		default:
			// TODO: warning
		}
	}
}

// Temporary qualifies an error type as describing temporary (i.e. retryable)
// errors.
//
// Temporary must appear in a Error expression.
//
// Temporary takes no argument.
//
// Example:
//
//    var _ = Service(
//        "divider",
//        Error(
//            "request_timeout",
//            Temporary(),
//        )
//    )
func Temporary() Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.AttributeExpr:
			if e.Meta == nil {
				e.Meta = make(expr.MetaExpr)
			}
			e.Meta["goa:error:temporary"] = nil
		default:
			// TODO: warning
		}
	}
}

// Timeout qualifies an error type as describing errors due to timeouts.
//
// Timeout must appear in a Error expression.
//
// Timeout takes no argument.
//
// Example:
//
//    var _ = Service(
//        "divider",
//        Error(
//            "request_timeout",
//            Timeout(),
//        })
//    })
func Timeout() Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.AttributeExpr:
			if e.Meta == nil {
				e.Meta = make(expr.MetaExpr)
			}
			e.Meta["goa:error:timeout"] = nil
		default:
			// TODO: warning
		}
	}
}

// Fault qualifies an error type as describing errors due to a server-side
// fault.
//
// Fault must appear in a Error expression.
//
// Fault takes no argument.
//
// Example:
//
//    var _ = Service(
//         "divider",
//         Error(
//                 "internal_error",
//                 Fault(),
//         )
//    })
func Fault() Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.AttributeExpr:
			if e.Meta == nil {
				e.Meta = make(expr.MetaExpr)
			}
			e.Meta["goa:error:fault"] = nil
		default:
			// TODO: warning
		}
	}
}
