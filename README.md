# go-sqlite-uuidtable
example codes

#Usage
##Usage of generate_uuid_table:
  -days int
        Number of days (default 120)
  -db string
        A name for the database (default "beacon.db")
  -start string
        Start date, something like 2020-11-10. Default: today's system date. (default "2020-11-17")

###example go run generate_uuid_table -db=beacon.db -start=2020-11-20 -days=10


##Usage of get_data:
  -date string
        Date, something like 2020-11-10. Default: today's system date.  (default "2020-11-17")
  -db string
        The name of the database (default "beacon.db")

###ecample go run get_data.go -date=2020-11-20
