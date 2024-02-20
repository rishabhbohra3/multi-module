import logging
import logging.handlers
import os
import uuid
import contextvars
from concurrent.futures import ThreadPoolExecutor


default_log_level = 'INFO'
logger_file_path = '/Users/rishabhbohra/Documents/bito-code/MyProjects/custom_logger/cra.logs'
# Define a context variable for the logger ID
logger_id_var = contextvars.ContextVar('logger_id', default=None)
thread_pool = ThreadPoolExecutor(max_workers=2)


def generate_logger_id():
    return uuid.uuid4()


class CustomLogger:
    def __init__(self, log_path, application_name, **kwargs):
        self.level = kwargs.get('level', default_log_level)
        self.log_path = log_path
        self.application_name = application_name

        logger = logging.getLogger('CRASyslogLogger')
        logger.setLevel(default_log_level)
        handler = logging.handlers.RotatingFileHandler(filename=logger_file_path,
                                                       maxBytes=5 * 1024 * 1024,
                                                       backupCount=5)
        syslog_format = '%(asctime)s %(levelname)s  %(name)s[%(process)d] %(funcName)s %(message)s'
        formatter = logging.Formatter(syslog_format, datefmt='%b %d %H:%M:%S')
        handler.setFormatter(formatter)
        if not logger.handlers:
            logger.addHandler(handler)
        self.logger_module = logger


    def get_logger(self):
        # Generate or retrieve the logger ID from the context
        logger_id = logger_id_var.get()
        if logger_id is None:
            logger_id = generate_logger_id()
            logger_id_var.set(logger_id)

        logger_name = str(logger_id)
        logger = logging.getLogger(logger_name)
        logger.setLevel(logging.INFO)

        handler = logging.StreamHandler()

        # Create a rotating file handler to implement log retention
        # log_file = os.path.join(self.log_path, 'cra.log')
        # handler = logging.handlers.RotatingFileHandler(log_file, maxBytes=5*1024*1024, backupCount=5)  # 5 MB per file, 5 files max

        # Define log format
        syslog_format = '%(asctime)s %(levelname)s  %(name)s[%(process)d] %(funcName)s %(message)s'
        formatter = logging.Formatter(syslog_format, datefmt='%b %d %H:%M:%S')
        handler.setFormatter(formatter)

        # Check if the handler is already added to avoid duplicate handlers
        if not logger.handlers:
            logger.addHandler(handler)

        return logger

    def log_message(self, level, message):
        logger = self.get_logger()
        logger_id = logger_id_var.get()
        formatted_message = f"[{logger_id}]: {message}"
        if level == "INFO":
            logger.info(formatted_message)
        elif level == "ERROR":
            logger.error(formatted_message)
        elif level == "WARNING":
            logger.warning(formatted_message)
        # Add more log levels as required


def print_logs(log_level, message, custom_logger, logger_id):
    # Set logger ID in context for a new operation or request
    logger_id_var.set(logger_id)

    # Use the context to log messages
    custom_logger.log_message(level=log_level, message=message)


def main():
    log_path = "/Users/rishabhbohra/Documents/bito-code/MyProjects/custom_logger/src"
    application_name = "YourApplicationName"
    custom_logger = CustomLogger(log_path, application_name)

    thread_pool.submit(print_logs, "INFO", "", custom_logger, "123456")
    thread_pool.submit(print_logs, "INFO", "", custom_logger, "12345612345")


if __name__ == '__main__':
    main()