edgar
- Sim

gera uma página inicial para qualquer coisa com um readme. Um substituto para o meu infindável bobagem.
Esta ferramenta destina-se a criar páginas para projetos que são baseados no arquivo README.md, e
Espero que seja particularmente útil para páginas github.

Basicamente, um gerador de site estático muito simples que leva um único arquivo de marcação e emite
página HTML de aparência razoável para ele.

STATUS: Este projeto é mantido. Vou responder a problemas, puxar pedidos e solicitações de recursos dentro de alguns dias. Faz
o que é suposto fazer.

Uso
- Não

```md
Uso de edgar:
- string de autor
O autor do arquivo HTML (default eyedeekay)
-css string
O arquivo CSS para usar, um padrão será gerado se não existir (default style.css)
-corda de doação
adicionar seção de doação para carteiras criptomoedas. Use os esquemas URL do endereço, separados por vírgulas (sem espaços). Mude-os antes de correr, a menos que queiras que o dinheiro vá até mim. (em inglês)
-filename string
O arquivo de marcação para converter para HTML, ou uma lista separada por vírgula de arquivos (default README.md,USAGE.md,index.html,docs/README.md)
- i2plink
adicione um link i2p ao rodapé da página. Logo cortesia de @Shoalsteed e @mark22k (correio padrão)
-nodona
desabilitar a seção do doar (mudar os endereços da carteira do doar antes de definir isso para true) (padrão verdadeiro)
- fora do arquivo de entrada. html
O nome do arquivo de saída (Somente usado para o primeiro arquivo, outros serão nomeados inputfile.html) (default index.html)
- string de texto
O arquivo de script a usar.
- agora flake
adicionar um floco de neve ao rodapé da página (padrão verdadeiro)
- string de suporte
mudar mensagem / CTA para seção de doações. (padrão "Desenvolvimento independente do edgar")
- corda de amarração
O título do arquivo HTML, se em branco ele será gerado a partir do primeiro h1 no arquivo markdown.
```
