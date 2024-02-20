from contextvars import ContextVar
from main import logger

logger_module = ContextVar('logger_module')
token = logger_module.set(logger)
