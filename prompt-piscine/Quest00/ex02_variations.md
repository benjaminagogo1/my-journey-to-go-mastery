# Modify your palindrome function to:
1 Ignore spaces and punctuation.

2 Be case-insensitive.

3 Return the position where the string stops being a palindrome (if not one).

# Ignore spaces and punctuation.
```python
def is_palindrome(s):

    s = ''.join(
        char for char in s
        if char.isalnum()
    )

    left = 0
    right = len(s) - 1

    while left < right:
        if s[left] != s[right]:
            return False
        left += 1
        right -= 1
    
    return True

examples = ["racecar", "hello", "A man a plan a canal Panama"]

for word in examples:
    result = is_palindrome(word)
    print(f"'{word}' -> {result}")
```
# case-insensitive
```python 
def is_palindrome(s):
    
    s = s.lower()
    s = ''.join(
        char for char in s 
        if char.isalnum()
    )
    
    left = 0
    right = len(s) - 1
    
    while left < right:
        if s[left] != s[right]:            
            return False
        left += 1
        right -= 1
    
    return True

examples = ["racecar", "hello", "A man a plan a canal Panama"]

for word in examples:
    result = is_palindrome(word)
    print(f"'{word}' = {result}")

```

# Return the position where the string stops being a palindrome (if not one).
```python
def is_palindrome(s):

    s = ''.join(
        char for char in s
        if char.isalnum()
    )

    left = 0
    right = len(s) - 1

    while left < right:
        if s[left] != s[right]:
           return left, right
        left += 1
        right -= 1
    
    return True

examples = ["racecar", "hello",  "A man a plan a canal Panama"]

for word in examples:
    result = is_palindrome(word)
    print(f"'{word}     {result}'")
```

# Reflect on what AI added that you didn't consider initially.


## Reverse String Comparison
```py
def is_palindrome(s):
    s = ''.join(c.lower() for c in s if c.isalnum())
    return s == s[::-1]
```
The above technique called reverse string comparison technique was provided by AI as another approach to implement my solution. The technique is not only shorter, but also simple and easy to understand as it contains fewer lines of code. Other things like edge cases, time complexity and the reason my program is on linear time complexity was provided. I also learned how to structure my program to handle input variants such as "!!!  ???", "Aa", "reifier".

# Edge cases
- "!!! ???"
- "Aa"
- ""
- "réifier"



