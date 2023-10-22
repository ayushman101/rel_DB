package BTree

import(
	"os"
	"log"
	"binary"
)

//  | type | nkeys | pointers | offsets | key-values
//	2B |   2B  | nkeys*8B | nkeys*2B| ...

// | klen  | vlen  | key  | value
// | 2B    | 2B    | ...  | ... 
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


const HEADER=4 // node type: 2B,  nKeys: 2B

const BTREE_PAGE_SIZE=4096

const BTREE_MAX_KEY_SIZE=1000

const BTREE_MAX_VAL_SIZE=3000

func init(){
	// 8=> size of each pointer, 2=> size of offset , 4=> klen and vlen
	node1max:= HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE

	if node1max> BTREE_PAGE_SIZE{
			log.Fatal(" node size greater than MAX PAGE SIZE")
			os.Exit(1)
		}
}


//header of BNode

func (node BNode) btype() uint16{
	return binary.LittleEndian.Uint16(node.data)
}

func (node BNode) nkeys() uint16{

	return binary.LittleEndian.Uint16(node.data[2:4])
}


func (node BNode) setHeader(btype uint16, nkeys uint16){
	binary.LittleEndian.PutUint16(node.data[0:2],btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

//pointers

func (node BNode) getPtr(idx uint16) uint64 {
	if idx >= node.nkeys(){
		log.Fatal("Index out of bounds")
		os.Exit(1)
	}

	pos:= HEADER + 8*idx

	return binary.LittleEndian.Uint64(node.data[pos:])
}

func (node BNode) setPtr(idx uint16, val uint64) {
	if idx >= node.nkeys() {
		log.Fatal("Index out of bounds")
		os.Exit(1)
	}
	
	
	pos:= HEADER + 8*idx

	binary.LittleEndian.PutUint64(node.data[pos:],val)

}


