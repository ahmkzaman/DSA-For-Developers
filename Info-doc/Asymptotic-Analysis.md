# Asymptotic Analysis

Asymptotic analysis is a method used to describe the limiting behavior of a function or algorithm as its input size increases towards infinity. It's particularly useful in computer science for understanding the efficiency and scalability of algorithms by focusing on how their performance changes with large inputs, ignoring constant factors and low-level details.

---

## 🔑 Key Concepts

- **Focus on Growth Rate**: Focuses on how the algorithm's time/space grows with input size.
- **Asymptotic Notations**: Uses mathematical tools like Big O, Big Omega, and Big Theta to describe performance.
- **Ignoring Constants**: Ignores constant multipliers and small inputs since they become negligible.
- **Scalability**: Helps evaluate how algorithms behave with very large datasets.
- **Performance Comparison**: Provides a framework to compare algorithms' efficiency at scale.

---

## 🧮 Notations

| Notation | Bound Type  | Description                                                | Example                        |
| -------- | ----------- | ---------------------------------------------------------- | ------------------------------ |
| **O(n)** | Upper Bound | Worst-case time. Algorithm takes _at most_ this much time. | Found book at last index       |
| **Ω(n)** | Lower Bound | Best-case time. Algorithm takes _at least_ this much time. | Found book at first index      |
| **Θ(n)** | Tight Bound | Exact bound. Algorithm always takes this much time.        | Same for best, worst, and avg. |

---

## 📈 Common Growth Functions

| Notation       | Name              | Example                           | Description                                  |
| -------------- | ----------------- | --------------------------------- | -------------------------------------------- |
| **O(1)**       | Constant Time     | `arr[5]`                          | Always takes the same time                   |
| **O(log n)**   | Logarithmic Time  | Binary Search                     | Very efficient; grows slowly                 |
| **O(n)**       | Linear Time       | Array Traversal                   | Time grows linearly with input size          |
| **O(n log n)** | Linearithmic Time | Merge Sort, Quick Sort (avg case) | Between linear and quadratic                 |
| **O(n²)**      | Quadratic Time    | Bubble Sort                       | Grows quickly as input size increases        |
| **O(n³)**      | Cubic Time        | Matrix Multiplication             | Even slower                                  |
| **O(nᵏ)**      | Polynomial Time   | Varies                            | Very slow for large k                        |
| **O(2ⁿ)**      | Exponential Time  | Recursive Fibonacci               | Becomes huge very quickly                    |
| **O(n!)**      | Factorial Time    | Brute-force Traveling Salesman    | Extremely slow; impractical for large inputs |

---

## 📊 Best, Worst, and Average Case

| Case Type      | Description                                 | Example (Linear Search)                |
| -------------- | ------------------------------------------- | -------------------------------------- |
| **Best Case**  | Fastest scenario                            | Element is at the first position       |
| **Worst Case** | Slowest scenario                            | Element at the end or not in the array |
| **Avg Case**   | Expected time across all possible scenarios | Element is somewhere in the middle     |

---

## ⚖️ The Domination Rule

When combining multiple terms:

- **Keep the term with the highest growth rate**.
- This becomes the **dominant term** in Big-O analysis.

For example:

```text
O(n³ + n² + log n) = O(n³)
```
