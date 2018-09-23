# simple formatter
a golang package for parsing a format expression and outputting strings in that format in a relatively efficient way

## Examples
### Simple Usage
```go
    fields:=string{"username","fullname"}

    sfp,err:= simpleformatter.NewParsedFormat(fields,"hi user {fullname}")
    if err!=nil{
            fmt.Println(err)
            return 
    }

    for _,record := range records{
        fmt.Println(sftp.GetFormattedValue(record))
    }

```
### Advanced Usage (more then one section)
```go
    defaultSectionDefinition := SectionDefinition{
        FieldNames: []string{"record.name"},
    }
	sections := []*SectionDefinition{
        &SectionDefinition{SectionName: "BEGIN", FieldNames: []string{}]},
        &SectionDefinition{SectionName: "END", FieldNames: []string{}},
    }

	err := NewParsedFormatWithSections("hello BEGIN{see the following list } {record.name}  END{ footer note }", &defaultSectionDefinition, sections...)
    if err!=nil{
        fmt.Println(err)
        return 
    }

    //Output the BEGIN section (contains no usage fields)
    fmt.Println(sections[0].GetFormattedValue(nil))
    //Output the "default" section
    for _,record := range records{
        fmt.Println(defaultSectionDefinition.GetFormattedValue(record))
    }
    //Output the END section (contains no usage fields)
    fmt.Println(sections[1].GetFormattedValue(nil))
```
## Installing
```shell
go get github.com/rhysbryasnt/simpleformatter
```
then add  github.com/rhysbryasnt/simpleformatter to your imports 
## running the tests
there is 98% test coverage to run the tests execute the following
```shell
go test
```
## Contributing
want to Contribute sure pull requests are welcome. I just ask that your raise an issue first
## License
this project is licensed under LGPL
