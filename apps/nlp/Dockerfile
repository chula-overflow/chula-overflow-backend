FROM python:3.9.15

WORKDIR /app

COPY apps/nlp/requirements.txt .

RUN pip install -r requirements.txt

COPY apps/proto /proto

COPY apps/nlp .

CMD ["python", "server.py"]