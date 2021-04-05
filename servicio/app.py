from flask import Flask, request
from flask_cors import CORS
from mysql import Database
import simplejson as json
import os

app = Flask(__name__)
cors = CORS(app, resources={r"/*": {"origin": "*"}})


@app.route('/', methods=['GET'])
def check():
    return str(os.getenv('SERVNAME'))+": hola mundo :)"


@app.route('/get_simple', methods=['GET'])
def get_simple():
    reporte = ''
    res = {}
    res['atendido'] = str(os.getenv('SERVNAME'))
    db = Database()
    res['data'] = db.get_simple(reporte)
    return json.dumps(res, default=str)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, threaded=True, use_reloader=True)
