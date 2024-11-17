document.addEventListener('DOMContentLoaded', function() {
    // Função para editar uma partida
    window.editPartida = function(id) {
        window.location.href = `/partida/formulario?id=${id}`;
    };

    // Função para excluir uma partida
    window.deletePartida = function(id) {
        fetch(`/api/partidas/${id}`, { method: 'DELETE' })
            .then(response => {
                if (response.ok) {
                    loadPartidas();
                }
            });
    };

    // Atualizar o título e carregar as partidas quando o usuário selecionar uma rodada
    dropRodada.addEventListener('change', function() {
        const rodadaSelecionada = dropRodada.value;
        document.getElementById('rodadaDisplay').textContent = rodadaSelecionada;
        loadPartidas(rodadaSelecionada); 
    });

    CarregarDropRodada()
    CarregamentoInicial()
});

function CarregamentoInicial(){
    const loadTimesTable = false;
    dropRodada.value = 1; 
    document.getElementById('rodadaDisplay').textContent = dropRodada.value;

    loadTimes(loadTimesTable).then(() => {
        loadPartidas(dropRodada.value);
    });
};

function CarregarDropRodada(){
    const dropRodada = document.getElementById('dropRodada');

    // Preencher o dropdown com as rodadas de 1 a 38
    for (let i = 1; i <= 38; i++) {
        let option = document.createElement('option');
        option.value = i;
        option.textContent = `Rodada ${i}`;
        dropRodada.appendChild(option);
    };
};
