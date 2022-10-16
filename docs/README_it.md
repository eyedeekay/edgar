edilizia
Traduzione:

genera una homepage per qualsiasi cosa con un readme. Un sostituto per le mie infinite sciocchezze.
Questo strumento è destinato a creare pagine per progetti che si basano sul file README.md, e
Spero che sia particolarmente utile per le pagine github.

Fondamentalmente, un generatore di sito statico molto semplice che prende un singolo file di markdown ed emette
pagina HTML ragionevole per esso.

STATUS: Questo progetto è mantenuto. Risponderò a problemi, estrarre richieste e funzionalità entro pochi giorni. Lo fa
cosa dovrebbe fare.

Utilizzo
-----

```md
Utilizzo di edgar:
- stringa di autori
L'autore del file HTML (default eyedeekay)
-css string
Il file CSS da usare, un default verrà generato se non esiste (default style.css)
- corde donate
aggiungere sezione donazione ai portafogli di criptovaluta. Utilizzare gli schemi URL dell'indirizzo, separati da virgole (senza spazi.) Cambiali prima di correre a meno che tu non voglia che i soldi vengano da me. (DEFOP)
-filename string
Il file markdown da convertire in HTML, o un elenco separato da virgola di file (default README.md,USAGE.md,index.html,docs/README.md)
- i2plink
aggiungere un link i2p al piè di pagina. Logo per gentile concessione di @Shoalsteed e @mark22k (default true)
- nodonate
disabilitare la sezione donare (cambiare gli indirizzi del portafoglio -donate prima di impostare questo a true) (default true)
- senza inputfile. html
Il nome del file di output (Solo usato per il primo file, altri saranno nominati inputfile.html) (default index.html)
- stringa del testo
Il file di script da usare.
- fiocco di neve
aggiungere un fiocco di neve alla pagina footer (default true)
- stringa di supporto
cambiare messaggio/CTA per la sezione donazioni. (default "Supporto lo sviluppo indipendente di edgar")
-title string
Il titolo del file HTML, se vuoto verrà generato dal primo h1 nel file markdown.
```
