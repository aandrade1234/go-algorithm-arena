# 20. Valid Parentheses

Given a string s containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
1. Open brackets must be closed in the correct order.
1. Every close bracket has a corresponding open bracket of the same type.

**Example 1:**

> **Input:** s = "()"<br>
  Output: true

**Example 2:**

> **Input:** s = "()[]{}"<br>
  **Output:** true

**Example 3:**

> **Input:** s = "(]"<br>
  **Output:** false

**Constraints:**

- `1 <= s.length <= 10^4`
- `s` consists of parentheses only `'()[]{}'`.

Source: [Leetcode](https://leetcode.com/problems/valid-parentheses)
