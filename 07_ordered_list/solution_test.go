package ordered_list

import (
	"testing"
	"constraints"
	"github.com/stretchr/testify/assert"
)

// helpers

func makeAscList(values ...int) OrderedList[int] {
	list := OrderedList[int]{_ascending: true}
	for _, v := range values {
		list.Add(v)
	}
	return list
}

func makeDescList(values ...int) OrderedList[int] {
	list := OrderedList[int]{_ascending: false}
	for _, v := range values {
		list.Add(v)
	}
	return list
}

func toSlice[T constraints.Ordered](list *OrderedList[T]) []T {
	var result []T
	for node := list.head; node != nil; node = node.next {
		result = append(result, node.value)
	}
	return result
}

// COUNT

func Test_GivenEmptyList_WhenGettingCount_ThenReturnsZero(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: true}

	// When
	count := list.Count()

	// Then
	assert.Equal(t, 0, count)
}

func Test_GivenListWithElements_WhenGettingCount_ThenReturnsCorrectNumber(t *testing.T) {
	// Given
	list := makeAscList(5, 2, 8, 1)

	// When
	count := list.Count()

	// Then
	assert.Equal(t, 4, count)
}

// ADD - ASCENDING ORDER

func Test_GivenEmptyAscendingList_WhenAddingElement_ThenBecomesHeadAndTail(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: true}

	// When
	list.Add(5)

	// Then
	assert.Equal(t, 1, list.Count())
	assert.Equal(t, []int{5}, toSlice(&list))
	assert.Equal(t, 5, list.head.value)
	assert.Equal(t, 5, list.tail.value)
}

func Test_GivenAscendingList_WhenAddingSmallerElement_ThenInsertsAtBeginning(t *testing.T) {
	// Given
	list := makeAscList(5, 10)

	// When
	list.Add(2)

	// Then
	assert.Equal(t, []int{2, 5, 10}, toSlice(&list))
	assert.Equal(t, 2, list.head.value)
}

func Test_GivenAscendingList_WhenAddingLargerElement_ThenInsertsAtEnd(t *testing.T) {
	// Given
	list := makeAscList(5, 10)

	// When
	list.Add(15)

	// Then
	assert.Equal(t, []int{5, 10, 15}, toSlice(&list))
	assert.Equal(t, 15, list.tail.value)
}

func Test_GivenAscendingList_WhenAddingMiddleElement_ThenInsertsInCorrectPosition(t *testing.T) {
	// Given
	list := makeAscList(5, 15)

	// When
	list.Add(10)

	// Then
	assert.Equal(t, []int{5, 10, 15}, toSlice(&list))
}

func Test_GivenAscendingList_WhenAddingMultipleElements_ThenMaintainsOrder(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: true}

	// When - add in random order
	list.Add(8)
	list.Add(3)
	list.Add(10)
	list.Add(1)
	list.Add(6)

	// Then - should be sorted
	assert.Equal(t, []int{1, 3, 6, 8, 10}, toSlice(&list))
}

func Test_GivenAscendingList_WhenAddingDuplicates_ThenMaintainsAllCopies(t *testing.T) {
	// Given
	list := makeAscList(5, 10)

	// When
	list.Add(5)
	list.Add(10)

	// Then
	assert.Equal(t, []int{5, 5, 10, 10}, toSlice(&list))
	assert.Equal(t, 4, list.Count())
}

// ADD - DESCENDING ORDER

func Test_GivenEmptyDescendingList_WhenAddingElement_ThenBecomesHeadAndTail(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: false}

	// When
	list.Add(5)

	// Then
	assert.Equal(t, 1, list.Count())
	assert.Equal(t, []int{5}, toSlice(&list))
}

func Test_GivenDescendingList_WhenAddingLargerElement_ThenInsertsAtBeginning(t *testing.T) {
	// Given
	list := makeDescList(10, 5)

	// When
	list.Add(15)

	// Then
	assert.Equal(t, []int{15, 10, 5}, toSlice(&list))
	assert.Equal(t, 15, list.head.value)
}

func Test_GivenDescendingList_WhenAddingSmallerElement_ThenInsertsAtEnd(t *testing.T) {
	// Given
	list := makeDescList(10, 5)

	// When
	list.Add(2)

	// Then
	assert.Equal(t, []int{10, 5, 2}, toSlice(&list))
	assert.Equal(t, 2, list.tail.value)
}

func Test_GivenDescendingList_WhenAddingMiddleElement_ThenInsertsInCorrectPosition(t *testing.T) {
	// Given
	list := makeDescList(15, 5)

	// When
	list.Add(10)

	// Then
	assert.Equal(t, []int{15, 10, 5}, toSlice(&list))
}

func Test_GivenDescendingList_WhenAddingMultipleElements_ThenMaintainsOrder(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: false}

	// When - add in random order
	list.Add(3)
	list.Add(10)
	list.Add(1)
	list.Add(8)
	list.Add(6)

	// Then - should be reverse sorted
	assert.Equal(t, []int{10, 8, 6, 3, 1}, toSlice(&list))
}

