# Auth Service - Exemplos de Uso

## Executar o servidor

```bash
cd /home/william/Projetos/cortaai-ms-auth-go
go run ./cmd/main.go
```

## Endpoints

### 1. Health Check
```bash
curl -X GET http://localhost:8001/hello
```

Resposta:
```json
{
  "message": "Hello World!"
}
```

### 2. Register (NOVO - IMPLEMENTADO ✅)

**Endpoint:**
```
POST /auth/register
```

**Request:**
```bash
curl -X POST http://localhost:8001/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "novousuario@example.com",
    "password": "senha123456",
    "username": "novousuario"
  }'
```

**Resposta de Sucesso (201):**
```json
{
  "message": "Usuário criado com sucesso",
  "user": {
    "id": 1,
    "email": "novousuario@example.com",
    "username": "novousuario"
  }
}
```

**Resposta de Erro - Email Duplicado (400):**
```json
{
  "error": "registration_failed",
  "description": "email já cadastrado"
}
```

**Resposta de Erro - Password Curta (400):**
```json
{
  "error": "registration_failed",
  "description": "password deve ter pelo menos 6 caracteres"
}
```

**Resposta de Erro - Requisição Inválida (400):**
```json
{
  "error": "invalid_request",
  "description": "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"
}
```

### 3. Login (IMPLEMENTADO ✅ - Agora com banco de dados)

**Endpoint:**
```
POST /auth/login
```

**Request:**
```bash
curl -X POST http://localhost:8001/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "novousuario@example.com",
    "password": "senha123456"
  }'
```

**Resposta de Sucesso (200):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 900,
  "user": {
    "id": 1,
    "email": "novousuario@example.com",
    "username": "novousuario"
  }
}
```

**Resposta de Erro - Email/Senha Inválidos (401):**
```json
{
  "error": "invalid_credentials",
  "description": "Email ou senha incorretos"
}
```

## Fluxo Completo de Teste

1. **Registrar um novo usuário:**
```bash
curl -X POST http://localhost:8001/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "teste@example.com",
    "password": "senha123456",
    "username": "teste"
  }'
```

2. **Login com as credenciais:**
```bash
curl -X POST http://localhost:8001/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "teste@example.com",
    "password": "senha123456"
  }'
```

3. **Guardar o access_token e usar em requisições futuras**

## Implementação Atual

- ✅ Register com validação e banco de dados real
- ✅ Login com JWT token (conectado ao banco)
- ✅ Validação de campos
- ✅ Tratamento de erros padronizado
- ✅ Hash de senha com bcrypt
- ✅ Migrations automáticas

## Próximos Passos

- [ ] Implementar Refresh Token
- [ ] Implementar Validate Token
- [ ] Implementar Logout
- [ ] Adicionar Middleware de autenticação
- [ ] Adicionar testes unitários
- [ ] Dockerizar

