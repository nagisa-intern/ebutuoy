# Dosukoi
from flask import Flask
import subprocess

app = Flask(__name__)  # Standard Flask app

@app.route("/")        # Standard Flask endpoint
def hello_world():
    return "Hello, World!"

@app.route("/webhook", methods=['POST'])
def on_push():
    try:
      res = subprocess.call('git pull ; git checkout -f infra ; docker-compose up -d') # /usr/local/bin/
      print(res)
    except Exception as e:
      print(e)
    return "OK"

if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=80)