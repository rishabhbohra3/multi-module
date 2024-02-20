import logging
import uuid


class ContextFilter(logging.Filter):

    def filter(self, record):
        record.uuid = uuid.uuid4()  # Generate a unique log ID for each message
        return True

class CustomLogger:
    def __init__(self, name, level=logging.INFO):
        self.logger = logging.getLogger(name)
        self.logger.setLevel(level)

        ch = logging.StreamHandler()
        ch.setLevel(level)

        syslog_format = '%(asctime)s %(logger_id)s MyApplication[%(process)d]: %(message)s'
        # formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - [%(unique_id)s] %(message)s')
        ch.setFormatter(syslog_format)

        self.logger.addHandler(ch)

        self.unique_ids = set()

    def add_unique_id(self, unique_id):
        """Add a unique_id to the logger's context."""
        if unique_id not in self.unique_ids:
            for handler in self.logger.handlers:
                handler.addFilter(self._create_filter(unique_id))
            self.unique_ids.add(unique_id)

    def _create_filter(self, unique_id):
        """Create a filter that adds the unique_id to log records."""
        def filter_func(record):
            record.unique_id = unique_id
            return True
        return logging.Filter(filter=filter_func)

    def log(self, level, msg, unique_id=None, *args, **kwargs):
        """Log a message with a specified unique_id."""
        if unique_id:
            self.add_unique_id(unique_id)
        self.logger.log(level, msg, *args, **kwargs)


def main():
    custom_logger = CustomLogger('my_app')

    request_id_1 = str(uuid.uuid4())
    custom_logger.log(logging.INFO, 'This is an info message for request 1', unique_id=request_id_1)

    request_id_2 = str(uuid.uuid4())
    custom_logger.log(logging.INFO, 'This is an info message for request 2', unique_id=request_id_2)


if __name__ == '__main__':
    main()
