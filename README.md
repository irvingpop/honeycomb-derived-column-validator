# Honeycomb Derived Column Validator

A command-line tool to validate and test Honeycomb derived column expressions.

## Overview

This tool helps validate derived column expressions used in Honeycomb by parsing them and checking for syntax errors. It can be used to test expressions before adding them to your Honeycomb datasets.

## Installation
`go install github.com/irvingpop/honeycomb-derived-column-validator`

## Ideas for running in Terraform
1. As an external data source like: https://gist.github.com/irvingpop/968464132ded25a206ced835d50afa6b
2. Entirely outside of TF plan runs, like with Github Actions
