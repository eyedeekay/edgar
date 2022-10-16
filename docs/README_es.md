edgar
=====

genera una página web para cualquier cosa con un readme. Un reemplazo para mis insensatez incontable.
Esta herramienta tiene como objetivo crear páginas para proyectos basados en el archivo README.md, y
Espero que sea particularmente útil para las páginas de github.

Básicamente, un generador de sitio estático muy simple que toma un solo archivo marcado y emite
página HTML razonable para ella.

STATUS: Este proyecto se mantiene. Responderé a cuestiones, pediré y presentaré solicitudes dentro de unos días. Lo hace
lo que se supone que debe hacer.

Usage
---

```md
Usage of edgar:
- cadena de autor
El autor del archivo HTML (default eyedeekay)
-css string
El archivo CSS para usar, se generará un defecto si no existe (estilo predeterminado.css)
-anotar la cuerda
añadir la sección de donación a carteras de criptomoneda. Utilice los esquemas URL de la dirección, separados por comas(no espacios). Cámbialos antes de correr a menos que quieras que me vaya el dinero. (default monero:4A2BwLabGUiU65C5JRfwXqFTwWPYNSmuZRjbTDjsu9wT6wV6kMFyXn83ydnVjVcR7BCsWh8B5b4Z9b6cmqfZiFd9sBUpWT,bitcoin:1D1sD2JE2
- la cadena de nombres de archivo
El archivo marcador para convertir a HTML, o una lista separada por coma de archivos (default README.md,USAGE.md,index.html,docs/README.md)
- i2plink
añadir un enlace i2p a la página de pie. Logo cortesía de @Shoalsteed y @mark22k (default true)
-nodonate
deshabilitar la sección de donar (cambiar las direcciones de la cartera -donar antes de ajustar esto a la verdad) (por defecto verdadero)
- fuera del archivo de entrada. html
El nombre del archivo de salida (sólo utilizado para el primer archivo, otros serán nombrados fichero de entrada.html) (índice predeterminado.html)
- cadena de comandos
El archivo script para usar.
-snowflake
añadir un copo de nieve a la página de pie (por defecto verdadero)
- cuerda de soporte
cambiar el mensaje/CTA para la sección de donaciones. (default "Desarrollo independiente de Edgar")
- titular
El título del archivo HTML, si está en blanco se generará desde la primera h1 en el archivo marcado.
```
