package tools

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

const GoTypeUnknown GoType = ""

func GetGoType(o interface{}) GoType {
	if g, ok := o.(GoType); ok {
		return g
	}

	if s, ok := o.(string); ok {
		return GoType(s)
	}

	return GoTypeUnknown
}

func (j GoType) MustGetOptionalTypeString() jen.Code {
	switch j {
	case GoTypeBoolean:
		return jen.Op("*").Bool()
	case GoTypeDuration:
		return jen.Op("*").Qual("time", "Duration")
	case GoTypeInt:
		return jen.Op("*").Int()
	case GoTypeString:
		return jen.Op("*").String()
	case GoTypeStringSlice:
		return jen.Index().String()
	default:
		panic(fmt.Errorf("unknown type '%s'", j))
	}
}

func (j GoType) MustGetTypeString() jen.Code {
	switch j {
	case GoTypeBoolean:
		return jen.Bool()
	case GoTypeDuration:
		return jen.Qual("time", "Duration")
	case GoTypeInt:
		return jen.Int()
	case GoTypeString:
		return jen.String()
	case GoTypeStringSlice:
		return jen.Index().String()
	default:
		panic(fmt.Errorf("unknown type '%s'", j))
	}
}

func (j GoType) MustGetDefaultValueCondition(id string) jen.Code {
	switch j {
	case GoTypeBoolean:
		return jen.Id(id)
	case GoTypeDuration:
		return jen.Id(id).Op("!=").Nil()
	case GoTypeInt:
		return jen.Id(id).Op("!=").Lit(0)
	case GoTypeString:
		return jen.Id(id).Op("!=").Lit("")
	case GoTypeStringSlice:
		return jen.Len(jen.Id(id)).Op(">").Lit(0)
	default:
		panic(fmt.Errorf("unknown type '%s'", j))
	}
}

func (j GoType) MustGetCommandArgsAppend(argsId, id, format string) jen.Code {
	switch j {
	case GoTypeBoolean:
		if format == "" {
			format = "%t"
		}

		return jen.Id(argsId).Op("=").Append(
			jen.Id(argsId),
			jen.Qual("fmt", "Sprintf").Call(jen.Lit(format), jen.Id(id)),
		)
	case GoTypeDuration:
		if format == "" {
			format = "%s"
		}

		return jen.Id(argsId).Op("=").Append(
			jen.Id(argsId),
			jen.Qual("fmt", "Sprintf").Call(jen.Lit(format), jen.Id(id)),
		)
	case GoTypeInt:
		if format == "" {
			format = "%d"
		}

		return jen.Id(argsId).Op("=").Append(
			jen.Id(argsId),
			jen.Qual("fmt", "Sprintf").Call(jen.Lit(format), jen.Id(id)),
		)
	case GoTypeString:
		if format == "" {
			return jen.Id(argsId).Op("=").Append(
				jen.Id(argsId),
				jen.Id(id),
			)
		}

		return jen.Id(argsId).Op("=").Append(
			jen.Id(argsId),
			jen.Qual("fmt", "Sprintf").Call(jen.Lit(format), jen.Id(id)),
		)
	case GoTypeStringSlice:
		if format == "" {
			return jen.Id(argsId).Op("=").Append(
				jen.Id(argsId),
				jen.Id(id).Op("..."),
			)
		}

		tmpId := fmt.Sprintf("_%s", id)

		return jen.For(jen.List(jen.Id("_"), jen.Id(tmpId)).Op(":=").Range().Id(id)).Block(
			jen.Id(argsId).Op("=").Append(
				jen.Id(argsId),
				jen.Qual("fmt", "Sprintf").Call(jen.Lit(format), jen.Id(tmpId)),
			),
		)
	default:
		panic(fmt.Errorf("unknown type '%s'", j))
	}
}

func (g *CommandGroup) IsRoot() bool {
	return g.Package == "@"
}
