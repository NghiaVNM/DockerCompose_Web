from flask import Flask
from pymongo import MongoClient

app = Flask(__name__)

client = MongoClient('mongo', 27017)
db = client['database']
users_collection = db['users']

@app.route('/')
def index():
    try: 
        users = list(users_collection.find({}, {'_id': 0}))
        return f"""
            <html>
                <head>
                    <title>Python Web App</title>
                </head>
                <body>
                    <h1>Hello World from Python</h1>
                    <h2>User List</h2>
                    <ul>
                        {''.join([f"<li>{user['name']} - {user['age']}</li>" for user in users])}
                    <ul>
                </body>
            </html>         
        """
    except Exception as e:
        return str(e), 500
    
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3001)