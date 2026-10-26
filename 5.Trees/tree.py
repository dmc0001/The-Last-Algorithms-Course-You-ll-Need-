class node_tree:
  def __init__(self,data,l_child=None,r_child=None):
    self.data = data
    self.l_child= l_child 
    self.r_child= r_child 


# PRE-ORDER TRAVERSEL
def walk_pre_order(tree,path):
  if tree == None:
    return
  path.append(tree.data)
  walk_pre_order(tree=tree.l_child,path=path)
  walk_pre_order(tree=tree.r_child,path=path)

def pre_order(tree):
    path = []
    walk_pre_order(tree, path)
    return path

# IN-ORDER TRAVERSEL
def walk_in_order(tree,path):
  if tree == None:
    return
  walk_in_order(tree=tree.l_child,path=path)
  path.append(tree.data)
  walk_in_order(tree=tree.r_child,path=path)

def in_order(tree):
    path = []
    walk_in_order(tree, path)
    return path

# POST-ORDER TRAVERSEL
def walk_post_order(tree,path):
  if tree == None:
    return
  walk_post_order(tree=tree.l_child,path=path)
  walk_post_order(tree=tree.r_child,path=path)
  path.append(tree.data)

def post_order(tree):
    path = []
    walk_post_order(tree, path)
    return path


if __name__ == "__main__":
    # Create tree:
    #
    #         1
    #        / \
    #       2   3
    #      / \   \
    #     4   5   6
    #

    root = node_tree(
        1,
        node_tree(
            2,
            node_tree(4),
            node_tree(5)
        ),
        node_tree(
            3,
            None,
            node_tree(6)
        )
    )

    path = []
    path2 = []
    path3 = []
    path = pre_order(root)
    path2 = in_order(root)
    path3= post_order(root)

    print(path)
    print(path2)
    print(path3)