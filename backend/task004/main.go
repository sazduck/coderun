package task004

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	tree := NewBST(r)

	buf := bufio.NewWriterSize(w, 1<<20)
	defer buf.Flush()

	h := tree.Height()
	buf.WriteString(strconv.Itoa(h))
	buf.WriteByte('\n')
	return nil
}

type BST struct {
	Root *Node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewBST(r io.Reader) *BST {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	tree := &BST{}
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		if i == 0 {
			break
		}
		tree.Insert(i)
	}

	return tree
}

func (t *BST) Insert(val int) {
	if t.Root == nil {
		t.Root = &Node{Val: val}
		return
	}

	curr := t.Root
	for {
		switch {
		case val < curr.Val:
			if curr.Left == nil {
				curr.Left = &Node{Val: val}
				return
			}
			curr = curr.Left
		case val > curr.Val:
			if curr.Right == nil {
				curr.Right = &Node{Val: val}
				return
			}
			curr = curr.Right
		default:
			return
		}
	}
}

func (t *BST) Height() int {
	if t.Root == nil {
		return 0
	}
	return t.Root.Height()
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return max(n.Left.Height(), n.Right.Height()) + 1

}

func (n *Node) WriteInOrder(w *bufio.Writer) {
	if n == nil {
		return
	}
	n.Left.WriteInOrder(w) // <

	w.WriteString(strconv.Itoa(n.Val)) // cur
	w.WriteByte(' ')

	n.Right.WriteInOrder(w) // >
}
