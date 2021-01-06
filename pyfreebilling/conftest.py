import pytest
from django.test import RequestFactory

from pyfreebilling.users.models import User
from pyfreebilling.users.tests.factories import UserFactory


@pytest.fixture(autouse=True)
def media_storage(settings, tmpdir):
    settings.MEDIA_ROOT = tmpdir.strpath


@pytest.fixture
def user() -> User:
    return UserFactory()

