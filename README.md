# eighty

eighty is an HTTP server for performing dynamic HTTP routing.

## Features

* graceful shutdown of the HTTP server
* zero downtime restarts (through [Einhorn](https://github.com/stripe/einhorn))

## Requirements

* Go 1.3+
* Git

### Installing Dependencies

    brew install go git

## Building

Run `make` and look for `bin/eighty`.

## Running

To run on port 80:

    bin/eighty

You can change the port through `-bind`:

    bin/eighty -bind=:8080

If you need to listen on a specific interface:

    bin/eighty -bind=127.0.0.1:80

To enable zero downtime restarts, use Einhorn:

    einhorn -b 127.0.0.1:80,r -m manual bin/eighty -bind=einhorn@0
