$('#login').on('submit',login);

function login(evento){
    evento.preventDefault();
    //ajax que faz a requisição pro servidor e obtem resposta. Ele sabe se done ou fail pelo statuscode
    $.ajax({
        url: "/login",
        method: "POST",
        data:{
            email: $('#email').val(),
            senha: $('#senha').val(),
        },
        dataType:'text'
    }).done(function(){
        window.location = "/home";
    }).fail(function(data, textStatus, jqXHR){
        console.log(data);
        console.log(textStatus);
        console.log(jqXHR);
        Swal.fire({
            title: "Senha ou usuário incorretos!",
            text: "Qual será?",
            icon: "error"
          });
    }).always(function(){
        console.log("terminou");
    })
}