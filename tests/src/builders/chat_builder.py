from faker import Faker


class UserBuilder:
    def __init__(self, users_id, messages_id):
        self.fake = Faker()
        self.result = {
            "users_id": self.fake,
            "messages_id": self.fake.password(),
        }

    def build(self):
        return self.result
