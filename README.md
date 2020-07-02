# Microserviço de login em golang

- Utilizando RabbitMQ
- Banco de dados Mysql
- Micro serviço de envio de email (já criado)!!!
- Linguagem Golang

# Logins e senhas padroes
- Mysql: Usuario: root   password: 123456
- RabbitMQ: Usuario: user  password: bitnami

# Fluxo

- Usuário se registra
- Recebe email de confirmação de cadastro para ativar conta e validar email.
- Quando a conta é ativada o sistema cria um token para recuperação de senha.
- Usuário loga no sistema caso os dados estão certos ele recebe um jwt que vai ser usado em todo processo para percorrer pelos outros subsistemas.
- Ao deslogar esse token é destruido.
- Caso ele logue novamente e não tenha deslogado o antigo token é destruido e o um novo é criado.
- Recuperação de senha.
- Bloqueio de conta.
- Ativação de conta administrador.

# Bibliotecas
- ORM - go get github.com/jinzhu/gorm
- Encrypt - go get golang.org/x/crypto/bcryp
- UUID - go get github.com/satori/go.uuid