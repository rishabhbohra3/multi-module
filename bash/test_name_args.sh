#!/bin/bash

# This script demonstrates how to use field names for arguments

# Function to display usage information
display_usage() {
    echo "Usage: $0 -name <name> -age <age> -city <city>"
}

# Check if there are no arguments
if [ "$#" -eq 0 ]; then
    display_usage
    exit 1
fi

echo $#
echo $key
echo $1
echo $2
# Parse command-line arguments
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -name|--name)
	    name="$2"
            shift
            ;;
	--name=*|-name=*)
	    name="${1#*=}"
	    shift
	    ;;
	--age=*|-age=*)
	    age="${1#*=}"
            shift
            ;;
        -age|--age)
            age="$2"
            shift
            ;;
	--city=*|-city=*)
            city="${1#*=}"
            shift
            ;;
        -city|--city)
            city="$2"
            shift
            ;;
        *)
            # Unknown option
            echo "Error: Unknown option: $1"
            display_usage
            exit 1
            ;;
    esac
    shift
done

# Display the parsed values
echo "Name: $name"
echo "Age: $age"
echo "City: $city"

