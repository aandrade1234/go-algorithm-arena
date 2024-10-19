# 269. Alien Dictionary

There is a new alien language that uses the English alphabet. However, the order of the letters is unknown to you.

You are given a list of strings `words` from the alien language's dictionary. Now it is claimed that the strings in `words` are sorted lexicographically
by the rules of this new language.

If this claim is incorrect, and the given arrangement of string in words cannot correspond to any order of letters, return `""`.

Otherwise, return a string of the unique letters in the new alien language sorted in lexicographically increasing order by the new language's rules. If there are multiple solutions, return any of them.

**Example 1:**
> **Input:** words = ["wrt","wrf","er","ett","rftt"]<br>
> **Output:** "wertf"<br>

**Example 2:**
> **Input:** words = ["z","x"]<br>
> **Output:** "zx"

**Example 3:**
> **Input:** words = ["z","x","z"]<br>
> **Output:** ""<br>
> **Explanation:** The order is invalid, so return "".

**Constraints:**
- `1 <= words.length <= 100`
- `1 <= words[i].length <= 100`
- `words[i]` consists of only lowercase English letters.

Source: [Leetcode](https://leetcode.com/problems/alien-dictionary/description/)
