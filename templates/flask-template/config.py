from dotenv import load_dotenv
import os.path

load_dotenv()

SECRET_KEY = os.getenv("SECRET_KEY")

BASE_DIR = os.path.abspath(os.path.dirname(__file__))
