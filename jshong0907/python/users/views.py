from django.contrib.auth import authenticate, get_user_model
from django.contrib.auth.hashers import make_password
from django.http import JsonResponse
from django.shortcuts import render

from rest_framework.generics import CreateAPIView
from rest_framework.permissions import AllowAny, IsAuthenticated
from rest_framework.views import APIView
from rest_framework_simplejwt.views import TokenObtainPairView

from users.serializers import UserSerializer, JwtSerializer


class SignupView(CreateAPIView):
    permission_classes = [AllowAny]
    serializer_class = UserSerializer


class SigninView(TokenObtainPairView):
    serializer_class = JwtSerializer


class UserView(APIView):
    permission_classes = [IsAuthenticated]

    def get(self, request, *args, **kwargs):
        return JsonResponse(
            data={'hello': request.user.email}
        )