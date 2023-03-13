# https://leetcode.com/problems/valid-parentheses/

def isValid(s: str) -> bool:
    if len(s) % 2 != 0:
        return False
    ob = ['(', '{', '[']
    cb = [')', '}', ']']
    stack = []
    
    for i in s:
        if i in ob:
            stack.append(i)
        if i in cb:
            if len(stack) > 0 and stack[-1] == ob[cb.index(i)]:
                stack = stack[:-1]
            else:
                return False
    return len(stack) == 0