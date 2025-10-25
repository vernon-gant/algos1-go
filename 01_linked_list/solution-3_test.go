package linkedlist

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func makeList(values ...int) LinkedList {
    list := LinkedList{}
    for _, v := range values {
        node := Node{value: v}
        list.AddInTail(node)
    }
    return list
}

func toSlice(list *LinkedList) []int {
    var result []int
    for temp := list.head; temp != nil; temp = temp.next {
        result = append(result, temp.value)
    }
    return result
}

// FIND

func Test_GivenEmptyList_WhenFindingValue_ThenErrorReturned(t *testing.T) {
    // Given
    list := LinkedList{}

    // When
    _, err := list.Find(42)

    // Then
    assert.Error(t, err, "should return error when searching in empty list")
}

func Test_GivenSingleNodeList_WhenFindingExistingValue_ThenNodeReturned(t *testing.T) {
    // Given
    list := makeList(7)

    // When
    node, err := list.Find(7)

    // Then
    assert.NoError(t, err)
    assert.Equal(t, 7, node.value)
}

func Test_GivenSingleNodeList_WhenFindingDifferentValue_ThenErrorReturned(t *testing.T) {
    // Given
    list := makeList(7)

    // When
    _, err := list.Find(99)

    // Then
    assert.Error(t, err)
}

func Test_GivenMultiNodeList_WhenFindingExistingValue_ThenCorrectNodeReturned(t *testing.T) {
    // Given
    list := makeList(10, 20, 30)

    // When
    node, err := list.Find(20)

    // Then
    assert.NoError(t, err)
    assert.Equal(t, 20, node.value)
}

func Test_GivenMultiNodeList_WhenFindingNonexistentValue_ThenErrorReturned(t *testing.T) {
    // Given
    list := makeList(10, 20, 30)

    // When
    _, err := list.Find(99)

    // Then
    assert.Error(t, err)
}

// FIND ALL

func Test_GivenEmptyList_WhenFindingAllValues_ThenReturnsEmptySlice(t *testing.T) {
    // Given
    list := LinkedList{}

    // When
    results := list.FindAll(42)

    // Then
    assert.Empty(t, results, "FindAll should return empty slice for empty list")
}

func Test_GivenSingleNodeList_WhenFindingAllMatchingValues_ThenReturnsSingleNode(t *testing.T) {
    // Given
    list := makeList(5, 4, 3, 2, 1)

    // When
    results := list.FindAll(5)

    // Then
    assert.Len(t, results, 1)
    assert.Equal(t, 5, results[0].value)
}

func Test_GivenSingleNodeList_WhenFindingAllNonMatchingValues_ThenReturnsEmptySlice(t *testing.T) {
    // Given
    list := makeList(5)

    // When
    results := list.FindAll(99)

    // Then
    assert.Empty(t, results)
}

func Test_GivenMultiNodeList_WhenFindingAllMatchingValues_ThenReturnsAllOccurrences(t *testing.T) {
    // Given
    list := makeList(1, 2, 3, 2, 4, 2)

    // When
    results := list.FindAll(2)

    // Then
    assert.Len(t, results, 3, "should find three nodes with value 2")
    for _, n := range results {
        assert.Equal(t, 2, n.value)
    }
}

func Test_GivenMultiNodeList_WhenFindingAllNonMatchingValues_ThenReturnsEmptySlice(t *testing.T) {
    // Given
    list := makeList(10, 20, 30)

    // When
    results := list.FindAll(99)

    // Then
    assert.Empty(t, results)
}

// COUNT

func Test_GivenEmptyList_WhenObtainingCount_ThenReturnsZero(t *testing.T) {
    // Given
    list := LinkedList{}

    // When/Then
    assert.Equal(t, 0, list.Count())
}

func Test_GivenSingleElementList_WhenObtainingCount_ThenReturnsOne(t *testing.T) {
    // Given
    list := makeList(5)

    // When/Then
    assert.Equal(t, 1, list.Count())
}

func Test_GivenMultipleElementsList_WhenObtainingCount_ThenReturnsCorrectValue(t *testing.T) {
    // Given
    list := makeList(1, 2, 3, 4, 5, 6, 7)

    // When/Then
    assert.Equal(t, 7, list.Count())
}

// DELETE

