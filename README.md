# guess-it-2

A command-line program written in Go that reads integers from standard input and predicts a range for the next value using linear regression.

Developed as part of the Zone01 curriculum, this project builds on the statistical calculations introduced in `math-skills` and `linear-stats`.

## How It Works

After each valid input, the program:

1. Adds the number to its history.
2. Keeps up to the 20 most recent observations.
3. Fits a regression line when at least two observations are available.
4. Estimates the next value and prints an integer range around it.

The regression equation is:

```text
y = a*x + b
```

Where:

- `a` is the slope.
- `b` is the intercept.
- `x` is the position within the current observation window, starting at zero.
- `y` is the estimated value.

For a window containing `n` observations, the next value is estimated at `x = n`.

With only one observation, the program uses that value as the prediction center.

The range uses a fixed margin around the estimate. The lower bound is rounded down and the upper bound is rounded up before conversion to integers. The margin is defined in `import.go`.

## Requirements

- Go compatible with the version declared in `go.mod`.
- A terminal.

The program uses only the Go standard library.

## Run

From the project directory:

```bash
go run .
```

Enter one integer per line and press Enter. The program prints a prediction range after each number and then waits for the next input.

To finish input on macOS or Linux, press `Ctrl+D` on an empty line.

## Build

Create an executable:

```bash
go build -o guess-it-2 .
```

Run it on macOS or Linux:

```bash
./guess-it-2
```

Input can also be redirected from a file containing one integer per line:

```bash
./guess-it-2 < numbers.txt
```

## Input and Output

Input consists of signed base-10 integers, one per line:

```text
100
110
120
```

Each output line contains two integer bounds separated by a space:

```text
lower upper
```

Each range predicts the value that comes **after** the number just received.

Predictions are written to standard output without labels or additional messages.

## Error Handling

The program reports an error to standard error and stops when it encounters:

- Non-integer input.
- An empty line.
- An integer outside the supported parsing range.
- An input-reading error.

Normal end-of-input stops the program without an extra prediction.

## Project Files

| File | Purpose |
| --- | --- |
| `main.go` | Starts the program and reports errors. |
| `import.go` | Reads input, maintains the observation window, and prints prediction ranges. |
| `regression.go` | Calculates the regression slope and intercept. |
| `go.mod` | Defines the Go module and language version. |
| `PRD.md` | Describes the project requirements and scope. |

## Limitations

Linear regression estimates a trend from recent observations; it cannot guarantee the next value.

Abrupt changes and irregular sequences can reduce prediction accuracy. The fixed margin does not automatically adapt to changes in variability.