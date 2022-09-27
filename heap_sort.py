def heapSort(array):
    buildMaxHeap(array)
    # start with the ending index all the way to 0
    for endIdx in range(len(array)-1, 0, -1):
        swap(0, endIdx, array)
        # reduce heap size by 1
        # sift down the value we just swapped
        heapify(0, endIdx - 1, array)
    return array

def buildMaxHeap(array):
    lastNonLeafNodeIdx = (len(array) - 2) // 2
    for currentIdx in range(lastNonLeafNodeIdx, -1, -1):
        heapify(currentIdx, len(array)-1, array)

# sift down
def heapify(currentIdx, endIdx, array):
    leftChildIdx = 2 * currentIdx + 1
    while leftChildIdx <= endIdx:
        rightChildIdx = 2 * currentIdx + 2 if 2 * currentIdx + 2 <= endIdx else -1
        largestChildIdx = leftChildIdx
        if rightChildIdx != -1 and array[leftChildIdx] < array[rightChildIdx]:
            largestChildIdx = rightChildIdx
        
        if array[currentIdx] < array[largestChildIdx]:
            swap(currentIdx, largestChildIdx, array)
            currentIdx = largestChildIdx
            leftChildIdx = 2 * currentIdx + 1
        else:
            return

def swap(i, j, array):
    array[i], array[j] = array[j], array[i]


if __name__ == "__main__":
    array = [1, -15, 22, 40, 9, 91]
    print(heapSort(array))

# output: [-15, 1, 9, 22, 40, 91]