# LeetCode Solutions

A collection of LeetCode algorithm and SQL problem solutions written in Go and SQL.

## Tech Stack

- **Go** (`1.22`, module `github.com/daopmdean/leetcode`)
- **SQL**

## Project Structure

```text
.
├── database/
│   └── easy/           # Easy SQL / database problems (6 files)
├── easy/               # Easy algorithm problems in Go (13 files)
├── interview_150/      # Top Interview 150 solutions in Go (34 files)
├── medium/             # Medium algorithm problems in Go (7 solutions, 1 test)
├── sql50/              # SQL 50 study plan solutions (1 file)
├── go.mod              # Go module definition
└── main.go             # Entry point
```

## Categories

| Directory | Language | Description | Count |
| --- | --- | --- | --- |
| [`easy/`](./easy) | Go | Easy algorithm problems | 13 |
| [`medium/`](./medium) | Go | Medium algorithm problems | 7 |
| [`interview_150/`](./interview_150) | Go | LeetCode Top Interview 150 problems | 34 |
| [`database/easy/`](./database/easy) | SQL | Easy database query problems | 6 |
| [`sql50/`](./sql50) | SQL | LeetCode SQL 50 study plan | 1 |

## Running Tests

Run unit tests across the Go packages with:

```sh
go test ./...
```
