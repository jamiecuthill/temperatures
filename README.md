# Find the nearest temperature

## Assumptions
1. Temperatures are shown as integers in the chart. The correct way to model temperatures is with floating point numbers.
2. The chart shows +8 in the negative space. Assuming this should be -8 for the purpose of this test.
3. If the input is between two values, the first nearest is returned.

## Running
Run the script with
```
go run main.go -origin=-15
```

## Tests
Run the tests with
```
go test .
```