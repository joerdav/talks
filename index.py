import flask
import os

# If `entrypoint` is not defined in app.yaml, App Engine will look for an app
# called `app` in `main.py`.
app = flask.Flask(__name__)


@app.route("/")
def hello():
    """Return a friendly HTTP greeting."""
    dirs = [f for f in os.listdir('./dist') if os.path.isdir('./dist/'+f)]

    return flask.render_template('index.html', talks=dirs)


if __name__ == "__main__":
    # Used when running locally only. When deploying to Google App
    # Engine, a webserver process such as Gunicorn will serve the app. This
    # can be configured by adding an `entrypoint` to app.yaml.
    app.run(host="localhost", port=8080, debug=True)

