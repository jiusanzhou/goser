package dsl

import (
	"fmt"

	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

// Attribute describes a field of an object.
//
// An attribute has a name, a type and optionally a default value, an example
// value and validation rules.
//
// The type of an attribute can be one of:
//
// * The primitive types Boolean, Float32, Float64, Int, Int32, Int64, UInt,
// UInt32, UInt64, String or Bytes.
//
// * A user type defined via the Type function.
//
// * An array defined using the ArrayOf function.
//
// * An map defined using the MapOf function.
//
// * An object defined inline using Attribute to define the type fields
// recursively.
//
// * The special type Any to indicate that the attribute may take any of the
// types listed above.
//
// Attribute must appear in ResultType, Type, Attribute or Attributes.
//
// Attribute accepts one to four arguments, the valid usages of the function
// are:
//
//    Attribute(name)       // Attribute of type String with no description, no
//                          // validation, default or example value
//
//    Attribute(name, fn)   // Attribute of type object with inline field
//                          // definitions, description, validations, default
//                          // and/or example value
//
//    Attribute(name, type) // Attribute with no description, no validation,
//                          // no default or example value
//
//    Attribute(name, type, fn) // Attribute with description, validations,
//                              // default and/or example value
//
//    Attribute(name, type, description)     // Attribute with no validation,
//                                           // default or example value
//
//    Attribute(name, type, description, fn) // Attribute with description,
//                                           // validations, default and/or
//                                           // example value
//
// Goser only keep one style:
//
//   Attribute(name, type, opts ...Option)
//
// Where name is a string indicating the name of the attribute, type specifies
// the attribute type (see above for the possible values), description a string
// providing a human description of the attribute and fn the defining DSL if
// any.
//
// When defining the type inline using Attribute recursively the function takes
// the second form (name and DSL defining the type). The description can be
// provided using the Description function in this case.
//
// Examples:
//
//    Attribute(
//        "name", String,
//        Pattern("^foo"),                 // Adds a validation rule
//    )
//
//    Attribute(
//        "driver", Person,
//        Required("name"),                // Add required field to list of
//    )                                    // fields already required in Person
//
//    Attribute(
//        "age", Int32,
//        Description("description"),
//        Minimum(2),                       // Sets both a description and
//                                          // validations
//    )
//
// The definition below defines an attribute inline. The resulting type
// is an object with three attributes "name", "age" and "child". The "child"
// attribute is itself defined inline and has one child attribute "name".
//
//    Attribute(
//        "driver", Object,                  // Define type inline
//        Description("Composite attribute") // Set description
//
//        Attribute("name", String),         // Child attribute
//        Attribute(
//            "age", Type(Int32),            // Another child attribute
//            Description("Age of driver"),
//            Default(42),
//            Minimum(2),
//        ),
//        Attribute(
//            "child", Object                // Defines a child attribute
//            Attribute("name", String),     // Grand-child attribute
//            Required("name"),
//        ),
//
//        Required("name", "age"),           // List required attributes
//    )
//
func Attribute(name string, opts ...Option) Option {
	// TODO:
	return func(v eval.Expression) {
		
	}
}

// Field is syntactic sugar to define an attribute with the "rpc:tag" meta
// set with the value of the first argument.
//
// Field must appear wherever Attribute can.
//
// Field takes the same arguments as Attribute with the addition of the tag
// value as first argument.
//
// Example:
//
//     Field(
//         1, "ID", 
//         Type(String),
//         Pattern("[0-9]+"),
//     })
//
func Field(tag interface{}, name string, opts ...Option) Option {
	return Attribute(name, append(opts, Meta("rpc:tag", fmt.Sprintf("%v", tag)))...)
}

// Default sets the default value for an attribute.
//
// Default must appear in an Attribute DSL.
//
// Default takes one parameter: the default value.
func Default(def interface{}) Option {
	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.AttributeExpr:
			if e.Type != nil && !e.Type.IsCompatible(def) {
				// TODO:
				// eval.ReportError("default value %#v is incompatible with attribute of type %s",
				// 	def, expr.QualifiedTypeName(e.Type))
				return
			}
			e.SetDefault(def)
		default:
			// TODO:
		}
	}
}

