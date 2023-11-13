FROM python:3.11-alpine

WORKDIR /historico

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN pip install --upgrade pip
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
RUN pip install gunicorn

COPY . .

EXPOSE 5000

ENV POSTGRES_CONN "postgresql+psycopg2://tcc:tcc@127.0.0.1:5432/tcc"

CMD ["gunicorn", "--bind", "0.0.0.0:5001", "main:app"]
