edgar
=====

generates a homepage for anything with a readme. A replacement for my endless makefile nonsense.
This tool is intended to create pages for projects which are based on the README.md file, and
I hope it's particularly useful for github pages.

Basically, a really simple static site generator that take a single markdown file and emits
reasonable-looking HTML page for it.

STATUS: This project is maintained. I will respond to issues, pull requests, and feature requests within a few days. It does
what it's supposed to do.

Usage
-----

```md
# Usage of ./edgar:
  -author string
    	The author of the HTML file (default is your git namespace)
  -css string
    	The CSS file to use, a default will be generated if one doesn't exist (default "style.css")
  -filename string
    	The markdown file to convert to HTML (default "README.md")
  -out string
    	The name of the output file (default "index.html")
  -script string
    	The script file to use (default "script.js")
  -snowflake
    	add a snowflake to the page footer (default true)
  -title string
    	The title of the HTML file
```
