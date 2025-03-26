# Testes de carga

Este repositório contém os scripts e configurações de teste de carga para o projeto, utilizando a
ferramenta [K6](https://k6.io/). Os testes de carga são essenciais para simular condições reais de uso e garantir que a
aplicação possa suportar o tráfego esperado sem degradação de desempenho.

## Instalação do K6

O K6 é uma ferramenta robusta para testes de carga, que permite simular tráfego em um ambiente de produção. Siga as
instruções abaixo para instalar o K6 no seu sistema:

### Windows

Para instalar o K6 no Windows, você pode baixar o executável diretamente
do [GitHub](https://github.com/grafana/k6/releases) ou usar um gerenciador de pacotes como o Chocolatey:

```bash
choco install k6
```

### macOS

No macOS, o K6 pode ser instalado via Homebrew:

```bash
brew install k6
```

### Linux

Para usuários de Linux, você pode instalar o K6 usando um gerenciador de pacotes apropriado para a sua distribuição. Por
exemplo, para Debian/Ubuntu:

```bash
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61
echo "deb https://dl.bintray.com/loadimpact/deb stable main" | sudo tee -a /etc/apt/sources.list
sudo apt-get update
sudo apt-get install k6
```

# Executando os Testes

Para executar os testes de carga, execute o comando abaixo, substituindo `<nome-do-script-de-teste>` pelo nome do seu
script de teste. Certifique-se de que o K6 está instalado corretamente em sua
máquina.

```bash
k6 run <nome-do-script-de-teste>.js
```