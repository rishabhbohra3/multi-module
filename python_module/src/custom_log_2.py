import logging
import logging.handlers
import os
import uuid
import contextvars
from concurrent.futures import ThreadPoolExecutor


logger_id_var = contextvars.ContextVar('logger_id', default=None)
thread_pool = ThreadPoolExecutor(max_workers=2)
logger_file_path = '/Users/rishabhbohra/Documents/bito-code/MyProjects/custom_logger/cra.logs'


class CustomLogger(logging.Logger):
    def __init__(self, name, level=logging.NOTSET):
        super().__init__(name, level)

    def _log(self, level, msg, args, exc_info=None, extra=None, stack_info=False):
        if extra is None:
            extra = {}

        # Add the unique_id to the extra dict
        logger_id = logger_id_var.get()
        if logger_id:
            extra['logger_id'] = logger_id
        else:
            extra['logger_id'] = ""

        # Call the original _log method with the modified extra dict
        super()._log(level, msg, args, exc_info, extra=extra, stack_info=stack_info)


# To make the logging system use our CustomLogger, we need to override the Logger class in the logging manager
logging.setLoggerClass(CustomLogger)

# Create a custom formatter that includes the unique_id in the log message
syslog_format = '%(asctime)s %(levelname)s  %(name)s[%(process)d] %(funcName)s [%(logger_id)s]: %(message)s'
formatter = logging.Formatter(syslog_format)

# Set up a console handler using this formatter
# handler = logging.StreamHandler()

handler = logging.handlers.TimedRotatingFileHandler(filename=logger_file_path,
                                                    when='midnight',
                                                    backupCount=10)
# handler = logging.handlers.RotatingFileHandler(filename=logger_file_path, maxBytes=5*1024*1024,
#                                                backupCount=5)
handler.setFormatter(formatter)

logger = logging.getLogger("my_custom_logger")
logger.setLevel(logging.DEBUG)
logger.addHandler(handler)


def print_logs(log_id):
    if log_id:
        logger_id_var.set(log_id)
    # Use the context to log messages
    logger.info("Test log from print function")


# Example usage
if __name__ == "__main__":
    logger_id = "12345"
    logger_id_var.set(logger_id)
    logger.info("This is a test log message")
    thread_pool.submit(print_logs, "thread1-abcd")
    thread_pool.submit(print_logs, "thread2-qwerty")
    thread_pool.submit(print_logs, None)


