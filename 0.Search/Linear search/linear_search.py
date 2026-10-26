def linear_search (n,element):
    for i in n:
        if (i == element):
          return True;
    return False



result = linear_search([1, 3, 67, 232, 0, 23, 1], 3);
print(result)
