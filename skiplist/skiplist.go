package skiplist

import (
	"github.com/slowmoon/leveldb/comp"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type SkipList struct {
	Level int
	head *Node
	lock sync.RWMutex
	comp comp.Comparator
	length int
}

func New(comp comp.Comparator, level int)*SkipList  {
    sk := new(SkipList)
    sk.Level = level
    sk.comp = comp
    sk.head =  NewNode(nil, nil, level)

    return  sk
}

func (s *SkipList)randomLevel()int  {
      level  := 1
     for ; rand.Int()%2==1 && level < s.Level;level++{}
     return  level
}

func (s *SkipList)Insert(key ,value interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()

	level := s.randomLevel()
    newNode := NewNode(key, value, level)

    node := s.head
    insert := false
    for i:= level-1;i >=0 ;i-- {
    	//insert the data
    	for {
    	   cur := node.getNext(i)
    	   if cur == nil || s.comp.Compare(cur.key, key) > 0 {
    	      //find one insert it here
               newNode.setNext(i, cur)
               node.setNext(i, newNode)
               insert = true
               break
		   }else if s.comp.Compare(cur.key, key) ==0 {
		         //equal update value
			   cur.value = value
			   return
		   }else {
		      node = cur
		   }
		}
	}
    if insert {
    	s.length++
	}
}

func (s *SkipList)findGreaterOrEqual(key interface{})(*Node, []*Node) {
    level := s.Level
    node := s.head
    founds := make([]*Node, level)
    for {
        node1:= node.getNext(level)
        if s.KeyIsAfterNode(key, node1) {
        	node = node1
        	break
		} else  {
			//less or equal
			founds[level-1] =  node1
			if level == 1 {
				return  node1, founds
			}else {
				level--
			}
		}
	}
    return  nil ,founds
}

func (s *SkipList)Contains(key interface{})bool {
    s.lock.RLock()
    defer  s.lock.RUnlock()
	for node := s.head.getNext(0);node!= nil ; node = node.getNext(0) {
	     if s.comp.Compare(node.key, key) ==0 {
	     	return  true
		 }
	}
	return  false
}

func (s *SkipList)KeyIsAfterNode(key interface{}, node *Node)bool  {
   return  (node!= nil)	&& s.comp.Compare(node.key, key) <0
}

func (s *SkipList)Iterator() *Iterator {
     return  &Iterator{
     	sk: s,
	 }
}
