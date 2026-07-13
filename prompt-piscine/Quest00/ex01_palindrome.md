# Write pseudoscde for a function that checks if a string is a palindrome.

1. FUNCTION is_palindrome(string), which takes string as a parameter
2. Normalizes the string 
3. Initailises two pointers, left = 0 and right = lenght - 1
4. compare the string from both ends
5. Then, return False if the character at left is not equal to the right.
6. And return True, if there is no mismatch.

# Implement your solution in Python.

```python
def is_palindrome(s):
    """
    Check if a string is a palindrome.
    """
    
    # Initialize pointers
    left = 0
    right = len(s) - 1
    
    # Compare characters from both ends
    while left < right:
        if s[left] != s[right]:            
            return False
        left += 1
        right -= 1
    
    return True

# Test cases

examples = ["racecar", "hello", "A man a plan a canal Panama"]

for word in examples:
    result = is_palindrome(word)
    print(f"'{word}' = {result}")
```

# What did you learn from solving it before asking AI?
I learned about;

- How loops work

- The meaning of palindrome

# How is your understanding different now?
My understanding is now much better as I have gained deeper understanding and seen different approaches to implement it. Specifically, how to structure my program to properly handle empty spaces, ascented letter and quotation marks enclosed by string "!!!".

# Could you now write similar function (e.g reverse a string) without help?
Yes, I can do it now.
