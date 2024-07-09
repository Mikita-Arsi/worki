import pytest
import requests

from tests.configuration import USERS_LIST
from tests.src.builders.user_builder import UserBuilder


@pytest.fixture
def get_user_builder():
    return UserBuilder()


@pytest.fixture
def get_users():
    response = requests.get(USERS_LIST)
    return response


def _create_user(user):
    response = requests.post(USERS_LIST, json=user)
    return response


def _delete_user(user_id):
    response = requests.delete(USERS_LIST + str(user_id))
    return response


@pytest.fixture
def create_user():
    return _create_user


@pytest.fixture
def delete_user():
    return _delete_user


print(USERS_LIST + str(1))
