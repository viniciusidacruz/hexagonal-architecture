# Arquitetura Hexagonal com Go

Este projeto demonstra a implementação da **Arquitetura Hexagonal** (também conhecida como Ports and Adapters) em Go, utilizando as melhores práticas de desenvolvimento e tecnologias modernas.

## 🏗️ Arquitetura

A aplicação segue o padrão de **Arquitetura Hexagonal**, que separa a lógica de negócio (domínio) das tecnologias externas através de interfaces bem definidas. Esta abordagem garante:

- **Independência de frameworks**: A lógica de negócio não depende de tecnologias específicas
- **Testabilidade**: Facilita a criação de testes unitários e de integração
- **Flexibilidade**: Permite trocar implementações sem afetar o core da aplicação
- **Manutenibilidade**: Código organizado e fácil de manter

### Estrutura do Projeto

```
├── application/          # Camada de aplicação (casos de uso)
│   ├── product.go       # Entidade e interfaces do domínio
│   ├── product_service.go # Serviços de aplicação
│   └── mocks/           # Mocks para testes
├── adapters/            # Adaptadores (Ports and Adapters)
│   ├── web/             # Adaptadores para web (HTTP)
│   │   ├── handler/     # Handlers HTTP
│   │   ├── dto/         # Data Transfer Objects
│   │   └── server/      # Servidor HTTP
│   ├── cli/             # Adaptadores para CLI
│   └── db/              # Adaptadores para banco de dados
├── infra/               # Infraestrutura
│   └── prisma/          # Configuração do Prisma
├── cmd/                 # Pontos de entrada da aplicação
└── prisma/              # Schema do banco de dados
```

## 🛠️ Tecnologias Utilizadas

### Core Framework

- **Go 1.24.3**: Linguagem principal, escolhida por sua performance, simplicidade e excelente suporte para concorrência

### Web Framework e Middleware

- **Gorilla Mux**: Router HTTP robusto e flexível para criação de APIs RESTful
- **Negroni**: Middleware elegante que fornece logging, recovery e funcionalidades de segurança
- **Por que usar**: Gorilla Mux oferece roteamento avançado com suporte a variáveis de URL, enquanto Negroni simplifica a adição de middleware essenciais

### Banco de Dados e ORM

- **PostgreSQL**: Banco de dados relacional robusto e confiável
- **Prisma**: ORM moderno com type-safety e excelente DX (Developer Experience)
- **Por que usar**: Prisma oferece geração automática de código Go type-safe, migrations automáticas e uma API intuitiva

### Validação e Utilitários

- **govalidator**: Biblioteca para validação de structs com tags
- **go.uuid**: Geração de UUIDs seguros e únicos
- **shopspring/decimal**: Manipulação precisa de valores decimais para operações financeiras
- **Por que usar**: Estas bibliotecas garantem integridade de dados e precisão em operações críticas

### CLI e Configuração

- **Cobra**: Framework para criação de CLIs poderosas e intuitivas
- **godotenv**: Carregamento de variáveis de ambiente a partir de arquivos .env
- **Por que usar**: Cobra facilita a criação de comandos CLI complexos, enquanto godotenv mantém configurações seguras

### Testes

- **testify**: Framework de testes com assertions e mocks
- **golang/mock**: Geração automática de mocks para interfaces
- **Por que usar**: Facilita a criação de testes unitários e de integração robustos

## 🚀 Como Executar

### Pré-requisitos

- Go 1.24.3 ou superior
- PostgreSQL
- Prisma CLI

### Configuração

1. Clone o repositório
2. Configure as variáveis de ambiente:

   ```bash
   # Copie o arquivo de exemplo
   cp env.example .env

   # Edite o arquivo .env com suas configurações
   nano .env
   ```

3. Execute as migrations do Prisma:
   ```bash
   prisma generate
   prisma db push
   ```

### Executando a Aplicação

#### Servidor HTTP

```bash
go run main.go http
```

A API estará disponível em `http://localhost:8080`

#### CLI

```bash
# Criar produto
go run main.go cli -a create -n "Produto Teste" -p 99.99

# Buscar produto
go run main.go cli -a get -i "uuid-do-produto"

# Habilitar produto
go run main.go cli -a enable -i "uuid-do-produto"

# Desabilitar produto
go run main.go cli -a disable -i "uuid-do-produto"
```

## 📡 API Endpoints

### Produtos

- `GET /product/{id}` - Buscar produto por ID
- `POST /product` - Criar novo produto
- `GET /product/{id}/enable` - Habilitar produto
- `GET /product/{id}/disable` - Desabilitar produto

### Exemplo de Criação de Produto

```bash
curl -X POST http://localhost:8080/product \
  -H "Content-Type: application/json" \
  -d '{"name": "Produto Teste", "price": 99.99}'
```

## 🧪 Testes

Execute os testes com:

```bash
go test ./...
```

A aplicação inclui testes unitários para:

- Entidades de domínio
- Serviços de aplicação
- Handlers HTTP
- Adaptadores CLI

## 🏛️ Princípios da Arquitetura Hexagonal

### 1. **Ports (Interfaces)**

- `ProductInterface`: Define o contrato da entidade Produto
- `ProductServiceInterface`: Define os casos de uso
- `ProductPersistenceInterface`: Define o contrato de persistência

### 2. **Adapters (Implementações)**

- **Primary Adapters**: Web handlers e CLI commands
- **Secondary Adapters**: Implementação do banco de dados com Prisma

### 3. **Application Core**

- Lógica de negócio isolada em `application/`
- Regras de validação e comportamento centralizadas
- Independência total de frameworks externos

## 🎯 Benefícios da Arquitetura

1. **Testabilidade**: Fácil mock de dependências externas
2. **Flexibilidade**: Troca de tecnologias sem afetar o core
3. **Manutenibilidade**: Código organizado e responsabilidades claras
4. **Escalabilidade**: Fácil adição de novos adaptadores
5. **Independência**: Core de negócio isolado de detalhes técnicos

## 📚 Conceitos Implementados

- **Domain-Driven Design (DDD)**: Entidades e regras de negócio bem definidas
- **Dependency Inversion**: Dependências através de interfaces
- **Single Responsibility**: Cada componente tem uma responsabilidade específica
- **Interface Segregation**: Interfaces pequenas e coesas
- **Clean Architecture**: Separação clara de camadas

## 🔧 Desenvolvimento

### Adicionando Novos Recursos

1. Defina a entidade no domínio (`application/`)
2. Crie os casos de uso necessários
3. Implemente os adaptadores (web, cli, db)
4. Adicione testes para cada camada

### Estrutura de Commits

- `feat`: Nova funcionalidade
- `fix`: Correção de bug
- `refactor`: Refatoração de código
- `test`: Adição ou correção de testes
- `docs`: Documentação

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

**Desenvolvido com ❤️ usando Go e Arquitetura Hexagonal**
