def remove_dup(str="pwwkew"):
    tmp_str = ''
    for i in range(len(str)):
        if str[i] not in tmp_str:
            tmp_str += str[i]
    return tmp_str

def contains_dup(str):
    str_arr = list(str)
    if len(str) != len(set(str_arr)):
        return True
    else:
        return False

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_len = 0
        cursor  = 0
        range_ = 1
        x = 0
        while x < 100:
            # for j in range(range_):
                # print(range_)
                # if contains_dup(s[cursor:range_]) == False and len(s[cursor:range_]) > max_len:
                #     max_len = len(s[cursor:range_])
                #     cursor +=1
                #     print("if")
                # if contains_dup(s[cursor:range_]):
                #     print(f"Duplicate found at index {j}")
            if contains_dup(s[cursor:range_]) == False and len(s[cursor:range_]) > max_len: 
                range_ += 1
                max_len = range_ - 1
                print("Increased")

                print(max_len)

            if contains_dup(s[cursor:range_]):
                cursor = s.find(s[range_-1]) + 1
                cursor = max_len
                
            x += 1

print(Solution.lengthOfLongestSubstring(Solution, "dvdf"))