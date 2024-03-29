import { describe, it, expect } from 'vitest'
import { binarySearch } from '.'

describe('Test binary search algorithm', () => {
  it('Should be able to search on an empty list', () => {
    expect(binarySearch([], 1)).toEqual(-1)
  })

  it('Should be able to search on a list with one item', () => {
    expect(binarySearch([1], 1)).toEqual(0)
  })

  it('Should be able to search an item in sorted array', () => {
    expect(binarySearch([1, 2], 1)).toEqual(0)
    expect(binarySearch([1, 2], 2)).toEqual(1)
    expect(binarySearch([1, 5, 10, 12], 1)).toEqual(0)
    expect(binarySearch([1, 5, 10, 12, 14, 17, 22, 100], 17)).toEqual(5)
    expect(binarySearch([1, 5, 10, 12, 14, 17, 22, 100], 1)).toEqual(0)
    expect(binarySearch([1, 5, 10, 12, 14, 17, 22, 100], 100)).toEqual(7)
  })

  it('Should be able to return -1 if the item does not exists', () => {
    expect(binarySearch([1, 5, 10, 12, 14, 17, 22, 100], 0)).toEqual(-1)
  })
})
