from dotenv import load_dotenv
import os

load_dotenv()

userNamedb = os.getenv("POSTGRES_USER")
password = os.getenv("POSTGRES_PASSWORD")
host = os.getenv("POSTGRES_HOST")
port = os.getenv("POSTGRES_PORT")
database_name = os.getenv("POSTGRES_DB")
lm_path = os.getenv("LM_STUDIO_API_URL")
lm_deepseek_model = os.getenv("LM_STUDIO_DEEPSEEK_MODEL")
lm_deepseek_timeout = float(os.getenv("LM_STUDIO_DEEPSEEK_TIMEOUT"))
