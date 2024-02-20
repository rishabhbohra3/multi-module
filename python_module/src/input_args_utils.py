def parse_input_to_dict(input_string):
    # Split the input string by spaces to get each key-value pair
    pairs = input_string.split()
    
    # Initialize an empty dictionary
    result_dict = {}
    
    # Process each pair
    for pair in pairs:
        # Each pair is expected to be in the format "--key=value"
        # Remove the leading '--' and split by '=' to separate the key and value
        key, value = pair.lstrip('-').split('=', 1)
        
        # Add the key and value to the dictionary
        result_dict[key] = value
    
    return result_dict

# Example usage
input_string = "--cli=server --key=value --key2=value2"
output_dict = parse_input_to_dict(input_string)
print(output_dict)
