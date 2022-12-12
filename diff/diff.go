package diff

import (
	"github.com/zyedidia/generic"
	"github.com/zyedidia/generic/stack"
)

//
// https://onmyway133.com/posts/diff-algorithm/
// Paul Heckel's Diff Algorithm
// IGListDiff algorithm
//

type Operation any
type Insert int
type Delete int
type Update int
type Move struct {
	From int
	To   int
}

type Result struct {
	Operations  []Operation
	OldIndexFor map[uint64]int
	NewIndexFor map[uint64]int
}

type entry struct {
	oldCounter int
	newCounter int
	oldIndexes stack.Stack[int]
	updated    bool
}

type record struct {
	entry *entry
	index int
}

func newRecord(entry *entry) *record {
	return &record{
		entry: entry,
		index: -1,
	}
}

func Compute[T any](oldArray, newArray []T, equals generic.EqualsFn[T], hash generic.HashFn[T]) Result {
	table := map[uint64]*entry{}
	oldCount := len(oldArray)
	newCount := len(newArray)

	// pass 1
	// create an entry for every item in the new array
	// increment its new count for each occurence
	// record `nil` for each occurence of the item in the new array
	newRecords := []*record{}
	for _, k := range newArray {
		h := hash(k)
		e, isEntry := table[h]
		if !isEntry {
			e = &entry{}
			table[h] = e
		}
		e.newCounter++
		e.oldIndexes.Push(-1)
		newRecords = append(newRecords, newRecord(e))
	}

	// pass 2
	// update or create an entry for every item in the old array
	// increment its old count for each occurence
	// record the original index of the item in the old array
	// MUST be done in descending order to respect the oldIndexes stack construction
	oldRecords := []*record{}
	for i := len(oldArray) - 1; i >= 0; i-- {
		k := oldArray[i]
		h := hash(k)
		e, isEntry := table[h]
		if !isEntry {
			e = &entry{}
			table[h] = e
		}
		e.oldCounter++
		e.oldIndexes.Push(i)
		oldRecords = append(oldRecords, newRecord(e))
	}

	// pass 3
	// handle data that occurs in both arrays
	for i, r := range newRecords {
		e := r.entry
		if oldIndex := e.oldIndexes.Pop(); oldIndex != -1 {
			if oldIndex < oldCount {
				n := newArray[i]
				o := oldArray[oldIndex]
				if !equals(n, o) {
					e.updated = true
				}
			}
			if e.newCounter > 0 && e.oldCounter > 0 {
				newRecords[i].index = oldIndex
				oldRecords[oldIndex].index = i
			}
		}
	}

	// storage for final results
	operations := []Operation{}
	oldIndexFor := map[uint64]int{}
	newIndexFor := map[uint64]int{}

	// track offsets from deleted items to calculate where items have moved
	// iterate old array records checking for deletes
	// increment offset for each delete
	runningOffset := 0
	deleteOffsets := make([]int, oldCount)
	for i, r := range oldRecords {
		deleteOffsets[i] = runningOffset
		if r.index == -1 {
			operations = append(operations, Delete(i))
			runningOffset++
		}
		oldIndexFor[hash(oldArray[i])] = i
	}

	//reset and track offsets from inserted items to calculate where items have moved
	runningOffset = 0
	insertOffsets := make([]int, newCount)
	for i, r := range newRecords {
		insertOffsets[i] = runningOffset
		if oldIndex := r.index; oldIndex != -1 {
			// note that an entry can be updated /and/ moved
			if r.entry.updated {
				operations = append(operations, Update(oldIndex))
			}
			// calculate the offset and determine if there was a move
			// if the indexes match, ignore the index
			deleteOffset := deleteOffsets[oldIndex]
			if oldIndex-deleteOffset+runningOffset != i && oldIndex != i {
				operations = append(operations, Move{From: oldIndex, To: i})
			}

		} else {
			// add to inserts if the opposing index is -1
			operations = append(operations, Insert(i))
			runningOffset++
		}
		newIndexFor[hash(newArray[i])] = i
	}

	return Result{
		Operations:  operations,
		OldIndexFor: oldIndexFor,
		NewIndexFor: newIndexFor,
	}
}

func ComputeOrdered[T any](oldArray, newArray []T, equals generic.EqualsFn[T], hash generic.HashFn[T]) Result {
	result := Compute[T](oldArray, newArray, equals, hash)
	insertions := []Operation{}
	updates := []Operation{}
	deletions := make([]Operation, len(oldArray))
	trackDeletion := func(index int, op Operation) {
		if deletions[index] == nil {
			deletions[index] = op
		}
	}

	for _, operation := range result.Operations {
		switch op := operation.(type) {
		case Insert:
			insertions = append(insertions, operation)
		case Delete:
			trackDeletion(int(op), operation)
		case Move:
			insertions = append(insertions, op.To)
			trackDeletion(op.From, Delete(op.From))
		case Update:
			updates = append(updates, op)
		}

	}

	ordered := []Operation{}
	for i := len(deletions) - 1; i >= 0; i-- {
		if d := deletions[i]; d != nil {
			ordered = append(ordered, d)
		}
	}
	ordered = append(ordered, insertions...)
	ordered = append(ordered, updates...)
	return Result{
		Operations:  ordered,
		OldIndexFor: result.OldIndexFor,
		NewIndexFor: result.NewIndexFor,
	}
}
