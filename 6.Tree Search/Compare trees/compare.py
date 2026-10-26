def compare(a,b):
  if a==None and b==None:
    return True
  if a!=None or b!=None:
    return False
  if a.v != b.v:
    return False
  return compare(a.left,b.left) and compare(a.right,b.right)