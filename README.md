# go-sqlite-uuidtable
example codes

# Usage
## Usage of generate_uuid_table:
```
  -days (int)
        Number of days (default 120)
  -db (string)
        A name for the database (default "beacon.db")
  -start (string)
        Start date, something like 2020-11-10. Default: today's system date. (default "2020-11-17")
```
### example:
```go run generate_uuid_table.go -days=10
db name:  beacon.db
start date:  2020-11-17
number of days:  10
+===============================================================+
|   DATE     |            UUID                  | MAJOR | MINOR |
+===============================================================+
| 2020-11-17 | 00655dbefb564127a4d60f493e9934d7 | 55395 | 31712 |
| 2020-11-18 | 9f9c5d3ededb4302ae45c99a1394dd9f | 47899 | 48015 |
| 2020-11-19 | 423b090584df4135805a5181f3f90feb | 44350 | 32878 |
| 2020-11-20 | a99ace54deba481c9a4f7cadf8e0437b | 58060 | 33715 |
| 2020-11-21 | 0486bf099b56409f84ec147fc07fb6cc | 20238 | 44001 |
| 2020-11-22 | 105d88c6031c43509328194f0cf3e812 | 35353 | 59055 |
| 2020-11-23 | 74ffdb583b13485ea09f84b46c5379c9 | 10568 |   374 |
| 2020-11-24 | 2751e2f233804074910f7f3db4c06ff1 | 64430 |    54 |
| 2020-11-25 | 33bbec767c2b4e7ab27f557281ee00cf | 30502 | 47534 |
| 2020-11-26 | d06219a151a645eabd79ec1896d0c796 | 36761 | 35579 |
+===============================================================+
```


## Usage of get_data:
```
  -date (string)
        Date, something like 2020-11-10. Default: today's system date.  (default "2020-11-17")
  -db (string)
        The name of the database (default "beacon.db")
```

### example:
```
go run get_data.go -date=2020-11-20
2020-11-20 | a99ace54deba481c9a4f7cadf8e0437b | 58060 | 33715
```
