import pytest
import requests

from tests.src.basic_classes.response import Response
from tests.src.enums.global_enums import GlobalErrorMessages
from tests.configuration import USERS_LIST, CHAT_LIST


def test_create():
    resp = requests.get(CHAT_LIST).json()
    chat_id = resp[0]["id"]
    if isinstance(resp[0]["users_id"], list):
        from_id = resp[0]["users_id"][0]
    else:
        from_id = resp[0]["users_id"]

    from_id = resp
    data = {
        "chat_id": chat_id,
        "from_id": from_id,
        "text": "Lorem Ipsum"
    }
    resp = requests.post(CHAT_LIST + str(chat_id) + "/messages", json=data)
    assert resp.status_code == 201, GlobalErrorMessages.WRONG_STATUS_CODE


def test_get_messages():
    resp = requests.get(CHAT_LIST).json()
    chat_id = resp[0]["id"]
    resp = requests.get(CHAT_LIST + str(chat_id))
    assert resp.status_code == 200


def test_delete():
    resp = requests.get(CHAT_LIST).json()
    chat_id = resp[0]["id"]
    resp = requests.get(CHAT_LIST + str(chat_id)).json()
    message_id = resp[0]["id"]
    resp = requests.delete(CHAT_LIST + str(chat_id) + "/message/" + str(message_id))
