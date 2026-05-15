# Health Checker - Monitoramento de Disponibilidade (Go & Next.js)

Uma aplicação fullstack desenvolvida para realizar monitoramento ativo de disponibilidade de aplicações e websites, fornecendo análises de status HTTP em tempo real. O backend foi construído visando alta performance e execução paralela de requisições, enquanto o frontend apresenta as informações de forma reativa e intuitiva.

## 🚀 Arquitetura e Decisões Técnicas

### Backend (Go)
O backend deste projeto foi inteiramente escrito em **Go**, adotando os princípios de bibliotecas padrão (Standard Library) sempre que possível para garantir uma aplicação enxuta e estável.

* **Concorrência Otimizada:** Para evitar gargalos durante o monitoramento de múltiplos sites, a verificação de status HTTP é realizada assincronamente através de `Goroutines`.
* **Sincronização de Estado:** Implementação de `sync.WaitGroup` para coordenar as rotinas e uso de `sync.Mutex` para prevenir *race conditions* durante a escrita de logs e tratamento de erros paralelos.
* **Persistência de Dados (SQLite):** Armazenamento estruturado utilizando o driver `go-sqlite3`, permitindo que o estado dos sites seja mantido mesmo após o reinício do servidor.
* **Middlewares e Roteamento:** Tratamento customizado de rotas e segurança (suporte a *CORS*) construídos sobre o multiplexador nativo do Go (`http.NewServeMux`).

### Frontend (Next.js)
O frontend adota uma abordagem moderna utilizando React e o App Router do Next.js.

* **Reatividade com SWR:** A listagem dos sites é consumida via `SWR` (stale-while-revalidate), configurada com *polling* automático a cada 30 segundos, garantindo que o dashboard reflita o estado real dos serviços sem recarregar a página.
* **Componentização e Tipagem:** Desenvolvimento modular com componentes funcionais e **TypeScript**, garantindo segurança de tipos em toda a aplicação.
* **UI com Tailwind CSS:** Interface responsiva e otimizada, utilizando feedbacks visuais de cores baseados no status code retornado pela API.

## 💻 Tecnologias Utilizadas

**Backend**
* [Go 1.24+](https://go.dev/)
* SQLite3 + [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
* API RESTful

**Frontend**
* [Next.js](https://nextjs.org/) (App Router)
* React & TypeScript
* [SWR](https://swr.vercel.app/) (Data Fetching)
* Tailwind CSS

## 🔧 Como Executar Localmente

### Pré-requisitos
* Go instalado (1.24 ou superior)
* GCC (Necessário para o CGO do SQLite)
* Node.js (v18+) e gerenciador de pacotes (npm/yarn)

### Passos para execução

**1. Clone o repositório** e acesse a pasta raiz.

**2. Inicie o Backend:**
```bash
go mod tidy
go run main.go
````
*O servidor inciará em `http://localhost:8080`.*

**3. Inicie o Frontend:**
```bash
cd web
npm install
npm run dev
```

*O dashboard estará disponível em `http://localhost:3000`.*

## 📖 Endpoints da API

* `GET /` - Lista todos os sites e executa as verificações de status concorrentes.

* `POST /adicionar` - Cadastra uma nova URL (JSON: `nome`, `url`).

* `PUT /atualizar?id={id}` - Atualiza dados de um site existente.

* `DELETE /remover?id={id}` - Remove um site do monitoramento.

---
Desenvolvido por [Juan Gonçalves Nascimento](https://github.com/ojuangoncalves)
