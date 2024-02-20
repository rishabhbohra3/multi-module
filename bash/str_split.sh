
#!/bin/bash

# Example comma-separated string
csv_string="apple,banana,orange,apple,grape,banana"

# Convert comma-separated string to set
IFS=', ' read -ra elements <<< "$csv_string"

echo "'${elements[@]}'"


# Create a set using associative array
declare -a unique_set

# Iterate through the array and add unique elements to the set
for element in "${elements[@]}"; do
  trimmed_element=$(echo "$element" | tr -d '[:space:]')
  if [[ ! " ${unique_set[@]} " =~ " $trimmed_element " ]]; then
    unique_set+=("$trimmed_element")
  fi
done

# Print the set
echo "Set:"
for element in "${unique_set[@]}"; do
  echo "$element"
done
