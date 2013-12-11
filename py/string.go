// String objects
//
// Note that string objects in Python are arrays of unicode
// characters.  However we are using the native Go string which is
// UTF-8 encoded.  This makes very little difference most of the time,
// but care is needed when indexing, slicing or iterating through
// strings.

package py

import (
	"unicode/utf8"
)

type String string

var StringType = NewType("string",
	`str(object='') -> str
str(bytes_or_buffer[, encoding[, errors]]) -> str

Create a new string object from the given object. If encoding or
errors is specified, then the object must expose a data buffer
that will be decoded using the given encoding and error handler.
Otherwise, returns the result of object.__str__() (if defined)
or repr(object).
encoding defaults to sys.getdefaultencoding().
errors defaults to 'strict'.`)

// Type of this object
func (s String) Type() *Type {
	return StringType
}

// Intern s possibly returning a reference to an already interned string
func (s String) Intern() String {
	// fmt.Printf("FIXME interning %q\n", s)
	return s
}

func (s String) M__bool__() Object {
	return NewBool(len(s) > 0)
}

func (s String) M__len__() Object {
	return Int(utf8.RuneCountInString(string(s)))
}

func (a String) M__add__(other Object) Object {
	if b, ok := other.(String); ok {
		return a + b
	}
	return NotImplemented
}

func (a String) M__radd__(other Object) Object {
	if b, ok := other.(String); ok {
		return b + a
	}
	return NotImplemented
}

func (a String) M__iadd__(other Object) Object {
	return a.M__add__(other)
}

func (a String) M__mul__(other Object) Object {
	if b, ok := convertToInt(other); ok {
		newString := String("")
		for i := 0; i < int(b); i++ {
			newString += a
		}
		return newString
	}
	return NotImplemented
}

func (a String) M__rmul__(other Object) Object {
	return a.M__mul__(other)
}

func (a String) M__imul__(other Object) Object {
	return a.M__mul__(other)
}

// Check interface is satisfied
var _ sequenceArithmetic = String("")
var _ I__len__ = String("")
var _ I__bool__ = String("")
