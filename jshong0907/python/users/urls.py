from django.urls import path

from users.views import SignupView, SigninView, UserView

urlpatterns = [
    path('auth/signup/', SignupView.as_view()),
    path('auth/signin/', SigninView.as_view()),
    path('user/', UserView.as_view()),
]