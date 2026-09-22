# guess-it-2 — Product Requirements Document

Status: Draft for implementation
Language: Go
Source: Zone01 guess-it-2 assignment

## 1. Purpose

Build a command-line program that reads numbers one at a time from standard input and prints a range predicting the next number.

Use the statistical knowledge developed in linear-stats, particularly linear regression, to balance prediction accuracy against range width.

## 2. User Flow

1. Receive one number from standard input.
2. Update the observed history.
3. Calculate a prediction range for the next number.
4. Print the range immediately.
5. Wait for the next input.

The program must process input continuously. It must not wait for the input stream to end before producing predictions.

## 3. Assignment Requirements

- Read one number per line from standard input.
- Print one prediction range after each valid input, including the first.
- Print the lower and upper bounds separated by a single space.
- Use only observations already received to predict the next value.
- Treat input positions as x values starting at 0.
- Treat input numbers as y values.
- Apply calculations from linear-stats.
- Prioritize performance and the balance between successful predictions and range width.
- Support evaluation using Data 4 and Data 5.
- Include all required solution files inside the tester's student/ folder.
- Provide an executable student/script.sh that launches the solution when invoked from the tester root.

## 4. Input and Output

The following behavior decisions belong to this implementation and are not additional rules quoted from the assignment.

### Input

- Accept one signed base-10 integer per line.
- Accept values within the supported integer parser range.
- Store observations in a format suitable for floating-point calculations.
- Treat blank lines, non-integer text, and out-of-range integers as invalid input.
- On invalid input or a read failure, report the error to standard error and stop.
- Do not print a prediction for an invalid line.
- At normal end-of-input, stop without printing an extra prediction.

Example input:

```text
100
110
120
```

### Output

Print integer bounds using this format:

```text
lower upper
```

Example output line:

```text
130 150
```

This is a format example, not a required prediction for a particular input.

Output rules:

- Print exactly one range per valid input.
- End each output line with a newline.
- Ensure lower <= upper.
- Write predictions to standard output.
- Write error messages to standard error.
- Do not print prompts, labels, explanations, or debug messages to standard output.
- Round fractional lower bounds down and fractional upper bounds up so rounding does not shrink the intended interval.

## 5. Prediction Approach

### Linear Regression

For two or more observations, calculate the regression line:

```text
y = a*x + b
```

Where:

- a is the slope.
- b is the intercept.
- x is the observation position.
- y is the estimated value.

If the program has received n observations at positions 0 through n-1, predict the next value using:

```text
nextX = n
prediction = a*nextX + b
```

### Initial Range

The first working version will use a fixed margin of 10 units:

```text
lower = prediction - 10
upper = prediction + 10
```

This margin is a temporary development baseline. It is not an assignment requirement or a proven final strategy.

### First Input

The current regression formula requires at least two observations.

After the first valid input:

- Use the received value as the prediction center.
- Apply the initial margin.
- Print a range immediately.
- Do not call the regression formula.

### Later Improvements

After evaluating the baseline, possible improvements include:

- Adjusting the range margin.
- Using a recent window of observations.
- Using additional statistics to estimate prediction uncertainty.

These are experiments for later development, not required features of the first working version.

## 6. Edge Cases

| Situation | Expected behavior |
| --- | --- |
| No input | Exit without a prediction. |
| One valid number | Print a fallback range without dividing by zero. |
| Two valid numbers | Begin regression-based prediction. |
| Repeated equal values | Produce finite, ordered bounds around a stable estimate. |
| Negative values | Process them as valid integers. |
| Invalid or blank line | Report an error to standard error and stop. |
| Normal end-of-input | Finish without an extra output line. |
| Sudden change in values | Continue producing valid ranges; prediction success is not guaranteed. |

Calculations and integer conversions must not produce invalid output or panic for the supported input range.

## 7. Performance and Scoring

The tester rewards successful predictions based on range size.

- Wider ranges may capture more values but receive fewer points per successful prediction.
- Narrower ranges receive more points when successful but may miss more values.

The supplied assignment does not specify:

- An exact scoring formula.
- A numeric passing score.
- A latency limit.
- A memory limit.

The goal is to find a measured balance between prediction accuracy and range width on Data 4 and Data 5.

Record results for the baseline and each tested change. Do not claim improved performance until the tester results support it.

Do not hardcode dataset values or use future observations.

## 8. Tester Integration

The final submission must include:

- A student/ folder inside the provided tester root.
- All files required to build and run the Go solution.
- An executable script.sh inside student/.
- Launch commands that work when the current directory is the tester root.

The solution must run successfully with both Data 4 and Data 5.

## 9. Acceptance Criteria

- [ ] The Go project builds and runs from the terminal.
- [ ] Each valid input immediately produces exactly one prediction range.
- [ ] Every range contains two integer bounds separated by a single space.
- [ ] The first input produces a fallback range.
- [ ] Regression is used only when at least two observations are available.
- [ ] After receiving 100, 110, and 120, the regression estimate for the next value is 130.
- [ ] Constant and negative sequences produce valid bounds.
- [ ] Fractional bounds are rounded outward correctly.
- [ ] Invalid input produces an error only on standard error.
- [ ] Normal end-of-input produces no extra prediction.
- [ ] No prompts or debug messages appear on standard output.
- [ ] student/script.sh is executable and starts the solution from the tester root.
- [ ] Data 4 and Data 5 run successfully.
- [ ] Baseline tester scores are recorded for comparison.

Checks may initially be performed manually. Add automated checks where they resolve a concrete correctness risk.

## 10. Code Organization

- Keep main() focused on starting the program and handling errors.
- Keep input handling in a separate file.
- Keep regression calculations in a separate file.
- Keep prediction logic separate when introducing its own function.
- Avoid unnecessary dependencies and unrelated refactoring.

The exact file structure may evolve as the implementation grows.

## 11. Development Plan

1. Complete the streaming baseline.
2. Add first-input handling.
3. Produce integer prediction bounds.
4. Verify input handling and output formatting locally.
5. Prepare student/ and the executable launch script.
6. Run Data 4 and Data 5.
7. Record baseline results.
8. Improve the prediction strategy based on measured results.

Follow this workflow:

```text
Small change → Run → Check → Commit → Next step
```

## 12. Out of Scope

- A graphical interface.
- A custom web server.
- A database.
- User accounts.
- Machine-learning frameworks.
- Guaranteed prediction of arbitrary sequences.
- Replacing the provided tester.

## 13. Open Decisions

The following decisions remain open until the baseline has been tested:

- The final range-margin strategy.
- Whether to use the complete history or a recent window.
- Whether additional statistics improve the score enough to justify their complexity.