# Microsoft Excel™ Plugin

Conduit connector for reading data out of an Microsoft Excel™ 2007 and later `xlsx`
formatted files.

## Source

### Change Data Capture (CDC)
todo

### Testing

`go test ./...`

#### Position Handling

The row of the sheet is used as the "position".


### Record Keys

The row number is used at the Record "key".

### Configuration

| name       | part of | description          | required | default value |
|------------|---------|----------------------|----------|---------------|
| `filepath` | source  | The path to the file | true     |               |
| `sheet`    | source  | The worksheet        | true     |               |


### Known Limitations

* Only supports Microsoft Excel™ 2007 and later

