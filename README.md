# Weather API

## Tabla de contenidos
[Descripción](#Descripción)  
[Escenario](#Escenario)  
[Endpoints - Ubicaciones](#GET---GetLocation)  
[Endpoints - Clima](#GET---GetWeather)

# 
## Descripción
Esta es una API REST desarrollada en golang para obtener información sobre el clima de ubicaciones específicas.
## Escenario
Para lograr el cometido, la API persiste en formato *json* una lista de ubicaciones establecidas por el usuario de la misma.
Luego se puede consultar en simultáneo la información climática para toda la lista.

## Endpoints
#### GET - Ping()
>localhost:8080/ping

**Response code:** 200
**Response body:**

    {
        "mensaje": "pong!"
    }

### GET - IP()
>localhost:8080/ip

|  | OK | Error interno |
|--|--|--|
| Response code | 200 | 500 |
| Response body| 1*| 2*|

**Response body: (1)**

    {
        "mensaje" : "179.0.238.1"
    }
}

**Response body: (2)**

    {	
	    "mensaje": "Error al obtener su IP"
	}

### GET - GetLocation()
>localhost:8080/location

|  | OK | Error interno |
|--|--|--|
| Response code | 200 | 500 |
| Response body| 1*| 2*|

**Response body: (1)**

    {
        "Status": "success",
        "Data": {
            "subdivision_1_name": "Cordoba",
            "city_name": "Almafuerte",
            "latitude": "-32.15560",
            "longitude": "-64.23890"
    }
}

**Response body: (2)**

    {	
	    "mensaje": "Error al obtener su ubicación"
	}

### GET - GetLocations()
>localhost:8080/locations

|  | OK | No encontrado | Error interno |
|--|--|--|--|
| Response code | 200 | 404 | 500 |
| Response body| 1* | 2* | 3* |

**Response body: (1)**

    [
        {
            "id": "236294580",
            "name": "Río Tercero, Pedanía Yucat, Departamento General San Martín, Córdoba, X5900, Argentina",
            "lat": "-32.3760926",
            "lon": "-63.450102"
        },
        ...
    }

**Response body: (2)**

    {	
	    "mensaje": "No hay ubicaciones cargadas"
	}

**Response body: (3)**

    {	
	    "mensaje": "Error al obtener los datos de la base de datos"
	}

### GET - GetLocationID()
>localhost:8080/locations/{id}

|  | OK | No encontrado | Error interno |
|--|--|--|--|
| Response code | 200 | 404 | 500 |
| Response body| 1* | 2* | 3* |

**Response body: (1)**

    {
        "id": "236689588",
        "name": "Almafuerte, Municipio de Almafuerte, Pedanía Salto, Departamento Tercero Arriba, Córdoba, X5854, Argentina",
        "lat": "-32.1917687",
        "lon": "-64.255066"
    }

**Response body: (2)**

    {	
	    "mensaje": "Ubicación no encontrada"
	}

**Response body: (3)**

    {	
	    "mensaje": "Error al obtener los datos de la base de datos"
	}

### POST - PostLocation()
>localhost:8080/location

**Request body:**

    {
        "City": "Rio Tercero",
        "State": "Cordoba",
        "Country": "Argentina"
    }

|  | OK | Ubicación existente  | Datos inválidos|
|--|--|--|--|
| Response code | 201 | 400 | 400 |
| Response body| 1* | 2* | 3* |

**Response body: (1)**

    {
	    "message": "Ubicación agregada con éxito"
	}

**Response body: (2)**

    {	
	    "message": "La ubicación ya se encuentra registrada"
	}
**Response body: (3)**

    {
	    "message": "JSON inválido"
	}

### DELETE- DeleteLocation()
>localhost:8080/location/{id}

|  | OK | No encontrado | Error interno |
|--|--|--|--|
| Response code | 204 | 404 | 500 |
| Response body| | 1*| 2* |

**Response body: (1)**

    {
	    "message": "Ubicación no encontrada"
	}

**Response body: (2)**

    {
	    "Error al borrar la ubicación"
	}

### PUT- UpdateLocation()
>localhost:8080/location

**Request body:**

    {
        "City": "Rio Tercero",
        "State": "Cordoba",
        "Country": "Argentina"
    }

|  | OK | No encontrado | Error interno |
|--|--|--|--|
| Response code | 200 | 404 | 500 |
| Response body| 1* | 2*| 3* |

**Response body: (1)**

    {
	    "message": "Ubicación actualizada con éxito"
	}

**Response body: (2)**

    {
	    "Ubicación no encontrada"
	}

**Response body: (3)**

    {
	    "Error al actualizar la ubicación"
	}

### GET - GetWeather()
>localhost:8080/weather

|  | OK | No encontrado | Error interno |
|--|--|--|--|
| Response code | 200 | 500 | 500 |
| Response body| 1* | 2* | 3* |

**Response body: (1)**

    {
        "weather": [
            {
                "description": "scattered clouds"
            }
        ],
        "main": {
            "temp": 29.28,
            "feels_like": 31.15,
            "humidity": 58
        },
        "name": "Almafuerte"
    }

**Response body: (2)**

    {	
	    "mensaje": "Error al obtener su ubicación"
	}

**Response body: (3)**

    {	
	    "mensaje": "Error al obtener el clima"
	}

### GET - GetWeatherID()
>localhost:8080/weather/{id}

|  | OK | No encontrado | Error interno |
|--|--|--|--|
| Response code | 200 | 404 | 500 |
| Response body| 1* | 2* | 3* |

**Response body: (1)**

    {
        "weather": [
            {
                "description": "scattered clouds"
            }
        ],
        "main": {
            "temp": 29.28,
            "feels_like": 31.15,
            "humidity": 58
        },
        "name": "Almafuerte"
    }

**Response body: (2)**

    {
        "mensaje": "Ubicación no encontrada"
    }

**Response body: (3)**

    {	
	    "mensaje": "Error al obtener el clima"
	}

### GET - GetAllWeathers()
>localhost:8080/allWeathers/

|  | OK | No encontrado | Error interno |
|--|--|--|--|
| Response code | 200 | 404 | 500 |
| Response body| 1* | 2* | 3* |

**Response body: (1)**

    [
        {
            "weather": [
                {
                    "description": "scattered clouds"
                }
            ],
            "main": {
                "temp": 29.28,
                "feels_like": 31.15,
                "humidity": 58
            },
            "name": "Almafuerte"
        },
        ...
    ]

**Response body: (2)**

    {	
	    "mensaje": "No hay ubicaciones cargadas"
	}

**Response body: (3)**

    {	
	    "mensaje": "Error al obtener el clima"
	}
