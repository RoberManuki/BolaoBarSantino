document.addEventListener('DOMContentLoaded', function() {

    const dropRodada = document.getElementById('dropRodada');
    
    // window.onload = function() {
    //     const rodadaAtual = sessionStorage.getItem('rodadaAtual');
    //     if (rodadaAtual) {
    //         loadPartidas(parseInt(rodadaAtual));
    //         sessionStorage.removeItem('rodadaAtual');
    //     }
    // };

    window.editPartida = function(id) {
        window.location.href = `/partida/formulario?id=${id}`;
    };

    window.deletePartida = function(id) {
        event.preventDefault(); 
        
        if (confirm('Tem certeza que deseja excluir essa partida?')) {
            fetch(`/api/partidas/${id}`, { method: 'DELETE' })
                .then(response => {
                    if (response.ok) {
                        alert('Partida excluída com sucesso!');
                        loadPartidas(dropRodada.value);
                    } else {
                        alert('Erro ao excluir a partida');
                    }
                })
                .catch(err => {
                    alert('Erro de rede: ' + err);
                });
        }
    };    

    // Atualizar o título e carregar as partidas quando o usuário selecionar uma rodada
    dropRodada.addEventListener('change', function() {
        const rodadaSelecionada = dropRodada.value;
        document.getElementById('rodadaDisplay').textContent = rodadaSelecionada;
        loadPartidas(rodadaSelecionada);
    });

    CarregarDropRodada();
    CarregamentoInicial();
});

function CarregamentoInicial() {
    const loadTimesTable = false;
    dropRodada.value = 1; 
    document.getElementById('rodadaDisplay').textContent = dropRodada.value;

    loadTimes(loadTimesTable).then(() => {
        loadPartidas(dropRodada.value);
    });
};

function CarregarDropRodada() {
    // Preencher o dropdown com as rodadas de 1 a 38
    for (let i = 1; i <= 38; i++) {
        let option = document.createElement('option');
        option.value = i;
        option.textContent = `Rodada ${i}`;
        dropRodada.appendChild(option);
    };
};
