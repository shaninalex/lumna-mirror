#!/bin/bash

templ generate
go run ./app serve --config=./config/config.yaml
