package tohtml

var DefaultCSS = `/* edgar default CSS file */

body {
    font-family: "Roboto";
    font-family: monospace;
    text-align: justify;
    background-color: #373636;
    color: whitesmoke;
    font-size: 1.15em;
}

ul {
    width: 55%;
    display: block;
}

ol {
    width: 55%;
    display: block;
}

li {
    margin-top: 1%;
}

p {
    max-width: 90%;
    margin-top: 1%;
    margin-left: 3%;
    margin-right: 3%;
}

img {
    float: left;
    top: 5%;
    left: 5%;
    max-width: 60%;
    display: inline;
    padding-right: 2%;
}

.inline {
    display: inline;
}

.link-button:focus {
    outline: none;
}

.link-button:active {
    color: red;
}

code {
    font-family: monospace;
    border-radius: 5%;
    padding: 1%;
    border-color: darkgray;
    font-size: .9em;
}

a {
    color: #C6D9FE;
    padding: 1%;
}

ul li {
    color: #C6D9FE;
}

iframe {
    background: aliceblue;
    border-radius: 15%;
    margin: 2%;
}

.container {
    display: inline-block;
    margin: 0;
    padding: 0;
}

.editor-toolbar a {
    display: inline-block;
    text-align: center;
    text-decoration: none !important;
    color: whitesmoke !important;
}

#feed {
    width: 60vw;
    height: unset !important;
    margin: 0;
    padding: 0;
    float: right;
    background-color: #373636;
    color: whitesmoke;
    border: #C6D9FE solid 1px;
}

.thread-post,
.thread {
    color: whitesmoke !important;
    background-color: #373636;
    border: 1px solid darkgray;
    font-size: inherit;
    padding-top: 1%;
    padding-bottom: 1%;
}

.thread-post {
    margin-left: 4%;
}

input {
    text-align: center;
    color: whitesmoke !important;
    background-color: #373636;
    border: 1px solid darkgray;
    font: normal normal normal 14px/1 FontAwesome;
    font-size: inherit;
    padding-top: 1%;
    padding-bottom: 1%;
}

.thread-hash {
    text-align: right;
    color: whitesmoke !important;
    background-color: #373636;
    border: 1px solid darkgray;
    font-size: inherit;
    padding-top: 1%;
    padding-bottom: 1%;
}

.post-body {
    text-align: left;
    color: whitesmoke !important;
    font-size: inherit;
    padding-top: 1%;
    padding-bottom: 1%;
}
#show {display:none; }
#hide {display:block; }
#show:target {display: block; }
#hide:target {display: none; }

#shownav {display:none; }
#hidenav {display:block; }
#shownav:target {display: block; }
#hidenav:target {display: none; }

#navbar {
	float: right;
	width: 15%;
}
#returnhome {
    font-size: xxx-large;
    display: inline;
}
h1 {
    display: inline;
}
`

var ShowHiderCSS = `/* edgar showhider CSS file */
#show {display:none; }
#hide {display:block; }
#show:target {display: block; }
#hide:target {display: none; }

#shownav {display:none; }
#hidenav {display:block; }
#shownav:target {display: block; }
#hidenav:target {display: none; }

#donate {display:none; }
#hidedonate {display:block; }
#donate:target {display: block; }
#hidedonate:target {display: none; }
`

var DarkLightCSS = `/* edgar darklight CSS file */
#checkboxDarkLight:checked + .container {
    background-color: #202020;
    filter: invert(100%);
}
#checkboxDarkLight{
    appearance: none;
    width: 80px;
    height: 40px;
    background: black;
    border-radius: 22px;
    cursor: pointer;
    outline: none;
}
#checkboxDarkLight::before{
    content: '';
    width: 40px;
    height: 35px;
    background-color:white;
    border-radius: 35px;
    cursor: pointer;
    transition: .3s linear;
}

`
