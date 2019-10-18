# Smarkets tools

A mishmash of tooling for smarkets API interactions

## Getting Started

Some functions require authentication to access to API; this is acheived by reading from a `.SMARKETS_TOKEN` file in the home directory

### Example uses

Tools can be used to store quote/bid data into a mysql database

```
connectionString := "mysql_user:mysql_password@/mysql_table? charset=utf8parseTime=True&loc=Local"  
```
```
marketID := "xxxxxxx"
```
```
smarketstools.InsertQuotesMysql(marketID, connectionString)
```

## Contributing

When it exists; please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details.


## Authors

* **Howard JOnes** - *Initial work* - [howardjonestm](https://github.com/howardjonestm)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

