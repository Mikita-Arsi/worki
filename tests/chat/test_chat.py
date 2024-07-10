import json
import pytest
import requests

from tests.src.basic_classes.response import Response
from tests.src.enums.global_enums import GlobalErrorMessages
from tests.configuration import USERS_LIST, CHAT_LIST


def test_create_delete_by_id():
    data = {
        "name": "test_create_delete_by_id"
    }
    resp = requests.post(CHAT_LIST, json=data)
    assert resp.status_code == 201, GlobalErrorMessages.WRONG_STATUS_CODE
    chat_id = resp.json()["id"]
    resp = requests.delete(CHAT_LIST + str(chat_id))
    assert resp.status_code == 204, GlobalErrorMessages.WRONG_STATUS_CODE


def test_get_chats():
    resp = requests.post(CHAT_LIST, json={"name": "test_get_chats"})
    chat_id = resp.json()["id"]
    assert resp.status_code == 201, GlobalErrorMessages.WRONG_STATUS_CODE
    resp = requests.get(CHAT_LIST).json()
    if isinstance(resp, list):
        assert resp[len(resp) - 1]["name"] == "test_get_chats"
    else:
        assert resp["name"] == "test_get_chats"
    resp = requests.delete(CHAT_LIST + str(chat_id))
    assert resp.status_code == 204, GlobalErrorMessages.WRONG_STATUS_CODE


def test_get_chat_by_id():
    resp = requests.post(CHAT_LIST, json={"name": "test_get_chat_by_id"})
    chat_id = resp.json()["id"]
    assert resp.status_code == 201, GlobalErrorMessages.WRONG_STATUS_CODE
    resp = requests.get(CHAT_LIST + str(chat_id))
    assert resp.status_code == 200, GlobalErrorMessages.WRONG_STATUS_CODE
    assert resp.json()["id"] == chat_id
    resp = requests.delete(CHAT_LIST + str(chat_id))
    assert resp.status_code == 204, GlobalErrorMessages.WRONG_STATUS_CODE


def test_add_delete_user(get_user_builder, create_user):
    user = get_user_builder.build()
    resp = Response(create_user(user)).assert_status_code(201).response_json
    user_id = resp["id"]
    resp = requests.post(CHAT_LIST, json={"name": "test_add_delete_user"})
    assert resp.status_code == 201, GlobalErrorMessages.WRONG_STATUS_CODE
    chat_id = resp.json()["id"]
    chat_user_data = {
        "chat_id": chat_id,
        "from_id": user_id,
    }
    headers = {'Content-Type': 'application/json'}
    resp = requests.post(CHAT_LIST + "add", headers=headers, data=json.dumps(chat_user_data))
    assert resp.status_code == 204, GlobalErrorMessages.WRONG_STATUS_CODE
    resp = requests.delete(CHAT_LIST + "user", headers=headers, data=json.dumps(chat_user_data))