func Test_GivenEmptyList_WhenDeletingValue_ThenSizeRemainsZero(t *testing.T) {
    // Given
    list := LinkedList{}

    // When
    list.Delete(42, false)

    // Then
    assert.Equal(t, 0, list.Count())
}

func Test_GivenSingleNodeList_WhenDeletingDifferentValue_ThenListUnchanged(t *testing.T) {
    // Given
    list := makeList(7)

    // When
    list.Delete(99, false)

    // Then
    assert.Equal(t, 1, list.Count())
    assert.Equal(t, []int{7}, toSlice(&list))
}

func Test_GivenSingleNodeList_WhenDeletingExistingValue_ThenListBecomesEmpty(t *testing.T) {
    // Given
    list := makeList(5)

    // When
    list.Delete(5, false)

    // Then
    assert.Equal(t, 0, list.Count())
}

func Test_GivenMultiNodeList_WhenDeletingNonexistentValue_ThenAllRemainIntact(t *testing.T) {
    // Given
    list := makeList(1, 2, 3, 4)

    // When
    list.Delete(99, false)

    // Then
    assert.Equal(t, 4, list.Count())
    assert.Equal(t, []int{1, 2, 3, 4}, toSlice(&list))
}

func Test_GivenMultiNodeListWithDuplicates_WhenDeletingValue_ThenOnlyFirstOccurrenceRemoved(t *testing.T) {
    // Given
    list := makeList(1, 2, 3, 2, 4)

    // When
    list.Delete(2, false)

    // Then
    assert.Equal(t, 4, list.Count())
    assert.Equal(t, []int{1, 3, 2, 4}, toSlice(&list))
}

func Test_GivenMultiNodeList_WhenDeletingHeadValue_ThenNextBecomesHead(t *testing.T) {
    // Given
    list := makeList(10, 20, 30)

    // When
    list.Delete(10, false)

    // Then
    assert.Equal(t, 2, list.Count())
    assert.Equal(t, []int{20, 30}, toSlice(&list))
}

func Test_GivenMultiNodeList_WhenDeletingTailValue_ThenTailMovesBack(t *testing.T) {
    // Given
    list := makeList(1, 2, 3)

    // When
    list.Delete(3, false)

    // Then
    assert.Equal(t, 2, list.Count())
    assert.Equal(t, []int{1, 2}, toSlice(&list))
}

// DELETE ALL

func Test_GivenEmptyList_WhenDeletingAllValues_ThenListRemainsEmpty(t *testing.T) {
    // Given
    list := LinkedList{}

    // When
    list.Delete(5, true)

    // Then
    assert.Equal(t, 0, list.Count())
}

func Test_GivenSingleNodeList_WhenDeletingAllMatchingValues_ThenListBecomesEmpty(t *testing.T) {
    // Given
    list := makeList(7)

    // When
    list.Delete(7, true)

    // Then
    assert.Equal(t, 0, list.Count())
}

func Test_GivenSingleNodeList_WhenDeletingAllDifferentValues_ThenListUnchanged(t *testing.T) {
    // Given
    list := makeList(7)

    // When
    list.Delete(99, true)

    // Then
    assert.Equal(t, 1, list.Count())
    assert.Equal(t, []int{7}, toSlice(&list))
}

func Test_GivenMultiNodeList_WhenDeletingAllMatchingValues_ThenListBecomesEmpty(t *testing.T) {
    // Given
    list := makeList(3, 3, 3, 3)

    // When
    list.Delete(3, true)

    // Then
    assert.Equal(t, 0, list.Count())
}

func Test_GivenMultiNodeList_WhenDeletingAllMatchingValues_ThenOnlyThoseValuesRemoved(t *testing.T) {
    // Given
    list := makeList(1, 2, 3, 2, 4, 2)

    // When
    list.Delete(2, true)

    // Then
    assert.Equal(t, 3, list.Count())
    assert.Equal(t, []int{1, 3, 4}, toSlice(&list))
}

func Test_GivenMultiNodeList_WhenDeletingAllNonexistentValues_ThenListUnchanged(t *testing.T) {
    // Given
    list := makeList(10, 20, 30)

    // When
    list.Delete(99, true)

    // Then
    assert.Equal(t, 3, list.Count())
    assert.Equal(t, []int{10, 20, 30}, toSlice(&list))
}
