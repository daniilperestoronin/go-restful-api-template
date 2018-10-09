# Go RESTful Api template

Simplest template of RESTful service written in [Golang](https://golang.org/).
Used only standard golang libraries.

## Getting Started

### Building

```bash
docker build -t go-record-service .
```

### Running

```bash
docker run -p 8080:8080 -e DB_DRIVER="postgres" -e DB_DATA_SOURCE="host=192.168.42.176 port= 5432 user=postgres password=postgres dbname=records sslmode=disable" --name go-record-service go-record-service
```

## Running the tests

Explain how to run the automated tests for this system

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

- [Daniil Perestoronin](https://github.com/daniilperestoronin)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details