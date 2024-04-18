from . import front_blue
from flask import request
from flask import jsonify
from datetime import timedelta

from flask_jwt_extended import create_access_token
from flask_jwt_extended import get_jwt_identity
from flask_jwt_extended import jwt_required
from flask_jwt_extended import JWTManager


@front_blue.route(f"/info", methods=["GET"])
def info():
    return jsonify({"msg": "info..."})
