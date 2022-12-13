package ecore

import "github.com/masagroup/soft.go/diff"

type diffToListVisitor struct {
	list EList
}

func (d *diffToListVisitor) HandleAdd(index int, element any) {
	d.list.Insert(index, element)
}

func (d *diffToListVisitor) HandleMove(oldIndex int, newIndex int, _ any) {
	d.list.Move(oldIndex, newIndex)
}

func (d *diffToListVisitor) HandleRemove(index int, _ any) {
	d.list.Remove(index)
}

func (d *diffToListVisitor) HandleReplace(index int, _ any, newElement any) {
	d.list.Set(index, newElement)
}

func ApplyDiffToList(diff diff.Diff[any], list EList) {
	diff.Accept(&diffToListVisitor{list: list})
}
