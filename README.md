A program to analyse a directory of worklogs;

1. With the title `YYYY-MM-DD.md`
2. In the format:

```
| Date    | {{date}}       |
| ------- | -------------- |
| Start   | {{time}}       |
| Finish  |                |
| WFH     | y              |
| Standup | Standup notes  |
| Note    |                |

### Worklog

- [ ] Items here
```

Two commands are available:

`work wfh --from="YYYY-MM-DD" --to="YYYY-MM-DD"`

Will output the number of days you worked from home between two dates.

`work hours --from="YYYY-MM-DD" --to="YYYY-MM-DD"`

Will output the average number of hours worked between two dates.

# Development

Prerequisites:

- golang 1.20

Build and usage:

1. `go build -o work main.go`
2. Copy the binary to the worklog location
3. Run the commands using `./work {command}`
