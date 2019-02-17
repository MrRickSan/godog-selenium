#language: pt

Funcionalidade: Login
    Para que eu possa ter acesso as minhas tarefas
    Sendo um usuario
    Posso me autenticar com os meus dados previamente cadastrados

    Contexto: Formulario de login
        Dado que eu acessei a pagina principal

    Cenario: Login do usuario
        Quando faco login com "eu@ricardo.com" e "123456"
        Entao sou atenticado com sucesso

    Cenario: Senha incorreta
        Quando faco login com "eu@ricardo.com" e "xpto123"
        Entao eu devo ver a seguinte mensagem "Senha invalida"
    
    Cenario: Email invalido
        Quando faco login com "eu#ricardo.com" e "xpto123"
        Entao eu devo ver a seguinte mensagem "Senha invalida"