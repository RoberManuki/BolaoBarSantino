document.addEventListener('DOMContentLoaded', function() {
    const loadTimesTable = false;

    // Carregar os times antes de carregar as partidas
    loadTimes(loadTimesTable).then(() => {
        loadPartidas(); // Chama loadPartidas após os times serem carregados
        console.log("teste times -->", times);
    });

    // Função para editar uma partida
    window.editPartida = function(id) {
        fetch(`/api/partidas/${id}`)
            .then(response => response.json())
            .then(partida => {
                document.getElementById('partidaId').value = partida.id;
                document.getElementById('timeCasa').value = partida.time_casa;
                document.getElementById('casaGols').value = partida.casa_gols;
                document.getElementById('timeFora').value = partida.time_fora;
                document.getElementById('foraGols').value = partida.fora_gols;
                document.getElementById('data').value = partida.data;
                document.getElementById('vencedor').value = partida.vencedor;
            });
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

    CarregarDropRodada()

    // Carregamento inicial
    dropRodada.value = 1; 
    document.getElementById('rodadaDisplay').textContent = dropRodada.value;
    loadPartidas(dropRodada.value);

    // Atualizar o título e carregar as partidas quando o usuário selecionar uma rodada
    dropRodada.addEventListener('change', function() {
        const rodadaSelecionada = dropRodada.value;
        document.getElementById('rodadaDisplay').textContent = rodadaSelecionada;
        loadPartidas(rodadaSelecionada); 
    });
});

function CarregarDropRodada(){
    const dropRodada = document.getElementById('dropRodada');

    // Preencher o dropdown com as rodadas de 1 a 38
    for (let i = 1; i <= 38; i++) {
        let option = document.createElement('option');
        option.value = i;
        option.textContent = `Rodada ${i}`;
        dropRodada.appendChild(option);
    }
};
