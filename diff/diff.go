package diff

import (
	"github.com/zyedidia/generic"
	"github.com/zyedidia/generic/set"
	"github.com/zyedidia/generic/stack"
)

//
// https://onmyway133.com/posts/diff-algorithm/
// Paul Heckel's Diff Algorithm
// IGListDiff algorithm
//

type MoveIndex struct {
	From int
	To   int
}

func makeMoveIndex(from, to int) MoveIndex {
	min := from
	max := to
	if to < from {
		min = to
		max = from
	}
	return MoveIndex{From: min, To: max}
}

func combineHashCode(hashCodes ...uint64) uint64 {
	var constant uint64 = 23
	var result uint64 = 17
	for _, hashCode := range hashCodes {
		result = result*constant + hashCode
	}
	return result
}

func newMoveIndexSet(in ...MoveIndex) set.Set[MoveIndex] {
	return set.NewHashset[MoveIndex](
		0,
		func(a, b MoveIndex) bool {
			return a.From == b.From && a.To == b.To
		},
		func(mi MoveIndex) uint64 {
			return combineHashCode(generic.HashInt(mi.From), generic.HashInt(mi.To))
		},
		in...,
	)
}

type Result interface {
	GetDeletes() set.Set[int]
	GetInserts() set.Set[int]
	GetUpdates() set.Set[int]
	GetMoves() set.Set[MoveIndex]
	GetOldIndexFor(uint64) (int, bool)
	GetNewIndexFor(uint64) (int, bool)
}

type resultImpl struct {
	deletes     set.Set[int]
	inserts     set.Set[int]
	updates     set.Set[int]
	moves       set.Set[MoveIndex]
	oldIndexFor map[uint64]int
	newIndexFor map[uint64]int
}

func newResultImpl() *resultImpl {
	return &resultImpl{
		deletes:     set.NewMapset[int](),
		inserts:     set.NewMapset[int](),
		updates:     set.NewMapset[int](),
		moves:       newMoveIndexSet(),
		oldIndexFor: map[uint64]int{},
		newIndexFor: map[uint64]int{},
	}
}

func (r *resultImpl) GetDeletes() set.Set[int] {
	return r.deletes
}

func (r *resultImpl) GetInserts() set.Set[int] {
	return r.inserts
}

func (r *resultImpl) GetUpdates() set.Set[int] {
	return r.updates
}

func (r *resultImpl) GetMoves() set.Set[MoveIndex] {
	return r.moves
}

func (r *resultImpl) GetOldIndexFor(h uint64) (i int, b bool) {
	i, b = r.oldIndexFor[h]
	return
}

func (r *resultImpl) GetNewIndexFor(h uint64) (i int, b bool) {
	i, b = r.newIndexFor[h]
	return
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
	result := newResultImpl()

	// track offsets from deleted items to calculate where items have moved
	// iterate old array records checking for deletes
	// increment offset for each delete
	runningOffset := 0
	deleteOffsets := make([]int, oldCount)
	for i, r := range oldRecords {
		deleteOffsets[i] = runningOffset
		if r.index == -1 {
			result.deletes.Put(i)
			runningOffset++
		}
		result.oldIndexFor[hash(oldArray[i])] = i
	}

	//reset and track offsets from inserted items to calculate where items have moved
	runningOffset = 0
	insertOffsets := make([]int, newCount)
	for i, r := range newRecords {
		insertOffsets[i] = runningOffset
		if oldIndex := r.index; oldIndex != -1 {
			// note that an entry can be updated /and/ moved
			if r.entry.updated {
				result.updates.Put(oldIndex)
			}
			// calculate the offset and determine if there was a move
			// if the indexes match, ignore the index
			deleteOffset := deleteOffsets[oldIndex]
			if oldIndex-deleteOffset+runningOffset != i && oldIndex != i {
				result.moves.Put(makeMoveIndex(oldIndex, i))
			}

		} else {
			// add to inserts if the opposing index is -1
			result.inserts.Put(i)
			runningOffset++
		}
		result.newIndexFor[hash(newArray[i])] = i
	}

	return result
}
