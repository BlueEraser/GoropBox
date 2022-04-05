from rest_framework.generics import CreateAPIView, RetrieveAPIView
from rest_framework.permissions import AllowAny, IsAuthenticated
from rest_framework_simplejwt.views import TokenObtainPairView

from users.serializers import UserSerializer, JwtSerializer


class SignupView(CreateAPIView):
    permission_classes = [AllowAny]
    serializer_class = UserSerializer


class SigninView(TokenObtainPairView):
    serializer_class = JwtSerializer


class UserView(RetrieveAPIView):
    permission_classes = [IsAuthenticated]
    serializer_class = UserSerializer

    def get_object(self):
        return self.request.user
