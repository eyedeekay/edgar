edgar
=====

generates a homepage for anything with a readme. A replacement for my endless makefile nonsense.
This tool is intended to create pages for projects which are based on the README.md file, and
I hope it's particularly useful for github pages.

Basically, a really simple static site generator that take a directory full of markdown files and emits
reasonable-looking HTML pages for them.

```sh
rm -f edgar
wget https://github.com/eyedeekay/edgar/releases/download/v0.34.2/edgar
sudo cp edgar /usr/bin/edgar
sudo chmod +x /usr/bin/edgar
```

STATUS: This project is maintained. I will respond to issues, pull requests, and feature requests within a few days. It does
what it's supposed to do.