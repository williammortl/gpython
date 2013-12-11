// Tuple objects

package py

var TupleType = ObjectType.NewType("tuple", "tuple() -> empty tuple\ntuple(iterable) -> tuple initialized from iterable's items\n\nIf the argument is a tuple, the return value is the same object.", TupleNew, nil)

type Tuple []Object

// Type of this Tuple object
func (o Tuple) Type() *Type {
	return TupleType
}

// TupleNew
func TupleNew(metatype *Type, args Tuple, kwargs StringDict) (res Object) {
	var iterable Object
	UnpackTuple(args, kwargs, "tuple", 0, 1, &iterable)
	if iterable != nil {
		return SequenceTuple(iterable)
	}
	return Tuple{}
}

// Copy a tuple object
func (t Tuple) Copy() Tuple {
	newT := make(Tuple, len(t))
	copy(newT, t)
	return newT
}
func (t Tuple) M__len__() Object {
	return Int(len(t))
}

func (t Tuple) M__bool__() Object {
	return NewBool(len(t) > 0)
}

func (t Tuple) M__iter__() Object {
	return NewIterator(t)
}

func (t Tuple) M__getitem__(key Object) Object {
	i := IndexIntCheck(key, len(t))
	return t[i]
}

func (a Tuple) M__add__(other Object) Object {
	if b, ok := other.(Tuple); ok {
		newTuple := make(Tuple, len(a)+len(b))
		copy(newTuple, a)
		copy(newTuple[len(b):], b)
		return newTuple
	}
	return NotImplemented
}

func (a Tuple) M__radd__(other Object) Object {
	if b, ok := other.(Tuple); ok {
		return b.M__add__(a)
	}
	return NotImplemented
}

func (a Tuple) M__iadd__(other Object) Object {
	return a.M__add__(other)
}

func (l Tuple) M__mul__(other Object) Object {
	if b, ok := convertToInt(other); ok {
		m := len(l)
		n := int(b) * m
		newTuple := make(Tuple, n)
		for i := 0; i < n; i += m {
			copy(newTuple[i:i+m], l)
		}
		return newTuple
	}
	return NotImplemented
}

func (a Tuple) M__rmul__(other Object) Object {
	return a.M__mul__(other)
}

func (a Tuple) M__imul__(other Object) Object {
	return a.M__mul__(other)
}

// Check interface is satisfied
var _ sequenceArithmetic = Tuple(nil)
var _ I__len__ = Tuple(nil)
var _ I__bool__ = Tuple(nil)
var _ I__iter__ = Tuple(nil)
var _ I__getitem__ = Tuple(nil)

// var _ richComparison = Tuple(nil)
