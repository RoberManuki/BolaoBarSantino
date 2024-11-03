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
                document.getElementById('rodada').value = partida.rodada;
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
});
