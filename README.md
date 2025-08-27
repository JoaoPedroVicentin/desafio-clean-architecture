
# Desafio GoExpert - Clean Architecture

O deafio inclui implementar o useCase para listar os pedidos e realizar a consulta com:
- Endpoint REST (GET /orders)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

## 📦 Como rodar o projeto

### 1. Clone o repositório

```bash
git clone https://github.com/JoaoPedroVicentin/desafio-clean-architecture.git
cd desafio-clean-architecture
```

### 2. Instale as dependências

```bash
go mod tidy
```

### 3. Suba os containers

```bash
docker compose up -d
```

### 4. Entre no container do mySql

```bash
docker exec -it <container-id> bash
```

### 5. Acesse e digite a senha: root

```bash
mysql -uroot -p orders
```

### 6. Gere a migration

```bash
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));
```

### 7. Inicie o programa

```bash
cd cmd/ordersystem
go run main.go wire_gen.go
```

### 8. Caso esteja tudo funcionando, a mensagem no terminal aparecerá

```bash
Starting web server on port :8000
Starting gRPC server on port 50051
Starting GraphQL server on port 8080
```

---

<div align="center">
<h3>👨‍💻</h3>
    <h3> Criado por João Pedro Vicentin!</h3>
    <div>
        <h3>
            <a href="https://www.linkedin.com/in/joaopedrovicentin/" target="_blank">Linkedin</a>
            <a href='https://github.com/JoaoPedroVicentin' target='_blank'>Github</a>
            <a href="https://contate.me/joao-pedro-lopes-vicentin" target="_blank">Whatsapp</a>
        </h3>
    </div>
</div>
