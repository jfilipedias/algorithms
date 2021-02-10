def insertion_sort():
    array = [44, 2, 5, 0, 15, 22]

    print('Original array:', array, '...', sep=' ')

    for index in range(1, len(array)):
        key = array[index]
        current_index = index - 1

        while (current_index >= 0 and array[current_index] > key):
            array[current_index + 1] = array[current_index]
            current_index -= 1
        
        array[current_index + 1] = key
    
    print('Sorted array:', array, '...', sep=' ')

if __name__ == "__main__":
    insertion_sort()