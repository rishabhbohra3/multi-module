#!/bin/bash

# This script demonstrates how to use command-line arguments

# Check if at least one argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <arg1> [arg2] [arg3] ..."
    exit 1
fi

# Access and display the provided arguments
echo "Number of arguments: $#"
echo "Script name: $0"
echo "First argument: $1"
echo "Second argument: $2"
echo "All arguments: \$@ -> $@"

# Loop through all arguments
echo "Loop through arguments using a for loop:"
for arg in "$@"; do
    echo "$arg"
done

