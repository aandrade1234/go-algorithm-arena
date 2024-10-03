# Office Desk Booking System

In a modern office environment, employees often work in dynamic teams and need the flexibility to book desks for their groups. Your company has developed an online desk booking system to manage desk allocations efficiently. The system needs to ensure that groups of employees are happy by allocating them desk blocks that can accommodate their entire group without splitting them up.

**Task:**

You are tasked with writing a function to optimize the desk booking system. The function should take into account the number of desks each group of employees requests and the sizes of available desk blocks in the office. The goal is to maximize the number of happy employee groups, where a group is happy if all its members can be seated together in a single desk block.

**Function Signature:**

```go
func scheduleDesks(userRequestSize []int, deskBlockSize []int) int
```

**Input:**

- `userRequestSize:` An array of integers where each element represents the number of desks requested by a group of employees.
- `deskBlockSize:` An array of integers where each element represents the size of a block of desks available in the office.

**Output:**

- An integer representing the number of happy employee groups.

**Example:**

```go
   userRequests := []int{2, 1, 4, 3}
   deskBlocks := []int{2, 3, 4, 5}

   // Output: Number of happy user groups: 4
   happyUsers := scheduleDesks(userRequests, deskBlocks)
```

Source: Google Interview
