# Bienvenido a Jarvis - Tu Asistente Personal en Go

¡Bienvenido a Jarvis, tu compañero digital diseñado para hacer tu vida más fácil y eficiente! Esta aplicación de Go está diseñada para ser tu asistente personal, brindándote una amplia gama de funciones para simplificar tu día a día.

## ¿Qué es Jarvis?

Jarvis es una aplicación de asistente personal que utiliza el lenguaje de programación Go para ofrecer un rendimiento rápido y eficiente. Su objetivo principal es ayudarte en diversas tareas diarias, desde recordatorios y gestionar tu agenda hasta proporcionar información útil de manera instantánea.

## Características Destacadas

- **Recordatorios y Tareas:** Configura recordatorios y gestiona tus tareas pendientes de manera sencilla.
- **Búsqueda Rápida:** Obten información instantánea sobre cualquier tema con solo unos pocos comandos.
- **Agenda Personal:** Organiza tu día, semana o mes con la funcionalidad de calendario integrada.
- **Automatización:** Crea secuencias de comandos personalizadas para automatizar tareas repetitivas.
- **Interfaz Intuitiva:** Una interfaz de usuario amigable y fácil de usar para una experiencia sin complicaciones.

## Requisitos del Sistema

- Go 1.20 o superior

## Como probarlo

#### curl
   ```bash
  curl -X POST localhost:8080/reply -d '{"message":"hola"}''
   ```
#### Respuesta
   ```bash
  {"response":"¡Hola! ¿Cómo puedo ayudarte hoy?"}
  
  
   ```

### Puntos vistos

1 - Integración API Open AI definiendo un contexto particular
2 - Uso de property tools
3 - Parserdor a consultas sql a partir de un esquema definido


### Fuentes 
https://cookbook.openai.com/examples/how_to_call_functions_with_chat_models
https://criptoya.com/api