// FIND - ASCENDING ORDER

func Test_GivenEmptyList_WhenFinding_ThenReturnsError(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: true}

	// When
	_, err := list.Find(5)

	// Then
	assert.Error(t, err)
}

func Test_GivenAscendingList_WhenFindingExistingElement_ThenReturnsNode(t *testing.T) {
	// Given
	list := makeAscList(1, 5, 10, 15)

	// When
	node, err := list.Find(10)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 10, node.value)
}

func Test_GivenAscendingList_WhenFindingNonexistent_ThenReturnsError(t *testing.T) {
	// Given
	list := makeAscList(1, 5, 10, 15)

	// When
	_, err := list.Find(7)

	// Then
	assert.Error(t, err)
}

func Test_GivenAscendingList_WhenFindingValueSmallerThanAll_ThenReturnsErrorEarly(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 15, 20, 25)

	// When
	_, err := list.Find(2)

	// Then - should terminate early when encountering 5 > 2
	assert.Error(t, err)
}

func Test_GivenAscendingList_WhenFindingValueLargerThanAll_ThenReturnsErrorAfterScan(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 15)

	// When
	_, err := list.Find(20)

	// Then
	assert.Error(t, err)
}

func Test_GivenAscendingListWithDuplicates_WhenFinding_ThenReturnsFirstOccurrence(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 10, 10, 15)

	// When
	node, err := list.Find(10)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 10, node.value)
}

// FIND - DESCENDING ORDER

func Test_GivenDescendingList_WhenFindingExistingElement_ThenReturnsNode(t *testing.T) {
	// Given
	list := makeDescList(15, 10, 5, 1)

	// When
	node, err := list.Find(10)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 10, node.value)
}

func Test_GivenDescendingList_WhenFindingNonexistent_ThenReturnsError(t *testing.T) {
	// Given
	list := makeDescList(15, 10, 5, 1)

	// When
	_, err := list.Find(7)

	// Then
	assert.Error(t, err)
}

func Test_GivenDescendingList_WhenFindingValueLargerThanAll_ThenReturnsErrorEarly(t *testing.T) {
	// Given
	list := makeDescList(20, 15, 10, 5)

	// When
	_, err := list.Find(25)

	// Then - should terminate early when encountering 20 < 25
	assert.Error(t, err)
}

// DELETE - ASCENDING ORDER

func Test_GivenEmptyList_WhenDeleting_ThenNothingHappens(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: true}

	// When
	list.Delete(5)

	// Then
	assert.Equal(t, 0, list.Count())
}

func Test_GivenSingleElementList_WhenDeletingThatElement_ThenBecomesEmpty(t *testing.T) {
	// Given
	list := makeAscList(5)

	// When
	list.Delete(5)

	// Then
	assert.Equal(t, 0, list.Count())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func Test_GivenAscendingList_WhenDeletingHead_ThenHeadUpdates(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 15)

	// When
	list.Delete(5)

	// Then
	assert.Equal(t, []int{10, 15}, toSlice(&list))
	assert.Equal(t, 10, list.head.value)
}

func Test_GivenAscendingList_WhenDeletingTail_ThenTailUpdates(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 15)

	// When
	list.Delete(15)

	// Then
	assert.Equal(t, []int{5, 10}, toSlice(&list))
	assert.Equal(t, 10, list.tail.value)
}

func Test_GivenAscendingList_WhenDeletingMiddle_ThenMaintainsOrder(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 15)

	// When
	list.Delete(10)

	// Then
	assert.Equal(t, []int{5, 15}, toSlice(&list))
}

func Test_GivenAscendingListWithDuplicates_WhenDeleting_ThenRemovesOnlyFirstOccurrence(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 10, 10, 15)

	// When
	list.Delete(10)

	// Then
	assert.Equal(t, []int{5, 10, 10, 15}, toSlice(&list))
	assert.Equal(t, 4, list.Count())
}

func Test_GivenAscendingList_WhenDeletingNonexistent_ThenUnchanged(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 15)

	// When
	list.Delete(7)

	// Then
	assert.Equal(t, []int{5, 10, 15}, toSlice(&list))
	assert.Equal(t, 3, list.Count())
}

// DELETE - DESCENDING ORDER

func Test_GivenDescendingList_WhenDeletingHead_ThenHeadUpdates(t *testing.T) {
	// Given
	list := makeDescList(15, 10, 5)

	// When
	list.Delete(15)

	// Then
	assert.Equal(t, []int{10, 5}, toSlice(&list))
	assert.Equal(t, 10, list.head.value)
}

func Test_GivenDescendingList_WhenDeletingMiddle_ThenMaintainsOrder(t *testing.T) {
	// Given
	list := makeDescList(15, 10, 5)

	// When
	list.Delete(10)

	// Then
	assert.Equal(t, []int{15, 5}, toSlice(&list))
}

// CLEAR

