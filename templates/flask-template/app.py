from flask import Flask
from flask_cors import CORS
from view.api import api_bp

app = Flask(__name__)

CORS(app,
	supports_credentials=True,
	origins=[
		'http://localhost:5173',
		'http://127.0.0.1:5173'
	]
)

app.config.from_pyfile('config.py')

app.register_blueprint(api_bp)

if __name__ == '__main__':
	# remova 'debug=True' para prod
	app.run(debug=True)
