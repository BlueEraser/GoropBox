from django.contrib.auth.hashers import make_password

from users.models import User


class UserService:
    @staticmethod
    def create(
            email: str,
            password: str,
            nick_name: str,
    ) -> User:
        hashed_password = make_password(password)
        return User.objects.create(
            email=email,
            password=hashed_password,
            nick_name=nick_name,
        )
