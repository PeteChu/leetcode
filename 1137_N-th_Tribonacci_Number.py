# https://leetcode.com/problems/n-th-tribonacci-number/

class Solution:
    
    def __init__(self):
        self.mem = {}
    
    def tribonacci(self, n: int) -> int:
        if n == 0:
            return 0
        if n < 3:
            return 1
        if n not in self.mem:
            self.mem[n] = self.tribonacci(n -1 ) + self.tribonacci(n - 2) + self.tribonacci(n - 3)
        return self.mem[n]



print(Solution().tribonacci(999))