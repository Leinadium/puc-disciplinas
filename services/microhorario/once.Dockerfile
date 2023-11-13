FROM python:latest

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY *.py .

CMD ["python", "main.py", "-f"]
