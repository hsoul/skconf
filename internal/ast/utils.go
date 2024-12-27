package ast

import "fmt"

func FindPropertyByName(propName string, properties []*PropertyDef) string {
	for _, prop := range properties {
		if ident, ok := prop.Key.(*Identifier); ok {
			if ident.Value == propName {
				switch v := prop.Value.(type) {
				case *Integer:
					return fmt.Sprintf("%d", v.Value)
				case *String:
					return v.Value
				}
			}
		}
	}
	return ""
}

func IsSameIdentifier(expr1, expr2 Expression) bool {
	ident1, ok1 := expr1.(*Identifier)
	ident2, ok2 := expr2.(*Identifier)
	if !ok1 || !ok2 {
		return false
	}
	return ident1.Value == ident2.Value
}
