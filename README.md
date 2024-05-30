# Rate Limiter em Go

Este projeto implementa um rate limiter em Go que pode ser configurado para limitar o número máximo de requisições por segundo com base em um endereço IP específico ou em um token de acesso. O rate limiter utiliza Redis para persistência e pode ser configurado via variáveis de ambiente.

## Funcionalidades

- **Limitação por IP**: Restringe o número de requisições recebidas de um único endereço IP dentro de um intervalo de tempo definido.
- **Limitação por Token**: Limita as requisições baseadas em um token de acesso único, permitindo diferentes limites de tempo de expiração para diferentes tokens.
- **Configuração via Variáveis de Ambiente**: Permite configurar o número máximo de requisições permitidas por segundo e o tempo de bloqueio.
- **Persistência com Redis**: Armazena informações de limitação em um banco de dados Redis.
- **Middleware**: Implementado como middleware para fácil integração com servidores web.

## Execução

- docker-compose up -d