func Test_GivenNonEmptyList_WhenClearing_ThenBecomesEmpty(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 15)

	// When
	list.Clear(true)

	// Then
	assert.Equal(t, 0, list.Count())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func Test_GivenAscendingList_WhenClearingWithDescending_ThenChangesOrder(t *testing.T) {
	// Given
	list := makeAscList(5, 10, 15)

	// When
	list.Clear(false)
	list.Add(10)
	list.Add(5)
	list.Add(15)

	// Then - should now be descending
	assert.Equal(t, []int{15, 10, 5}, toSlice(&list))
}

// COMPARE

func Test_GivenTwoIntegers_WhenComparing_ThenReturnsCorrectResult(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: true}

	// When/Then
	assert.Equal(t, -1, list.Compare(5, 10))
	assert.Equal(t, 0, list.Compare(5, 5))
	assert.Equal(t, 1, list.Compare(10, 5))
}

// COMPLEX SCENARIOS

func Test_GivenList_WhenAlternatingAddAndDelete_ThenMaintainsOrder(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: true}

	// When
	list.Add(10)
	list.Add(5)
	list.Add(15)
	list.Delete(5)
	list.Add(12)
	list.Add(3)
	list.Delete(15)

	// Then
	assert.Equal(t, []int{3, 10, 12}, toSlice(&list))
}

func Test_GivenDescendingList_WhenMultipleOperations_ThenMaintainsOrder(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: false}

	// When
	list.Add(5)
	list.Add(15)
	list.Add(10)
	list.Delete(15)
	list.Add(20)
	list.Add(8)

	// Then
	assert.Equal(t, []int{20, 10, 8, 5}, toSlice(&list))
}

// TASK 8: REMOVE DUPLICATES

func Test_GivenEmptyList_WhenRemovingDuplicates_ThenRemainsEmpty(t *testing.T) {
	// Given
	list := OrderedList[int]{_ascending: true}

	// When
	list.RemoveDuplicates()

	// Then
	assert.Equal(t, 0, list.Count())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func Test_GivenSingleElementList_WhenRemovingDuplicates_ThenUnchanged(t *testing.T) {
	// Given
	list := makeAscList(5)

	// When
	list.RemoveDuplicates()

	// Then
	assert.Equal(t, []int{5}, toSlice(&list))
	assert.Equal(t, 1, list.Count())
}

func Test_GivenListWithNoDuplicates_WhenRemovingDuplicates_ThenUnchanged(t *testing.T) {
	// Given
	list := makeAscList(1, 2, 3, 4, 5)

	// When
	list.RemoveDuplicates()

	// Then
	assert.Equal(t, []int{1, 2, 3, 4, 5}, toSlice(&list))
	assert.Equal(t, 5, list.Count())
}

func Test_GivenAscendingListWithDuplicates_WhenRemovingDuplicates_ThenKeepsOnlyUnique(t *testing.T) {
	// Given
	list := makeAscList(1, 2, 2, 3, 3, 3, 4, 5, 5)

	// When
	list.RemoveDuplicates()

	// Then
	assert.Equal(t, []int{1, 2, 3, 4, 5}, toSlice(&list))
	assert.Equal(t, 5, list.Count())
}

func Test_GivenDescendingListWithDuplicates_WhenRemovingDuplicates_ThenKeepsOnlyUnique(t *testing.T) {
	// Given
	list := makeDescList(5, 5, 4, 3, 3, 3, 2, 2, 1)

	// When
	list.RemoveDuplicates()

	// Then
	assert.Equal(t, []int{5, 4, 3, 2, 1}, toSlice(&list))
	assert.Equal(t, 5, list.Count())
}

func Test_GivenListWithAllDuplicates_WhenRemovingDuplicates_ThenKeepsOnlyOne(t *testing.T) {
	// Given
	list := makeAscList(7, 7, 7, 7, 7)

	// When
	list.RemoveDuplicates()

	// Then
	assert.Equal(t, []int{7}, toSlice(&list))
	assert.Equal(t, 1, list.Count())
}

func Test_GivenListWithDuplicatesAtHeadAndTail_WhenRemovingDuplicates_ThenUpdatesHeadAndTail(t *testing.T) {
	// Given
	list := makeAscList(1, 1, 1, 2, 3, 5, 5, 5)

	// When
	list.RemoveDuplicates()

	// Then
	assert.Equal(t, []int{1, 2, 3, 5}, toSlice(&list))
	assert.Equal(t, 1, list.head.value)
	assert.Equal(t, 5, list.tail.value)
}

func Test_GivenListWithConsecutiveDuplicates_WhenRemovingDuplicates_ThenRemovesAll(t *testing.T) {
	// Given
	list := makeAscList(1, 2, 2, 2, 2, 3, 4, 4, 5, 5, 5, 5, 5)

	// When
	list.RemoveDuplicates()

	// Then
	assert.Equal(t, []int{1, 2, 3, 4, 5}, toSlice(&list))
	assert.Equal(t, 5, list.Count())
}