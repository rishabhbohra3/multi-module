import datetime
import time

# Get the current date and time as the start time
start_time = datetime.datetime.now()
print(f"Start Time: {start_time}")

# Simulate some operation or delay (e.g., wait for 5 seconds)
time.sleep(75)  # Wait for 5 seconds

# Get the current date and time again as the end time
end_time = datetime.datetime.now()
print(f"End Time: {end_time}")

# Calculate the difference between the end time and the start time
time_diff = end_time - start_time

# Print the difference
print(f"Difference: {time_diff}")

# If you want to print the difference in a specific format (e.g., hours, minutes, seconds)
hours, remainder = divmod(time_diff.total_seconds(), 3600)
minutes, seconds = divmod(remainder, 60)
print(f"Difference: {int(hours)} hours, {int(minutes)} minutes, {int(seconds)} seconds")
