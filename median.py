def findMedian(arr1, arr2):
  # Concat array 
  res_array = arr1 + arr2
  #res_array = res_array.sort()
  res_array.sort()
  print(res_array)
  median = 0
  if len(res_array)%2 != 0:
    median = res_array[((len(res_array)+1)/2]
  else:
    median = (res_array[((len(res_array)/2) res_array[((len(res_array)/2 + 1) )/2
    retrun median
    
  
arr1 = [3, 2, 1, 4, 6 , 7]
arr2 = [1, 2]

print(findMedian(arr1, arr2))

