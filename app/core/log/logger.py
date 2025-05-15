# app/core/logging_config.py

import logging
import colorlog

def setup_logging():
    handler = colorlog.StreamHandler()

    formatter = colorlog.ColoredFormatter(
        fmt="%(asctime)s - %(log_color)s%(levelname)-8s%(reset)s - %(name)s - %(message)s",
        log_colors={
            "DEBUG":    "cyan",
            "INFO":     "green",
            "WARNING":  "yellow",
            "ERROR":    "red",
            "CRITICAL": "bold_red",
        },
        reset=True
    )

    handler.setFormatter(formatter)

    root_logger = logging.getLogger()
    root_logger.setLevel(logging.INFO)
    root_logger.handlers = [handler]

    logging.getLogger("databases").setLevel(logging.WARNING)