// Example provides an example value for a type, a parameter, a header or any
// attribute. Example supports two syntaxes: one syntax accepts two arguments
// where the first argument is a summary describing the example and the second a
// value provided directly or via a DSL which may also specify a long
// description. The other syntax accepts a single argument and is equivalent to
// using the first syntax where the summary is the string "default". When using
// a DSL the Value function can be used to provide the example value.
//
// If no example is explicitly provided in an attribute expression then a random
// example is generated unless the "swagger:example" meta is set to "false".
// See Meta.
//
// Example must appear in a Attributes, Attribute, Params, Param, Headers or
// Header DSL.
//
// Example takes one or two arguments: an optional summary and the example value
// or defining DSL.
//
// Examples:
//
//    Params(func() {
//        Param("ZipCode:zip-code", String, "Zip code filter", func() {
//            Example("Santa Barbara", "93111")
//            Example("93117") // same as Example("default", "93117")
//        })
//    })
//
//    Attributes(func() {
//        Attribute("ID", Int64, "ID is the unique bottle identifier")
//        Example("The first bottle", func() {
//            Description("This bottle has an ID set to 1")
//            Value(Val{"ID": 1})
//        })
//        Example("Another bottle", func() {
//            Description("This bottle has an ID set to 5")
//            Value(Val{"ID": 5})
//        })
//    })
//
func Example(summary string, opts ...Option) Option {
	
	if summary == "" {
		summary = "default"
	}

	ex := &expr.ExampleExpr{Summary: summary}

	for _, o := range opts {
		o(ex)
	}

	// TODO: finish value
	if ex.Value == nil {
		// TODO:
	}

	return func(v eval.Expression) {
		switch e := v.(type) {
		case *expr.AttributeExpr:
			e.UserExamples = append(e.UserExamples, ex)
		default:
			// TODO: warning
		}
	}
}

func parseAttributeArgs(baseAttr *expr.AttributeExpr, args ...interface{}) (expr.DataType, string, func()) {
	var (
		dataType    expr.DataType
		description string
		fn          func()
		ok          bool
	)

	parseDataType := func(expected string, index int) {
		if name, ok2 := args[index].(string); ok2 {
			// Lookup type by name
			if dataType = expr.Root.UserType(name); dataType == nil {
				eval.InvalidArgError(expected, args[index])
			}
			return
		}
		if dataType, ok = args[index].(expr.DataType); !ok {
			eval.InvalidArgError(expected, args[index])
		}
	}
	parseDescription := func(expected string, index int) {
		if description, ok = args[index].(string); !ok {
			eval.InvalidArgError(expected, args[index])
		}
	}
	parseDSL := func(index int, success, failure func()) {
		if fn, ok = args[index].(func()); ok {
			success()
			return
		}
		failure()
	}

	success := func() {}

	switch len(args) {
	case 0:
		if baseAttr != nil {
			dataType = baseAttr.Type
		} else {
			dataType = expr.String
		}
	case 1:
		success = func() {
			if baseAttr != nil {
				dataType = baseAttr.Type
			}
		}
		parseDSL(0, success, func() { parseDataType("type, type name or func()", 0) })
	case 2:
		parseDataType("type or type name", 0)
		parseDSL(1, success, func() { parseDescription("string or func()", 1) })
	case 3:
		parseDataType("type or type name", 0)
		parseDescription("string", 1)
		parseDSL(2, success, func() { eval.InvalidArgError("func()", args[2]) })
	default:
		eval.ReportError("too many arguments in call to Attribute")
	}

	return dataType, description, fn
}
