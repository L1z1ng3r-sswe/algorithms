class Solution:
  def valid_tree(self, n: int, edges: list[list[int]]) -> bool:
    if len(edges) != n-1: # disconnected graph
      return False
    
    dsu = DSU(n)
    
    for u, v in edges:
      if not dsu.union(u, v):
        return False
    
    return True


class DSU:
  def __init__(self, n: int):
    self.parent = list[int] = [range(n)]
    self.size = list[int] = [1] * n

  def find(self, x: int) -> int:
    root = x

    while root != self.parent[root]:
      root = self.parent[root]

    while x != self.parent[x]:
      nxt = self.parent[x]
      self.parent[x] = root
      x = nxt

    return root
  
  
  def union(self, a:int , b: int) -> bool:
    ra, rb = self.find(a), self.find(b)

    if ra == rb: # cycle detected
      return False
    
    if self.size[ra] < self.size[rb]:
      self.parent[ra] = rb
      self.size[rb] += self.size[ra]
    else:
      self.parent[rb] = ra
      self.size[ra] += self.size[rb]

    return True