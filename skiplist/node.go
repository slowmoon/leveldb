package skiplist

type Node struct {
	key interface{}
	value interface{}
	next []*Node
}

func NewNode(key , value interface{}, height int)*Node  {

	return  &Node{
		key: key,
		value: value,
		next: make([]*Node, height, height),
	}
}

func (n *Node)getNext(level int)*Node  {
    return  n.next[level-1]
}

func (n *Node)setNext(level int , x *Node)  {
    n.next[level-1] = x
}

