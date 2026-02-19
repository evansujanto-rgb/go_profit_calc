# Profit Calculator

A simple command-line profit calculator written in Go that helps calculate and track business metrics including Earnings Before Tax (EBT), Profit, and Profit Ratio.

## Features

- Calculate key business metrics:
  - Earnings Before Tax (EBT)
  - Profit (after tax)
  - Profit Ratio (EBT/Profit)
- Interactive command-line interface
- Automatic data persistence to file
- Input validation to prevent negative values

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd profit_calc
```

2. Build the project:
```bash
go build
```

## Usage

Run the compiled binary:
```bash
./profit_calc
```

The program will prompt you to enter three values:
1. Revenue: The total income of the business
2. Expense: The total expenses of the business
3. Tax Rate: The tax rate as a percentage (e.g., enter 10 for 10%)

After entering the values, the program will:
1. Calculate and display:
   - EBT (Earnings Before Tax)
   - Profit (after tax)
   - Profit Ratio (EBT/Profit)
2. Save the results to `profit_data.txt`

### Example Output

```
Please enter Revenue: 5000
Please enter Expense: 3000
Please enter Tax Rate: 10
EBT: 2000.0
Profit: 1800.0
Ratio: 1.11
```

## Data Storage

The program automatically saves the calculation results to `profit_data.txt` in the following format:
```
EBT: <value>
Profit: <value>
Ratio: <value>
```

## Error Handling

The program includes basic error handling:
- Prevents negative input values
- Validates all numeric inputs
- Provides clear error messages for invalid inputs

## Project Structure

- `profit_calc.go`: Main program file containing all the logic
- `profit_data.txt`: Data file storing the calculation results
- `profit_calc`: Compiled binary (generated after building)

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the MIT License. 