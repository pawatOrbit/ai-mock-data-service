# AI Mock Data Service

A FastAPI-based backend service for generating AI-powered mock data from table schemas using LM Studio. It connects to a PostgreSQL database and integrates with a language model to create realistic mock data for development and testing.

---

## 🧰 Features

- Generate SQL `INSERT` mock data from table definitions.
- Connect to LM Studio via HTTP API.
- Asynchronous PostgreSQL access via `asyncpg` and `databases`.
- Configurable via `.env` file.
- Clean logging with color-coded output.
- Modular folder structure with separation of concerns (HTTP, DB, Client, etc.).

---

## 📦 Requirements

- Python 3.9+
- PostgreSQL (local or remote)
- LM Studio running and accessible via HTTP

---


## ⚙️ Setup Instructions

### 1. Clone the Repository

```
git clone https://github.com/pawatOrbit/ai-mock-data-service.git
cd ai-mock-data-service
```

### 2. Create and Activate Virtual Environment

```
python3 -m venv env
source env/bin/activate 
```

### 3. Install Dependencies

```
pip install -r requirements.txt
```

### 4. Set Up Environment Variables

Create a ```.env``` file in the root directory with the following variables:
```
# PostgreSQL Config
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=mock_data_db

# LM Studio Config
LM_STUDIO_API_URL=http://localhost:1234/v1/completions
LM_STUDIO_DEEPSEEK_MODEL=deepseek-model-name
LM_STUDIO_DEEPSEEK_TIMEOUT=30
```

## 🚀 Running the App

### Run using Uvicorn (Development)
```
uvicorn app.main:app --reload
```

The API will be available at:
http://localhost:8000/docs – Swagger UI

Developed by Pawat.t – this project is part of HACKATHON on Orbit digital Merchant Backend team 