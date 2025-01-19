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

---

### Day 3

#### Calculating positions for the shares

- Postion is : Entry price, number of shares, target profit, stop less price
- For calculation we have coded some hard coded values and some formulas
- To store the result, a struct Position is created

```go type
Position struct {
EntryPrice float64 // the price at which to buy or sell
Shares int // amount of shares to buy / sell
TakeProfitPrice float64 // the price at which to take exit and make profit
StopLossPrice float64 // the price at which to stop my loss if stock doesn't go our way
Profit float64 // expected final profit
}
```

- then `calculate()` function
- Now we have instance of Stock struct and Position Struct for each stocks we have. We need to combine both under
  Selection Struct having Ticker and Postion : this will help user to determine if to buy the stock or not

```go
type Selection struct {
Ticker string
Position
}
```

---

### Day 4

#### Fetching news

- We will fetch news on each stock from `seekingalpha.com`
- first we created a http client `http.NewRequest(http.MethodGet, url + ticker, nil)`, this takes 3 argument : method, url, req-body if any
- We added a header, and then created http Client to make request using `Do()` function.
- `Do()` fnc return a pointer to http response object

##### Mapping response data and creating structs to read data from response

Demo response

```go
{
    "data": [
        {
            "id": "4395525",
            "type": "news",
            "attributes": {
                "publishOn": "2025-01-16T05:56:23-05:00",
                "isLockedPro": false,
                "commentCount": 10,
                "gettyImageUrl": "https://static.seekingalpha.com/cdn/s3/uploads/getty_images/2193527787/image_2193527787.jpg",
                "videoPreviewUrl": null,
                "videoDuration": null,
                "themes": {},
                "title": "Biden warns of ultrarich oligarchy, tech-industrial complex in farewell speech",
                "isPaywalled": false
```

- So here we will focus on PublishOn and Title attributes from the `res.Body.data`, So we have created a struct
  ```go
  type attributes struct {
  PublishOn time.Time `json:publishOn`
  Title     string    `json:title`
  }
  ```
  the `` helps to clearify the mapping process that PublishOn will be mapped to json - publishOn attribute, it is a good practise
- Now as `attributes` is inside `data`, we need to have data types that can contain the required format.

  ```go
  type attributes struct {
  PublishOn time.Time `json:publishOn`
  Title     string    `json:title`
  }

  type seekingAlphaNews struct {
    Attributes attributes `json:attributes`
  }
  type SeekingAlphaResponse struct {
    Data []seekingAlphaNews `json:data`      // response can be mapped to SeekinAlphaNews
  }
  ```

- Now we map all the req data in our struct
  ```go
  res := &SeekingAlphaResponse{}
  json.NewDecoder(response.Body).Decode(res)
  ```
- Now finally we have articles slice of Article Strcut objects, which are returned along with the error.

### Day 5

- Now we will update Selection struct to have slice of articles for each stock
  ```go
  type Selection struct {
  Ticker string
  Position
  Articles []Article
  }
  ```
- We will call `FetchNews()` func in the for_range loop when creating selection stock and set the articles as well

### Day 6

- Now we will write the output into the file as JSON using `Deliver()` fnc.
  - First using os package file is created
  - then using encoder for JSON, the values from selections slice is written in file !!

#### Happy Stoking with Go !!
