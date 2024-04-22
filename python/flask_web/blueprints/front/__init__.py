from flask import Blueprint

front_blue = Blueprint("front", __name__)

from .views import *
