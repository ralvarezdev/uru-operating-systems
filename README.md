# uru-operating-systems

**Note:** Archived and read-only. Kept for reference from the Operating Systems college course.

My projects from the Operating Systems college course. The `tasks/` module is a CLI that simulates CPU scheduling algorithms over tasks loaded from CSV data.

Go module: `github.com/ralvarezdev/uru-operating-systems/tasks` (Go 1.23.4).

## What it does

`tasks/cmd/main.go` presents a menu to run one of three scheduling algorithms: FIFO, LIFO, or Round Robin. Each `internal/algorithms/*.go` file implements one algorithm against `Task` structs that track id, arrival time, duration, and computed end time / processing time / waiting time / service index. `internal/algorithms/average.go` and `timing.go` provide aggregate metrics helpers.

Task data is read from CSV files in `tasks/data/` (`all.csv`, `class.csv`, `classroom.csv`, and the `tmax-*` sample sets).

## Project structure

```
tasks/
  cmd/main.go                    # interactive CLI menu
  internal/
    algorithms/                  # types.go, fifo.go, lifo.go, rr.go, average.go, timing.go
    data/                        # read.go, write.go, errors.go — CSV I/O
  data/*.csv                     # sample task datasets
  compile.bat, run.bat
```

## Running

```bash
cd tasks
go run ./cmd           # or: compile.bat / go build ./cmd && ./cmd
```

Pick an option from the menu to run FIFO, LIFO or Round Robin against the CSV data in `tasks/data/`.

## License

GNU General Public License v3.0 (see `LICENSE`).
