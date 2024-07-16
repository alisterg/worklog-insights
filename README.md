A program to analyse a directory of worklogs;

1. With the title `YYYY-MM-DD.md`
2. In the format:

```
| Date    | YYYY-MM-DD     |
| ------- | -------------- |
| Start   | 05:00          |
| Finish  | 15:00          |
| WFH     | y              |
| Standup | Standup notes  |
| Note    |                |

### Worklog

- [ ] Items here
```

# Usage

`work --from="YYYY-MM-DD" --to="YYYY-MM-DD"`

Will output:

1. The number of days you worked from home
2. The average number of hours worked per day

NOTE: `from` is inclusive and `to` is exclusive

# Development

Prerequisites:

- golang 1.20

Build and usage:

1. `go build -o work main.go`
2. Copy the binary to the worklog location
3. Run the commands using `./work {command}`
