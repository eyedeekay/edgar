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
Usage of edgar:
  -author string
    	The author of the HTML file (default "eyedeekay")
  -css string
    	The CSS file to use, a default will be generated if one doesn't exist (default "style.css")
  -donate string
    	add donation section to cryptocurrency wallets. Use the address URL schemes, separated by commas(no spaces). Change them before running unless you want the money to go to me. (default "monero:4A2BwLabGUiU65C5JRfwXqFTwWPYNSmuZRjbTDjsu9wT6wV6kMFyXn83ydnVjVcR7BCsWh8B5b4Z9b6cmqjfZiFd9sBUpWT,bitcoin:1D1sDmyZAs5q2Lb29q8TBnGhEJK7vfp5PJ,ethereum:0x539a4356bb0566a39376CaC3F50B558F77E84eC9")
  -filename string
    	The markdown file to convert to HTML, or a comma separated list of files (default "README.md,USAGE.md,index.html,docs/README.md")
  -i2plink
    	add an i2p link to the page footer. Logo courtesy of @Shoalsteed and @mark22k (default true)
  -nodonate
    	disable the donate section(change the -donate wallet addresses before setting this to true) (default true)
  -out inputfile.html
    	The name of the output file(Only used for the first file, others will be named inputfile.html) (default "index.html")
  -script string
    	The script file to use.
  -snowflake
    	add a snowflake to the page footer (default true)
  -support string
    	change message/CTA for donations section. (default "Support independent development of edgar")
  -title string
    	The title of the HTML file, if blank it will be generated from the first h1 in the markdown file.
```
