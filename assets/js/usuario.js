$('#parar-de-seguir').on('click',pararDeSeguir);
$('#seguir').on('click',seguir);
$('#editar-usuario').on('submit', editar);
$('#editar-senha').on('submit', editarSenha);
$('#deletar-usuario').on('click', deletar);

function seguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);
    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: 'POST'
    }).done(function(){
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function(){
        $('#seguir').prop('disabled', false);
        alert("Erro seguir");
    })
}

function pararDeSeguir(){
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disable', true);
    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: 'POST'
    }).done(function(){
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function(){
        $('#parar-de-seguir').prop('disabled', false);
        alert("Erro ao parar de seguir");
    })
}

function editar(evento) {
    evento.preventDefault();
    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function(){
        Swal.fire({
            title: "Sucesso",
            text: "Edição realizada com sucesso!",
            icon: "success"
          }).then(function(){
            window.location = "/perfil";
          });
    }).fail(function(){
        Swal.fire("Erro", "Programador foi preguiçoso demais, ou o email ou o nick já estão em uso", "error");
    });
}

function editarSenha(evento){
    evento.preventDefault();
    if ($('#nova').val() != $('#confirma').val()){
        Swal.fire("Opa", "Senhas não coinscidem", "error");
        return;
    }
    $.ajax({
        url: "/editar-senha",
        method: "POST",
        data: {
            atual: $('#atual').val(),
            nova: $('#nova').val(),
        }
    }).done(function(){
        Swal.fire({
            title: "Sucesso",
            text: "Edição realizada com sucesso!",
            icon: "success"
          }).then(function(){
            window.location = "/perfil";
          });
    }).fail(function(){
        Swal.fire("Erro", "Senha incorreta!", "error");
    });
}

function deletar(){
    Swal.fire({
        title: "Atenção!",
        text: "Essa é uma ação irreversível!",
        showCancelButton: true,
        cancelButton: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao){
        if (confirmacao.value) {
            $.ajax({
                url: "/deletar-usuario",
                method: "DELETE"
            }).done(function(){
                Swal.fire({
                    title: "Sucesso",
                    text: "Exclusão realizada com sucesso!",
                    icon: "success"
                  }).then(function(){
                    window.location = "/logout";
                  })
            }).fail(function(){
                Swal.fire("Erro", "Ocorreu um erro", "error");
            });
        }
    })
}