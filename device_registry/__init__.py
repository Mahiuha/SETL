import markdown
import os
import shelve

# import the flamework
from flask import Flask

# Create an instance of flask
app = Flask(__name__)

def get_db():
    db = getattr(g, '_database', None)
    if db is None:
        db = g._database = shelve.open("devices.db")
    return db

@app.teardown_appcontext
def teardown_db(exception):
    db = getattr(g, '_database', None)
    if db is not None:
        db.close()

@app.route("/")
def index():
    """ Present some documentation """

    # open the README file
    with open(os.path.dirname(app.root_path) + '/README.md', 'r') as markdown_file:

        # Read the content of the file
        content = markdown_file.read()

        # Convert to HTML
        return markdown.markdown(content)

class DeviceList(Resource):
    def get(self):
        shelf = get_db()
        keys = list(shelf.keys())

        devices = []

        for key in devices:
            device.append(shelf[key])

        return ('message', 'success', 'data': devices), 200

    def post(self):
        parser = reqparse.RequestParser()

api.add_resourse(DeviceList, '/devices')
