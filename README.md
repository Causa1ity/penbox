# PenBox

## Usage

Run penbox on the address:
```shell
$ ./penbox.exe ":20080"
```

The function can then be executed by initiating a request.

e.g.:
```shell
$ curl -d "action=forge&secret=app&data={\"user\":\"admin\"}" localhost:20080/payloads/flask/session
{"secret":"app","session":"eyJ1c2VyIjoiYWRtaW4ifQ.YUrFdg.RArBxGNrB-3812rGAoChxfygj-4","data":{"user":"admin"}}
$ curl -d "action=parse&secret=app&session=eyJ1c2VyIjoiYWRtaW4ifQ.YUrFdg.RArBxGNrB-3812rGAoChxfygj-4" localhost:20080/payloads/flask/session
{"secret":"app","session":"eyJ1c2VyIjoiYWRtaW4ifQ.YUrFdg.RArBxGNrB-3812rGAoChxfygj-4","data":{"user":"admin"}}
```

## Functions

* Flask
  * Parse and Forge Session: /payloads/flask/session