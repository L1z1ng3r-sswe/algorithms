class DSU:
  def __init__(self, n: int):
    self.parent = [i for i in range(n)]
    self.size = [1] * n # optimized union with size
  
  # bfs find with path compression
  def find(self, x: int) -> int:
    root = x
    while self.parent[root] != root:
      root = self.parent[root]
    
    # path compression
    while self.parent[x] != x:
      nxt = self.parent[x]
      self.parent[x] = root
      x = nxt

    return root

  def union(self, a: int, b: int):
    ar, br = self.find(a), self.find(b)
    if ar == br:
      return
    
    if self.size[ar] < self.size[br]:
      self.parent[ar] = br
      self.size[br] += self.size[ar]
    else:
      self.parent[br] = ar
      self.size[ar] += self.size[br]

class Solution:
  def numberOfConnected(n: int, edges: list[list[int]]) -> int:
    dsu = DSU(n)
    for a, b in edges:
      dsu.union(a, b)

    res = 0
    for i, parent in enumerate(dsu.parent):
      if i == parent:
        res += 1

    return res
  
# n - number of nodes
# m - number of edges

# time: O(m * a(n))
# space: O(n)