$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

function criarPublicacao(evento) {
    evento.preventDefault();
    //ajax que faz a requisição pro servidor e obtem resposta. Ele sabe se done ou fail pelo statuscode
    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data:{
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        },
        dataType:'text'
    }).done(function(){
        window.location="/home";
    }).fail(function(){
        alert("Erro ao criar publicação");
    })
}

function curtirPublicacao(evento) {
    //obtendo id da publicação clicada pegando a div mais perto
    evento.preventDefault();
    const clicado = $(evento.target);
    const publicacaoId = clicado.closest('div').data('publicacao-id');
    //desativando o botao de curtir até que a função de curtir acabe
    clicado.prop('disabled', true);
    //fazendo requisição de curtida
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function(){
        //aumentando numero de curtidas na tela na hora
        const contCurtidas = clicado.next('span');
        const qtdCurtidas = parseInt(contCurtidas.text());
        contCurtidas.text(qtdCurtidas + 1 );
        //após isso adcionando nova classe para caso o usuário clique dnv descurta
        //perceba que só funcionará se a a página não for atualizada
        clicado.addClass('descurtir-publicacao');
        clicado.addClass('text-danger');
        clicado.removeClass('curtir-publicacao');
    }).fail(function(){
        alert("Erro ao curtir publicação");
    }).always(function(){
        clicado.prop('disabled', false);
    });
}

function descurtirPublicacao(evento) {
    //obtendo id da publicação clicada pegando a div mais perto
    evento.preventDefault();
    const clicado = $(evento.target);
    const publicacaoId = clicado.closest('div').data('publicacao-id');
    //desativando o botao de curtir até que a função de descurtir acabe
    clicado.prop('disabled', true);
    //fazendo requisição de curtida
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function(){
        //diminuindo numero de curtidas na tela na hora
        const contCurtidas = clicado.next('span');
        const qtdCurtidas = parseInt(contCurtidas.text());
        contCurtidas.text(qtdCurtidas - 1 );
        //após isso adcionando nova classe para caso o usuário clique dnv descurta
        //perceba que só funcionará se a a página não for atualizada
        clicado.removeClass('descurtir-publicacao');
        clicado.removeClass('text-danger');
        clicado.addClass('curtir-publicacao');
    }).fail(function(){
        alert("Erro ao descurtir publicação");
    }).always(function(){
        clicado.prop('disabled', false);
    });
}

function atualizarPublicacao(evento) {
    //desativando botão até que a função seja executada
    $(this).prop('disabled', true);
    const publicacaoId = $(this).data('publicacao-id');
    //ajax que faz a requisição pro servidor e obtem resposta. Ele sabe se done ou fail pelo statuscode
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(){
        Swal.fire({
            title: "Feito!",
            text: "Publicação editada com sucesso!",
            icon: "success"
          }).then(function(){
            window.location="/home"
          })
    }).fail(function(){
        alert("Erro ao editar publicação");
    }).always(function(){
        $('#atualizar-publicacao').prop('disabled', false);
    })
}

function deletarPublicacao(evento){
    //obtendo id da publicação clicada pegando a div mais perto
    evento.preventDefault();
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que quer excluir essa publicação?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao){
        if (!confirmacao.value) return;
        const clicado = $(evento.target);
        const publicacao = clicado.closest('div');
        const publicacaoId = publicacao.data('publicacao-id');
        //desativando o botao de deletar até que a função de deletar acabe
        clicado.prop('disabled', true);
        //fazendo requisição de deletar
        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function(){
            publicacao.fadeOut("slow", function(){
                $(this).remove();
            });
        }).fail(function(){
            alert("Erro ao deletar publicação");
        });
    })
}