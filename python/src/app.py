import os
from flask import Flask, render_template
import firebase_admin
from firebase_admin import credentials, firestore

app = Flask(__name__)

# Initialize Firestore DB
cred = credentials.Certificate("./ServiceAccount.json")
default_app = firebase_admin.initialize_app(cred)

db = firestore.client()

@app.route('/')
def index():
    return render_template('index.html')

port = int(os.environ.get('PORT', 8080))
if __name__ == '__main__':
    app.run(threaded=True, host='localhost', port=port)

