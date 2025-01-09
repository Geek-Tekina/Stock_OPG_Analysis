Overview :

Building A Stock Analysis CLI Application
-> It uses OPG startegy, Open Price Gap (need not go in depth of formulas used)

The project can be broken into phases : 
-> Load Stocks from a CSV File 
-> Filter out unworthy Stocks
-> Do some processing on the stocks - position (quantity and price)
-> fetch news on each stocks
-> Produce output as JSON
---------------------------------------------------------------------------------------------------------------

Day wise Progress -

Day: 1
Our CSV file is in this format 
row 1 = Ticker | row 2 = Gap | row 3 = Opening Price  

    Reading the CSV File 
    -> Before reading the CSV file, program needs to open the file. We use os.Open("/path) method.
    -> csv.NewReader is used to read the csv file

    Stock Struct
    -> Now we need a type to store the data, thus we create a Struct named Stock having Ticker, Gap and OpeningPrice
    -> then we will have a slice of Stock (Struct) Objects

    Here we encounter error, as the values from CSV files are string which can be intialised to Stock Struct having float64 type. So we use strconv.ParseFloat() method to convert string to float 32/64

    Closing the File
    -> We use defer f.close()
---------------------------------------------------------------------------------------------------------------
