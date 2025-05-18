# Bohrium

API REST en Go que permite obtener las skins de jugadores premium de Minecraft usando su nombre de usuario.

## Características

- Obtener la skin completa de un jugador
- Obtener solo la cabeza del jugador
- Soporte para CORS
- Estructura modular y mantenible
- Manejo de errores robusto

## Estructura del Proyecto

```
minecraft-skin-api/
├── handlers/           # Manejadores HTTP
│   └── skin_handler.go # Maneja las peticiones de skins
├── models/            # Modelos de datos
│   └── models.go      # Estructuras de datos
├── services/          # Lógica de negocio
│   └── minecraft_service.go # Servicio para interactuar con la API de Minecraft
└── main.go           # Punto de entrada de la aplicación
```

## Requisitos

- Go 1.21 o superior

## Instalación

1. Clona este repositorio:
```bash
git clone <url-del-repositorio>
cd minecraft-skin-api
```

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

### Obtener la skin completa de un jugador

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

### Obtener solo la cabeza del jugador

**Endpoint:** `GET /skin/:username?head=true`

**Ejemplo:**
```bash
curl http://localhost:8080/skin/Notch?head=true
```

**Respuesta:**
```json
{
    "username": "Notch",
    "uuid": "069a79f444e94726a5befca90e38aaf5",
    "skin_url": "https://crafatar.com/avatars/069a79f444e94726a5befca90e38aaf5"
}
```

## Códigos de Error

- `400 Bad Request`: Cuando el nombre de usuario no está presente en la URL
- `404 Not Found`: Cuando el jugador no existe
- `500 Internal Server Error`: Cuando hay un error al procesar la petición

## Notas

- La API solo funciona con cuentas premium de Minecraft
- Si el jugador no existe o no tiene una skin, la API devolverá un error
- La API incluye CORS habilitado para permitir peticiones desde cualquier origen
- Las skins se obtienen directamente de los servidores oficiales de Minecraft
- Las cabezas se obtienen a través del servicio Crafatar

## Contribuir

Las contribuciones son bienvenidas. Por favor, asegúrate de:

1. Hacer fork del repositorio
2. Crear una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir un Pull Request 
