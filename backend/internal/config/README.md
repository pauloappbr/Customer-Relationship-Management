# Gestão de Configurações (Viper)

Este pacote centraliza todas as configurações do Backend, servindo como a **Única Fonte de Verdade**.

## Como adicionar uma nova variável?
1. Adicione a chave no arquivo `.env` na raiz do projeto.
2. Adicione a chave no arquivo `.env.example` para documentação de outros devs.
3. Adicione o novo campo na struct `Config` dentro de `config.go`.
4. Use a tag `mapstructure:"NOME_DA_VARIAVEL"`.

O sistema usa **Reflection** para vincular automaticamente as variáveis de ambiente à struct. Não é necessário atualizar a lógica de carregamento.

## Validação (DoD)
A aplicação implementa o padrão **Fail Fast**. Se uma variável marcada como obrigatória no método `LoadConfig` estiver vazia, o container não subirá e um erro JSON será disparado nos logs.