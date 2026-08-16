# Desafio Técnico - Korp

Repositório contendo a solução desenvolvida para o desafio técnico da empresa Korp, estruturado em três frentes: Desenvolvimento e Infraestrutura, Monitoramento e Observabilidade, e Automação de Deploy.

---

## Arquitetura do Projeto

O ecossistema foi construído utilizando containers Docker gerenciados via Docker Compose, divididos nos seguintes serviços:

1. **API (Go):** Um servidor HTTP desenvolvido em Go (porta 8080) que expõe o endpoint `/projeto-korp` (retornando nome e horário em UTC) e o endpoint `/metrics` com as métricas instrumentadas.
2. **NGINX:** Atua como proxy reverso (porta 8088), redirecionando o tráfego externo de forma segura para a nossa API interna.
3. **Prometheus:** Coleta e armazena as métricas de disponibilidade e volume de requisições da API.
4. **Grafana:** Interface de visualização que consome os dados do Prometheus e exibe o dashboard de monitoramento.

---

## Como Executar o Projeto Localmente

Certifique-se de ter o **Docker** e o **Docker Compose** instalados em sua máquina. Na raiz do projeto, execute o comando:

```bash
docker compose up -d --build

Após subir os containers, os seguintes serviços estarão disponíveis:

    API via NGINX: http://localhost:8088/projeto-korp

    Métricas do Prometheus (API): http://localhost:8088/metrics

    Prometheus Dashboard: http://localhost:9090

    Grafana (Painel de Monitoramento): http://localhost:3030 (Usuário/Senha: admin / admin)

Parte 2: Monitoramento e Observabilidade

O serviço está instrumentado com a biblioteca oficial do Prometheus em Go, rastreando:

    http_requests_total (Contador do volume total de requisições para o endpoint).

    service_availability (Métrica de gauge indicando o status de saúde da API: 1 para UP, 0 para DOWN).

Provisionamento Automático (Bônus)
O Grafana está configurado para ser provisionado de forma totalmente automatizada (sem necessidade de configuração manual). Os datasources, provedores e o layout do dashboard estão estruturados dentro da pasta grafana/provisioning/.

Parte 3: Automação com Ansible

Para garantir a reprodutibilidade do ambiente em um servidor Linux remoto de forma automatizada (Infrastructure as Code), utilizamos um playbook do Ansible.

Para executar o deploy completo em sua Máquina Virtual com zero interação manual, siga os passos:

1 -    Configure o IP da sua VM no arquivo hosts.ini.

2 -    No terminal Linux (ou WSL), acesse o diretório do projeto e execute a sequência de provisionamento:



# 1. Acesse o diretório onde o projeto foi clonado/salvo
cd /caminho/para/o/desafio-korp

# 2. Desative a checagem de chave do SSH para automação contínua
export ANSIBLE_HOST_KEY_CHECKING=False

# 3. Execute o provisionamento via Ansible
ansible-playbook -i hosts.ini playbook.yml -k -K



---