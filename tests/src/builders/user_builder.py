from faker import Faker


class UserBuilder:
    def __init__(self):
        self.fake = Faker()
        self.result = {
            "firstname": self.fake.first_name(),
            "password": self.fake.password(),
            "username": self.fake.user_name()
        }

    def build(self):
        return self.result
