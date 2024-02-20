import logging

# Create a custom logger
logger = logging.getLogger('my_custom_logger')
logger.setLevel(logging.DEBUG)  # Set the logging level

# Create formatter for console output
console_formatter = logging.Formatter('%(name)s - %(levelname)s - %(message)s')

# Create formatter for file output
file_formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s', datefmt='%Y-%m-%d %H:%M:%S')

# Create console handler and set level to debug
console_handler = logging.StreamHandler()
console_handler.setLevel(logging.DEBUG)
console_handler.setFormatter(console_formatter)  # Assign console formatter

# Create file handler and set level to error
file_handler = logging.FileHandler('error.log')
file_handler.setLevel(logging.ERROR)
file_handler.setFormatter(file_formatter)  # Assign file formatter

# Add handlers to the logger
logger.addHandler(console_handler)
logger.addHandler(file_handler)

# Example log messages
logger.debug('This is a debug message')
logger.info('This is an info message')
logger.warning('This is a warning message')
logger.error('This is an error message')
logger.critical('This is a critical message')
logger.removeHandler(file_handler)
logger.critical('This is a critical message 2')
