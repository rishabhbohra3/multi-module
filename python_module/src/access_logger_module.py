from context_vars import logger_module

app_logger_module = logger_module.get()
app_logger_module.info("contest vars - logger module", extra={'logger_id': "logger_id"})
