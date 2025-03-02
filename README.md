# Activity Rank - Algorithm & Execution

## Overview
This Go project calculates an activity score for repositories based on commit data from a CSV file. It ranks repositories and displays the top 10 most active ones.

## How It Works
The algorithm follows these steps:

1. **Read CSV Data**: Extract commit details from a CSV file.
2. **Process Data**: For each commit, retrieve:
   - Number of files changed
   - Lines added
   - Lines deleted
3. **Calculate Activity Score**:
   
   ```
   Score = (files changed * 2) + additions + deletions
   ```
4. **Sort Repositories**: Rank in descending order based on score.
5. **Display Top 10 Repositories**.

## How to Run

### Requirements
- Go installed
- `commits.csv` file in the project directory

### Steps

1. Clone or download this repository.
2. Navigate to the project folder.
3. Run the program:
   
   ```sh
   go run main.go
   ```

This will print the top 10 most active repositories in the terminal.

## Example Output

```sh
Top 10 most active repositories:
1. repo-A - Score: 3456
2. repo-B - Score: 3120
3. repo-C - Score: 2987
...
```

## Tests

Run unit tests
   ```sh
   go test
   ```