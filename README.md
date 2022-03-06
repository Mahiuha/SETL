# SETL (Safe Energy Transmission Lines) :penguin:

<p align="left">
    <img src="assets/images/logo.png">
</p>

## Registering a Device.

### Usage

```json
{
    "data": "Mixed type holding the content of the response",
    "message": "Description of what happened"
}
```

Subsequent response definitions will only detail the expected value of the `data field`

### List all Devices

**Definition**

`GET /devices`

**Response**

- `200 OK` on success

```json
[
    {
        "identifier": "vibration-sensor",
        "name": "Vibration Sensor",
        "device_type": "Sensor",
        "controller_gateway": "127.0.0.1"
    },
    {
        "identifier": "alarm",
        "name": "Alarm",
        "device_type": "output",
        "controller_gateway": "127.0.0.2"
    },
    {
        "identifier": "gsm",
        "name": "GSM",
        "device_type": "gsm",
        "controller_gateway": "127.0.0.3"
    }
]
```

### Registering a new device

**Definition**

`POST /devices`

**Arguments**

- `"identifier":string` a globally unique identifier for this device.
- `"name":string` a friendly name for the device.
- `"device_type":string` the type of the device as understood by the client.
- `"controller_gateway":string` the ip address of the devices controller.

If a device with the given identifier already exists, the existing device will be overwritten.

- `201 Created` on success

```json
{
        "identifier": "vibration-sensor",
        "name": "Vibration Sensor",
        "device_type": "Sensor",
        "controller_gateway": "127.0.0.1"
}
```

## Lookup device

`GET /device/<identifier>`

**Response**

- `404 Not Found` if the device does not exist
- `200 OK` on success

```json
{
    "identifier": "vibration-sensor",
    "name": "Vibration Sensor",
    "device_type": "Sensor",
    "controller_gateway": "127.0.0.1"
}
```

## Delete a device

**Definition**

`DELETE /device/<identifier>`

**Response**

- `404 Not Found` if the device does not exist
- `204 No Content` on success

### Technologies used. :whale2:


