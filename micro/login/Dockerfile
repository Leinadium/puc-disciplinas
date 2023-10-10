FROM python:3.11-alpine

WORKDIR /login

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN pip install --upgrade pip
COPY requirements.txt /login/requirements.txt
RUN pip install -r requirements.txt
RUN pip install gunicorn

COPY main.py /login/main.py

EXPOSE 5000

CMD ["gunicorn", "--bind", "0.0.0.0:5000", "main:app"]

