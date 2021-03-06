import json
from datetime import datetime, timedelta

import requests

genbook_url = 'https://www.genbook.com/bookings/api/serviceproviders'
barbers_urls = {
    'jordan':  f'{genbook_url}/30230662/services/989056738/resources/989056742?',
    'pete':    f'{genbook_url}/31191440/services/10282291592/resources/10282190294',
    'brandon': f'{genbook_url}/30377943/services/2394050193/resources/2394025610',
    'luis':    f'{genbook_url}/30250062/services/1173749692/resources/1173749696',
    'zach':    f'{genbook_url}/30302725/services/1547629284/resources/1547629288',
    'paul':    f'{genbook_url}/30309745/services/1603733980/resources/1603733984',
    'kegan':   f'{genbook_url}/30352805/services/2098565278/resources/2098565282?',
}


def get_this_weeks_appointments(barbers_datetime_appointments):
    barbers_formated_appointments = {}
    week_from_today = datetime.today() + timedelta(days=7)
    for barber, appointments in barbers_datetime_appointments.items():
        last_minute_appointments = [apt.strftime('%b/%d/%Y')
                                    for apt in appointments
                                    if apt < week_from_today]
        barbers_formated_appointments[barber] = last_minute_appointments

    return barbers_formated_appointments


def format_appointments(barbers_appointments):
    barbers_formated_appointments = {}
    for barber, appointments in barbers_appointments.items():
        formated_appointments = [datetime.strptime(apt[:-1], '%Y%m%d')
                                 for apt in appointments]
        barbers_formated_appointments[barber] = formated_appointments

    return barbers_formated_appointments


def parse_barbers_appointments(barbers_json):
    barbers_appointments = {}
    for barber, api_json in barbers_json.items():
        barbers_appointments[barber] = api_json['dates']

    return barbers_appointments


def main():
    barbers_api_responses = {}
    for barber, url in barbers_urls.items():
        try:
            api_response = requests.get(url)
        except Exception as e:
            raise e
        try:
            barbers_api_responses[barber] = api_response.json()
        except json.JSONDecodeError as e:
            raise e

    barbers_appointments = parse_barbers_appointments(barbers_api_responses)
    barbers_formated_appointments = format_appointments(barbers_appointments)
    appointments = get_this_weeks_appointments(barbers_formated_appointments)

    from pprint import pprint
    pprint(appointments)
    # TODO Do something with the data?
    # Store data in sqlite database?


if __name__ == '__main__':
    main()
