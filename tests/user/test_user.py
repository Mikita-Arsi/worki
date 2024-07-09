import pytest
import requests
import time

from tests.src.basic_classes.response import Response
from tests.src.enums.global_enums import GlobalErrorMessages
from tests.configuration import USERS_LIST


def test_create_delete_by_id(create_user, get_user_builder, delete_user, get_users):
    user = get_user_builder.build()
    resp = Response(create_user(user)).assert_status_code(201).response_json
    assert resp['firstname'] == user['firstname'], GlobalErrorMessages.CREATE_ERROR
    assert resp['username'] == user['username'], GlobalErrorMessages.CREATE_ERROR
    user_id = resp['id']
    resp = delete_user(user_id)
    assert resp.status_code == 204, GlobalErrorMessages.WRONG_STATUS_CODE
    deleted_user = requests.get(USERS_LIST + str(user_id))
    assert deleted_user.status_code == 500, GlobalErrorMessages.WRONG_STATUS_CODE
    # if isinstance(users_list, list) and len(users_list) >= user_id:
    #     deleted_user = users_list[user_id - 1]
    #     assert deleted_user['id'] != user_id, GlobalErrorMessages.DELETE_ERROR
    # elif isinstance(users_list, dict):
    #     assert users_list['id'] != user_id, GlobalErrorMessages.DELETE_ERROR



# def test_get_users(get_users, get_user_builder, create_user, delete_user_id):
#     user = get_user_builder.build()
#     print(user)
#     resp = Response(create_user(user)).assert_status_code(201).response_json
#     print(resp)
#     print(Response(get_users).assert_status_code(200).response_json)
#     resp = delete_user_id(resp['id'])
#     assert resp.status_code == 204, GlobalErrorMessages.WRONG_STATUS_CODE


def test_get_delete_user_by_username(get_user_builder, create_user, delete_user):
    user = get_user_builder.build()
    resp = Response(create_user(user)).assert_status_code(201).response_json
    user_id = resp['id']
    resp = requests.get(USERS_LIST + user['username'],)
    assert resp.status_code == 200, GlobalErrorMessages.WRONG_STATUS_CODE
    resp = resp.json()
    assert resp['firstname'] == user['firstname'], GlobalErrorMessages.CREATE_ERROR
    assert resp['username'] == user['username'], GlobalErrorMessages.CREATE_ERROR
    resp = delete_user(user['username'])
    assert resp.status_code == 204, GlobalErrorMessages.WRONG_STATUS_CODE
    users_list = requests.get(USERS_LIST).json()
    if isinstance(users_list, list) and len(users_list) >= user_id:
        deleted_user = users_list[user_id - 1]
        assert deleted_user['id'] != user_id, GlobalErrorMessages.DELETE_ERROR
    elif isinstance(users_list, dict):
        assert users_list['id'] != user_id, GlobalErrorMessages.DELETE_ERROR


