from . import login_blue
from flask import request
from flask import jsonify
from datetime import timedelta

from flask_jwt_extended import create_access_token
from flask_jwt_extended import get_jwt_identity
from flask_jwt_extended import jwt_required
from flask_jwt_extended import JWTManager

# 这里只是演示了10个路由，实际开发中可以根据需要对1000个路由进行类似的扩展


@login_blue.route(f"/login", methods=["POST"])
def login():
    name = request.json.get("username", None)
    pwd = request.json.get("password", None)
    if name is None or pwd is None:
        return jsonify({"msg": "Bad username or password"}), 401
    identity = f"{name}xxxx{pwd}"
    token = create_access_token(identity=identity, expires_delta=timedelta(days=10))
    return jsonify(token=token)
