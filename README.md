[![Coverage Status](https://coveralls.io/repos/github/babbageLabs/brinch/badge.png?branch=master)](https://coveralls.io/github/babbageLabs/brinch?branch=master&kill_cache=1)

# About

Brinch is a cli tool used to generate [Json Schema](https://json-schema.org/specification.html) assets derived from postgres objects. The
objective is to use the [Json Schema](https://json-schema.org/specification.html) assets to create
a fully functional and tested back end server as well as basic ui Forms/pages

## Prerequisites

A working knowledge of golan

## Getting started

create a ```.brinch.yaml``` file in your directory which serves as the config file.
refer to ```.brinch.exaple.yaml```

## TODO

- [ ] Refactor Db initialization feature
- [ ] Unit test the db initialization logic
- [x] Unit test lib
- [ ] Work on mapping postgres native types to json schema properties
- 