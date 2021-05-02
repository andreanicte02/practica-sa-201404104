from flask import Flask
app = Flask (__name__)

@app.route('/')
def hello():
    return 'Andrea Vicente Campos - 201404104 - todogood'
app.run(host='0.0.0.0')