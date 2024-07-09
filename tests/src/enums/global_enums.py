from enum import Enum


class GlobalErrorMessages(Enum):
    WRONG_STATUS_CODE = "Received status code is not equal to expected."
    DELETE_ERROR = "Deleted user still exists"
    CREATE_ERROR = "User has not been created"