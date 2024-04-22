from flask import Flask, g
import time
from flask_jwt_extended import JWTManager, jwt_required, get_jwt_identity

app = Flask(__name__)

from blueprints.login import login_blue
from blueprints.front import front_blue
from blueprints.backend import backend_blue


def register_request_handlers(blueprint):
    @blueprint.before_request
    @jwt_required()
    def before_request_func():
        pass

    @blueprint.before_request
    def start_timer():
        g.t_start = time.time()  # 存储开始时间

    @blueprint.after_request
    def calculate_duration(response):
        if hasattr(g, "t_start"):
            duration = time.time() - g.t_start  # 计算请求耗时
            response.headers["X-Request-Duration-ms"] = f"{duration * 1000:.2f}"
            # print(f"Duration ms:{duration * 1000:.2f})")
        return response


app.config["JWT_SECRET_KEY"] = "your-secret-key"  # Change this to your real secret key
jwt = JWTManager(app)

register_request_handlers(front_blue)
register_request_handlers(backend_blue)

app.register_blueprint(login_blue, url_prefix="")
app.register_blueprint(front_blue, url_prefix="/front")
app.register_blueprint(backend_blue, url_prefix="/backend")


if __name__ == "__main__":
    app.run(debug=True)
