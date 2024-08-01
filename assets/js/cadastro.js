$('#formulario-cadastro').on('submit',criarusuario);

function criarusuario(evento){
    evento.preventDefault();
    if ($('#senha').val() != $('#confirmaSenha').val()){
        Swal.fire("Opa", "Senhas não coinscidem", "error");
        return;
    }
    //ajax que faz a requisição pro servidor e obtem resposta. Ele sabe se done ou fail pelo statuscode
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data:{
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function(){
        Swal.fire({
            title: "Sucesso",
            text: "Cadastro realizado com sucesso!",
            icon: "success"
          }).then(function(){
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    senha: $('#senha').val(),
                }
            }).done(function(){
                window.location = "/home";
            }).fail(function(){
                alert("Problemas no servidor. Falha ao redirecionar para home. Favor faça o login");
            })
          })
    }).fail(function(erro){
        console.log(erro);
        Swal.fire("Erro", "Programador foi preguiçoso demais, ou o email ou o nick já estão em uso", "error");
    });

}