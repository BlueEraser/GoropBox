from rest_framework import serializers
from rest_framework_simplejwt.serializers import TokenObtainPairSerializer

from users.models import User
from users.services import UserService


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ('email', 'password', 'nick_name')
        extra_kwargs = {
            'password': {'write_only': True},
        }

    def create(self, validated_data):
        return UserService().create(**validated_data)


class JwtSerializer(TokenObtainPairSerializer):
    username_field = 'email'

    @classmethod
    def get_token(cls, user):
        token = super().get_token(user)
        token['email'] = user.email
        token['nick_name'] = user.email
        return token
