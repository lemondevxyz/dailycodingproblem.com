# balanced_parentheses
Given a string that contains parentheses, this program will detect any invalid ones.

Example:
```golang
solve_without_order("()") // true
solve_without_order("(()") // false
solve_without_order(")())") // true
solve_without_order("(())") // true

solve_with_order("()") // true
solve_with_order("(()") // false
solve_with_order(")())") // false
solve_with_order("(())") // true
```
