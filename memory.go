package turtl

import (
	"fmt"
)

/*
GC is the garbage collector
*/
type GC struct {
	/*
		heap is a memory store managed by the GC. Elements in
		heap are of type *Object, which wrap data references
		and member functions and into a single struct.
	*/
	heap *heap
	/*
		pmap is a map from identifiers to pointers to Objects
		on the heap. Identifiers are uint64 hash values of their
		names. The current hash function is xxHash. For more
		info see http://cyan4973.github.io/xxHash/
	*/
	pmap map[uint64]*Object
}

type pointer struct {
	address uint
	size    uint
}

/*
Collect the garbage
*/
func (gc *GC) Collect() {

}

/*
Allocate some memory
*/
func (gc *GC) Allocate(x []byte) {

}

/*
LookupID returns an Object pointer from the heap given an
identifier
*/
func (gc *GC) LookupID(id uint64) (*Object, error) {
	var obj *Object
	var ok bool
	if obj, ok = gc.pmap[id]; !ok {
		return nil, unknownIDErr{ID: id}
	}
	return obj, nil
}

/*
NumObjects currently held by the GC
*/
func (gc *GC) NumObjects() int {
	return len(gc.pmap)
}

/*
Object wraps data and functions into a element which
is tracked by the GC
*/
type Object struct {
	address   uint64           // location on the heap
	dataPtr   uint64           // location of data segment
	dataSize  uint64           // size of data segment
	funcTable map[uint8]uint64 // member function pointer table
}

/*
RunMethod of this Object given a method ID

x := RunMethod()
*/
func (obj *Object) RunMethod(mthID uint16, args []int) error {
	return nil
}

/*
dataBuffer is the main storage type for managed
data, including ints, floats and strings which
belong to Objects owned by the GC
*/
type dataBuffer []byte

/*
heap is a managed memory store owned by a GC
*/
type heap struct {
	buff []Object
}

const (
	initialHeapSize uint32 = 12000000  // 12 Mb
	maximumHeapSize uint32 = 256000000 // 256 Mb
)

/*
NewHeap returns a new GC collected heap. It's
initialized to ${initialHeapSize}
*/
func newHeap() heap {
	buff := make([]Object, 0, initialHeapSize)
	heap := heap{
		buff: buff,
	}
	return heap
}

/*
grow the heap by x number of bytes
*/
func (heap *heap) grow(x int) error {
	if len(heap.buff)+x > int(maximumHeapSize) {
		return &heapGrowErr{
			HeapSize: len(heap.buff),
			GrowSize: x,
		}
	}
	newBuff := make([]Object, 0, len(heap.buff)+x)
	for i, b := range heap.buff {
		newBuff[i] = b
	}
	heap.buff = newBuff
	return nil
}

/*
Shrink the heap by x number of bytes
*/
func (heap *heap) shrink(x int) error {
	return nil
}

/*
heapGrowErr reports errors due to memory
constraints when trying to grow the heap
*/
type heapGrowErr struct {
	HeapSize int
	GrowSize int
}

func (he *heapGrowErr) Error() string {
	return fmt.Sprintf("[heap]: current size - %d ", he.HeapSize)
}

/*
unknownIDErr reports pointer lookup errors.
This means the supplied identifier was not
found in the GC's pmap
*/
type unknownIDErr struct {
	ID uint64
}

func (ukID unknownIDErr) Error() string {
	return fmt.Sprintf("[gc]: unknown identifier: %d", ukID.ID)
}
