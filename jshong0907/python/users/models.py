from django.db import models
from encrypted_model_fields.fields import EncryptedCharField
from django_extensions.db.models import TimeStampedModel


# Create your models here.
class User(TimeStampedModel):
    email = models.EmailField(
        verbose_name='이메일',
        max_length=50,
        unique=True,
    )
    password = EncryptedCharField(
        verbose_name='비밀번호',
        max_length=255,
    )
    nick_name = models.CharField(
        verbose_name='닉네임',
        max_length=10,
    )
