def quick_sort(array, start, end):
    if start < end:
        pivotIndex = partition(array, start, end)
        quick_sort(array, start, pivotIndex - 1)
        quick_sort(array, pivotIndex + 1, end)

def partition(array, start, end):
    pivotIndex = start
    pivotValue = array[end]

    for index in range(start, end -1):
        if array[index] < pivotValue:
            swap(array, index, pivotIndex)
            pivotIndex + 1            

    swap(array, pivotIndex, end)

    return pivotIndex

def swap(array, indexA, indexB):
    temp = array[indexA]
    array[indexA] = array[indexB]
    array[indexB] = temp

if __name__ == "__main__":
    array = [23, 6, 16, 8, 235, 0]

    print('Original array:', array, sep=' ')

    quick_sort(array, 0, len(array) - 1)

    print('Sorted array:', array, sep=' ')