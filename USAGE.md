Edgar(Everything does get a README): Static Site Generator for the Paradoxically Prolific
=========================================================================================

This is a static site generator which is intended to generate a page for a piece of software from markdown documents stored inside it's VCS.
It will generate a page from any directory containing markdown files, but it's especially useful for Github Pages with the `.nojekyll` option.

```
Usage of ./edgar:
  -author string
    	The author of the HTML file (default "eyedeekay")
  -checkin
    	commit page generation changes after running (default true)
  -css string
    	The CSS file to use, a default will be generated if one doesn't exist (default "style.css")
  -donate string
    	add donation section to cryptocurrency wallets. Use the address URL schemes, separated by commas(no spaces). Change them before running unless you want the money to go to me. (default "monero:4A2BwLabGUiU65C5JRfwXqFTwWPYNSmuZRjbTDjsu9wT6wV6kMFyXn83ydnVjVcR7BCsWh8B5b4Z9b6cmqjfZiFd9sBUpWT,bitcoin:1D1sDmyZAs5q2Lb29q8TBnGhEJK7vfp5PJ,ethereum:0x539a4356bb0566a39376CaC3F50B558F77E84eC9")
  -filename string
    	The markdown file to convert to HTML, or a comma separated list of files (default "README.md,USAGE.md,docs/index.html,index.html,docs/README.md,docs/README_ar.md,docs/README_de.md,docs/README_es.md,docs/README_fr.md,docs/README_it.md,docs/README_ja.md,docs/README_pt.md,docs/README_ru.md,docs/README_zh.md,docs/index.html,docs/index_ar.html,docs/index_de.html,docs/index_es.html,docs/index_fr.html,docs/index_it.html,docs/index_ja.html,docs/index_pt.html,docs/index_ru.html,docs/index_zh.html")
  -help
    	Show usage.
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

