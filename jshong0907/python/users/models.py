from django.db import models
from encrypted_model_fields.fields import EncryptedCharField
from django_extensions.db.models import TimeStampedModel
from django.contrib.auth.models import AbstractBaseUser, UserManager


class User(AbstractBaseUser):
    USERNAME_FIELD = 'email'

    email = models.EmailField(
        verbose_name='이메일',
        max_length=50,
        unique=True,
    )
    nick_name = models.CharField(
        verbose_name='닉네임',
        max_length=10,
    )

    objects = UserManager()
