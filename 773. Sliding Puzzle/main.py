import queue

class Solution(object):
    def slidingPuzzle(self, board):
        """
        :type board: List[List[int]]
        :rtype: int
        """
        m = {0: (1,3), 1: (0,2,4),2:(1,5),3:(0,4),4:(1,3,5),5:(2,4)}

        flat = [item for sub in board for item in sub]
        visited = []

        q = queue.Queue()
        moves = 0

        q.put(flat)

        while not q.empty():
            for i in range(q.qsize()):
                move = q.get()
                if tuple(move) == (1,2,3,4,5,0):
                    return moves
                idx = move.index(0)

                for next in m[idx]:
                    t = move.copy()
                    t[idx], t[next] = t[next], t[idx]
                    if not t in visited:
                        q.put(t)
                        visited.append(t)
            
            moves+=1
                    
        return -1

s = Solution()
print(s.slidingPuzzle([[1,2,3],[4,0,5]]))