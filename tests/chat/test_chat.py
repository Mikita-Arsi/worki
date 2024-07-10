import pytest
import requests

from tests.src.basic_classes.response import Response
from tests.src.enums.global_enums import GlobalErrorMessages
from tests.configuration import USERS_LIST, CHAT_LIST


def test_create():
    resp = requests.get(USERS_LIST).json()
    user_id1 = resp[0]["id"]
    user_id2 = resp[1]["id"]

    from_id = resp
    data = {
        "id": [
            user_id1, user_id2
        ],
    }
    resp = requests.post(CHAT_LIST, json=data)
    assert resp.status_code == 201, GlobalErrorMessages.WRONG_STATUS_CODE


def test_get_chats():
    resp = requests.get(CHAT_LIST).json()
    assert resp.status_code == 200


def test_add_user():
    data = {
        "id": 1
    }
    resp = requests.post(CHAT_LIST + "addUser", json=data)
    assert resp.status_code == 200


def test_remove_user():
    data = {
        "id": 1
    }
    resp = requests.post(CHAT_LIST + "removeUser", json=data)
    assert resp.status_code == 200


def test_get_chat_users():
    resp = requests.get(CHAT_LIST + str(0) + "/users").json()
    assert resp.status_code == 200


def test_get_chat_messages():
    resp = requests.get(CHAT_LIST + str(0) + "/messages").json()
    assert resp.status_code == 200


def test_delete():
    resp = requests.get(CHAT_LIST).json()
    chat_id = resp[0]["id"]
    resp = requests.delete(CHAT_LIST + str(chat_id))
    assert resp.status_code == 204




