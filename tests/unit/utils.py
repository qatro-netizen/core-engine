# utils.py

import os
import logging
import re
import json
from datetime import datetime

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def is_valid_file_path(file_path: str) -> bool:
    """
    Check if the file path exists and is not empty.
    """
    return os.path.exists(file_path) and file_path != ''

def get_current_datetime() -> str:
    """
    Return the current date and time in ISO format.
    """
    return datetime.utcnow().isoformat() + 'Z'

def load_json_file(file_path: str) -> dict:
    """
    Load JSON data from a file.
    """
    try:
        with open(file_path, 'r') as file:
            return json.load(file)
    except json.JSONDecodeError as e:
        logger.error(f'Failed to load JSON file: {e}')
        return {}
    except FileNotFoundError as e:
        logger.error(f'JSON file not found: {e}')
        return {}

def clean_string(input_string: str) -> str:
    """
    Remove leading and trailing whitespace from a string and replace multiple spaces with a single space.
    """
    return re.sub(r'\s+', ' ', input_string).strip()