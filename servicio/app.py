from flask import Flask, request
from flask_cors import CORS
from mysql import Database
import simplejson as json
import os

app = Flask(__name__)
cors = CORS(app, resources={r"/*": {"origin": "*"}})


@app.route('/', methods=['GET'])
def check():
    return str(os.getenv('SERVNAME'))+": Todo good"


@app.route('/add_reporte', methods=['POST'])
def add_reporte():
    # {"carnet":0, "nombre":"nombre", "curso":"curso"}
    content = request.get_json()
    content['procesado'] = str(os.getenv('SERVNAME'))
    db = Database()
    res = db.add_reporte(content)
    res['atendido'] = str(os.getenv('SERVNAME'))
    return json.dumps(res)


@app.route('/get_reporte', methods=['GET'])
def get_reporte():
    reporte = request.args.get('reporte')
    res = {}
    res['atendido'] = str(os.getenv('SERVNAME'))
    db = Database()
    res['data'] = db.get_reporte(reporte)
    return json.dumps(res, default=str)


@app.route('/get_lista_reporte', methods=['GET'])
def get_lista_reporte():
    carnet = request.args.get('carnet')
    res = {}
    res['atendido'] = str(os.getenv('SERVNAME'))
    db = Database()
    res['data'] = db.get_lista_reporte(carnet)
    return json.dumps(res, default=str)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, threaded=True, use_reloader=True)
