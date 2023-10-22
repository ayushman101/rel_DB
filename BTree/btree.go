package BTree

type BNode struct {
	data []byte //can be dumped on the disk easily.
}


const (
	BNODE_NODE=1
	BNODE_LEAF=2
)

type BTree struct {
	//pointer( not an in-memory pointer)
	//a 64 bit integer referencing a page on the disk
	root uint64

	//callbacks for managing on disk pages

	get func(uint64) BNode  // dereference a pointer
	
	new func(BNode) uint64 // allocate a new page

	del func(uint64)    // deallocate a page
}


