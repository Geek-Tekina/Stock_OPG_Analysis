# Stock Analysis CLI Application

## Overview

This CLI application is built to analyze stocks using the **OPG (Open Price Gap)** strategy.

The project can be broken down into the following phases:

- **Load Stocks from a CSV File**
- **Filter out unworthy Stocks**
- **Process Stocks (position: quantity and price)**
- **Fetch News for each Stock**
- **Produce Output in JSON Format**

---

## Day-Wise Progress

### Day 1

#### CSV File Format

The CSV file is structured as follows:

- **Row 1**: Ticker
- **Row 2**: Gap
- **Row 3**: Opening Price

#### Reading the CSV File

- Before reading the CSV file, the program needs to open the file. This is done using the `os.Open("/path")` method.
- We use `csv.NewReader` to read the contents of the CSV file.

#### Stock Struct

- To store the stock data, we create a `Stock` struct with the following fields:
  - **Ticker**
  - **Gap**
  - **Opening Price**
- We then use a slice to hold multiple `Stock` struct objects.

#### Data Type Conversion

- CSV file values are read as strings, but the `Stock` struct fields (`Gap` and `OpeningPrice`) are of type `float64`.
- To handle this, we use `strconv.ParseFloat()` to convert the string values to `float64`.

#### Closing the File

- We use `defer f.Close()` to ensure the file is closed properly after the operation is complete.

---

### Day 2

#### Filtering Unworthy Stocks

- Used `DeleteFunc()` of slices to delete the unworthy stocks.
