package comp

type Comparator interface {

	Compare(a, b interface{}) int
}

type ComparatorFunc func(a, b interface{})int

func (c ComparatorFunc)Compare(a, b interface{})int  {
     return  c(a, b)
}

var IntComparator ComparatorFunc  = func(a, b interface{}) int {
	aInt, bInt := a.(int), b.(int)
	return  aInt - bInt
}



