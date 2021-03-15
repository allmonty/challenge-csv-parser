# csv-parser-challenge

Rain's job challenge

Read the [description](https://github.com/allmonty/csv-parser-challenge/blob/main/Rain%20CSV%20Parser%20v2.txt)

The idea of my solution is to have maps that define multiple possibilities for row names that represent the same thing.
The program will use those maps to identify the desired columns in every CSV and create another CSV with just those columns.
Any line with missing required fields will be stored in `flagged.csv` with the reason and original content.

---

### Testing

To run the test just run in the root folder:
```sh
make test
```

### Using the parser

To run the parser execute in the root folder:
```sh
go build
go run main.go <csv_file_paths ...>
```

You can see an example of running the parser in the [Makefile](https://github.com/allmonty/csv-parser-challenge/blob/main/Makefile)

After running, the output of the parser will be saved in the `result` folder, there will be two files:
- `parsed.csv` with all the successfuly processed lines
- `flagged.csv` with all the lines that had any missing required column
