#!/bin/bash

# Value to check (with leading/trailing spaces and newlines)
value_to_check="   
cherry   
"

# Trim leading and trailing spaces and newlines from the value
value_to_check=$(echo "$value_to_check" | tr -d '[:space:]')

echo $value_to_check
