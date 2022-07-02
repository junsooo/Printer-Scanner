# Printer-Scanner

## Overview
Search for printer pages with default password in specific ip range

## Flow
- Get ip range from user via command line
- Verify certain ip provides printer page and using default password there.
  - Example: Samsung printer default id/password: admin/sec00000
- Prints user the result.

## Usage
```
go build
./Printer-Scanner -start={IP Address} -end={IP Address}
```

## Tech stack
 - Golang 1.18

## List of currently supported printers
TODO

## TODO
 - [ ] Actual HTTP communication logic
 - [ ] Exclude some address with wildcard
 - [ ] Unit test
 - [ ] Integration test
