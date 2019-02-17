#language: pt

Funcionalidade: Login
    Para que eu possa ter acesso as minhas tarefas
    Sendo um usuario
    Posso me autenticar com os meus dados previamente cadastrados

    Contexto: Formulario de login
        Dado que eu acessei a pagina principal

    @success
    Cenario: Login do usuario
        Quando faco login com "euricardo@getnada.com" e "Passw0rd!"
        Entao sou atenticado com sucesso

    Cenario: Senha incorreta
        Quando faco login com "euricardo@getnada.com" e "xpto123"
        Entao eu devo ver a seguinte mensagem "Wrong username or password. Please retry"
    
    Cenario: Email invalido
        Quando faco login com "euricardo#getnada.com" e "xpto123"
        Entao eu devo ver a seguinte mensagem "Wrong username or password. Please retry"