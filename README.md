# Printer-Scanner

## Overview
Search for printer pages with default password in specific ip range

## Usage
```
go build
./Printer-Scanner -start={IP Address} -end={IP Address}
```

## Code Flow
- Get ip range from user via command line
- Verify certain ip provides printer page and using default password there.
  - Example: Samsung printer default id/password: admin/sec00000
- Prints user the result.
  - If ip address is not printer page, don't print any result.
  - If ip address is printer page but not using default password, "{IP_ADDR},X" will be result of line.
  - If ip address is printer page and using default password, "{IP_ADDR},O,{default id},{default password}" will be result of line.

## Tech stack
 - Golang 1.18

## List of currently supported printers
TODO

## TODO
 - [ ] Actual HTTP communication logic
 - [ ] Actual HTTPS communication logic
 - [ ] Verbose
 - [ ] Multithreading or use worker pool
 - [ ] Exclude some address with wildcard
 - [ ] Unit test
 - [ ] Integration test
