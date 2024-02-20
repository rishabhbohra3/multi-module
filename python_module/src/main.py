import logging
import logging.handlers

# Create a logger
logger = logging.getLogger('MySyslogLogger')
logger.setLevel(logging.INFO)

extra = {'logger_id': "1234"}

# syslog format: "<Timestamp> <Host> <AppName>[<PID>]: <Message>"
syslog_format = '%(asctime)s %(logger_id)s MyApplication[%(process)d]: %(message)s'

formatter = logging.Formatter(syslog_format, datefmt='%b %d %H:%M:%S')
formatter.default_msec_format = '%s.%03d'

handler = logging.StreamHandler()
handler.setFormatter(formatter)
logger.addHandler(handler)

logger.info('Initialized logger module', extra=extra)
