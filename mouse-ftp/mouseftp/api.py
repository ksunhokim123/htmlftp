import json

from flask import Flask, request, jsonify
app = Flask(__name__)

def start(server, address):
    global ftp_server
    ftp_server = server
    app.run(host=address[0], port=address[1])

@app.route('/users', methods=['POST'])
def add_user():
    username = request.form.get('username')
    password = request.form.get('password')
    try:
        ftp_server.add_user(username, password)
    except ValueError as err:
        return jsonify(
            error= str(err)
        )
    return jsonify(
        error= 'success'
    )

@app.route('/users', methods=['GET'])
def get_users():
    return jsonify(
        error= 'success',
        users= ftp_server.get_users()
    )

@app.route('/users/<username>', methods=['DELETE'])
def remove_user(username):
    try:
        ftp_server.remove_user(username)
    except KeyError as err:
        # TODO
        return jsonify(
            error= str(err) + " doesn't exist"
        )
    return jsonify(
        error= 'success'
    )
