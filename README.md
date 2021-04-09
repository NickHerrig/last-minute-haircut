# Last Minute Haircut 
> Last minute haircut notifys me when my barber has appointments available within the week. 

![](static/paramount.jpeg)

## Installation

OS X & Linux:

```sh
git clone https://github.com/nickherrig/last-minute-haircut
make run
```

## API

```sh
make serve
```

### Get all openenings
```sh
> curl localhost:8000
{'jordan': [], 'pete': [], 'brandon': [], 'luis': [], 'zach': [], 'paul': [], 'kegan': []}
```

### Get single barber openenings
```sh
> curl localhost:8000/jordan
{"jordan": []}
```

## Docker

```
---
version: "3"
services:
  last-minute-haircut:
    image: ghcr.io/NickHerrig/last-minute-haircut:0.1.0
    container_name: last-minute-haircut
    ports:
      - 8000:8000
    restart: unless-stopped
```