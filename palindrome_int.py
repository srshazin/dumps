# Problem: Check if number is pallindrome
# Algorithm:
#1 Find the reverse of the number 
#2 Compare the numbers
def isPallindrome(n):
  # First identify the reverse of the number 
  num = n
  reverse = 0
  while num != 0:
    # 123%10 = 3 so we get the last digit
    # then we multiply current reversed int by 10 and add the remainder
    reverse = reverse * 10 + int(num%10)
    # 123/10 = 12 so everytime we are removing the last digit
    num = int(num/10)
  # Finally check if the original and the reversed number is same 
  if num == reversed:
    return True
  else:
    return False
  
