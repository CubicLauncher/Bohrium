# Bohrium

Esta API permite obtener la URL de la skin de un jugador de Minecraft premium usando su nombre de usuario.

## Requisitos

- Go 1.21 o superior

## Instalación

1. Clona este repositorio
2. Instala las dependencias:
```bash
go mod download
```

## Ejecución

Para iniciar el servidor:
```bash
go run main.go
```

El servidor se iniciará en `http://localhost:8080`

## Uso de la API

### Obtener la skin de un jugador

**Endpoint:** `GET /skin/:username`

**Ejemplo:**
```bash
curl http://localhost:8080/skin/Notch
```

**Respuesta:**
```json
{
    "username": "Notch",
    "uuid": "069a79f444e94726a5befca90e38aaf5",
    "skin_url": "https://textures.minecraft.net/texture/..."
}
```

## Notas

- La API solo funciona con cuentas premium de Minecraft
- Si el jugador no existe o no tiene una skin, la API devolverá un error
- La API incluye CORS habilitado para permitir peticiones desde cualquier origen 
