from flask import Blueprint

backend_blue = Blueprint("backend", __name__)

from .views import *
