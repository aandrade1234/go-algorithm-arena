# 939. Minimum Area Rectangle

You are given an array of points in the X-Y plane `points` where `points[i] = [xi, yi]`.

Return the minimum area of a rectangle formed from these points, with sides parallel to the X and Y axes. If there is not any such rectangle, return `0`.

**Example 1:**

<img src="./image1.jpeg">

> **Input:** points = [[1,1],[1,3],[3,1],[3,3],[2,2]] <br>
> **Output:** 4

**Example 2:**

<img src="./image2.jpeg">

> **Input:** points = [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]] <br>
> **Output:** 2

**Constraints:**
- `1 <= points.length <= 500`
- `points[i].length == 2`
- `0 <= xi, yi <= 4 * 10^4`
- All the given points are unique.
