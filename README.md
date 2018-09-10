# toyrobot

```
$ go build

$ ./toyrobot PLACE 1,2,EAST MOVE LEFT LEFT MOVE RIGHT MOVE REPORT
$ -> position: 1,3,NORTH

$ ./toyrobot PLACE 3,2,SOUTH REPORT RIGHT MOVE REPORT LEFT MOVE REPORT RIGHT MOVE REPORT
$ -> position: 3,2,SOUTH
$ -> position: 2,2,WEST
$ -> position: 2,1,SOUTH
$ -> position: 1,1,WEST
```

