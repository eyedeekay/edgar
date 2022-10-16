essig
======

erzeugt eine Homepage für alles mit einem Lesegerät. Ein Ersatz für meinen endlosen Makefile Unsinn.
Dieses Tool soll Seiten für Projekte erstellen, die auf der README.md-Datei basieren, und
Ich hoffe, es ist besonders nützlich für github-Seiten.

Grundsätzlich ein wirklich einfacher statischer Site-Generator, der eine einzelne Markdown-Datei nimmt und emittiert
vernünftig aussehende HTML-Seite dafür.

STATUS: Dieses Projekt wird beibehalten. Ich werde innerhalb von wenigen Tagen auf Probleme reagieren, Anfragen ausziehen und Anfragen ausführen. Es tut
was es tun soll.

Verwendung
----

```md
Verwendung von Essig:
- autor string
Der Autor der HTML-Datei (default eyedeekay)
-css string
Die zu verwendende CSS-Datei, ein Standard wird generiert, wenn man nicht existiert (default style.css)
-donate string
fügen Sie Spendenabschnitt zu Kryptowährung Geldbörsen. Verwenden Sie die Adress-URL-Systeme, getrennt durch commas(keine Leerzeichen). Ändern Sie sie vor dem Laufen, es sei denn, Sie wollen das Geld zu mir gehen. (Standard-Monero:4A2BwLabGUiU65C5JRfwXqFTwWPYNSmuZRjbTDjsu9wT6wV6kMFyXn83ydnVjVcR7BCsWh8B5b4Z9b6cmqjfZiFd9sBD9sBD5
- dateiname string
Die Markdown-Datei, um in HTML zu konvertieren, oder eine Komma getrennte Liste von Dateien (Standard README.md,USAGE.md,index.html,docs/README.md)
-i2plink
fügen Sie einen i2p-Link zum Seitenfuß hinzu. Logo mit freundlicher Genehmigung von @Shoalsteed und @mark22k (Standard wahr)
-nodonat
deaktivieren sie den spendenabschnitt (ändern sie die -donate wallet adressen, bevor sie diese auf true setzen) (standard wahr)
- aus eingabedatei. html
Der Name der Ausgabedatei (nur für die erste Datei verwendet, andere werden als inputfile.html) (Standardindex.html) bezeichnet
- schriftzeichenfolge
Die Skriptdatei zu verwenden.
-snowflake
fügen sie eine schneeflocke zum seitenfuss (standard true) hinzu
- unterstützungsstring
nachrichten/CTA für Spenden Sektion ändern. (Standard "Unterstützung unabhängige Entwicklung von Essig")
-titel string
Der Titel der HTML-Datei, wenn leer es aus dem ersten h1 in der Markdown-Datei generiert wird.
```
