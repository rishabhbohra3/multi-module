#!/bin/bash

input_text="$1"  # Accept the input text as an argument

# Remove leading and trailing spaces, then remove '/' and replace spaces and newlines with nothing
# cleaned_text=$(echo "$input_text" | sed -e 's/^\s*//' -e 's/\s*$//' -e 's/\// /g' | tr -d '[:space:]')
cleaned_text=$(echo "$input_text" | sed -E 's|/(\s*)review|\1review|')

echo "$cleaned_text"

