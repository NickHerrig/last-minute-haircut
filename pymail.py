import os
import smtplib
import ssl


sender_email = "nickherrigdeveloper@gmail.com"
password = os.environ['DEV_EMAIL_PASS']
port = 465


def send_email(send_to, subject, message):
    """Send an email from nickherrigdeveloper@gmail.com.

    Keyword arguments:
    send_to -- the receiver, can be an email or phone number
    subject -- the subject line
    message -- the contents

    """

    email = 'Subject: {}\n\n{}'.format(subject, message)
    context = ssl.create_default_context()

    with smtplib.SMTP_SSL("smtp.gmail.com", port, context=context) as server:
        server.login(sender_email, password)
        server.sendmail(sender_email, send_to, email)
