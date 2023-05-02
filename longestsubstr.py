def contains_dup(str):
    str_arr = list(str)
    if len(str) != len(set(str_arr)):
        return True
    else:
        return False

def LongestSubString(s):
    str = s
    max_len = 1
    cursor = 2
    i = 0
    while max_len < len(str):
        if contains_dup(str[0:cursor]) == False:
            if max_len < len(str[0:cursor]):
                max_len = cursor
            cursor += 1
        if contains_dup(str[0:cursor]):
            # Find where the duplicte is found and take the first index
            split_index = str.find(str[cursor-1]) +1
            str = str[split_index:]
            #Reseting the cursor 
            cursor = 2
    return

LongestSubString("dvdf")