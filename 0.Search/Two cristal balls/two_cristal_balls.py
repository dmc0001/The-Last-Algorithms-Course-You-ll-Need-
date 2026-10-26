import math

def two_cristal_balls(n):
    jump_amount =  math.floor(math.sqrt(len(n)))
    i = jump_amount
    while i < len(n):
        if(n[i] == True):
            break
        i += jump_amount

    j = i - jump_amount
    while  j <= len(n):
         if(n[j] == True):
            return j
         j += 1
        
    return -1



result = two_cristal_balls([False,False,False,False,True,True,True])
print(result)
   


