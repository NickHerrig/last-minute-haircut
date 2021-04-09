FROM python:3.8.4

COPY requirements.txt /

RUN pip install -r requirements.txt

COPY haircut.py /

CMD [ "uvicorn", "haircut:app", "--host", "0.0.0.0"]