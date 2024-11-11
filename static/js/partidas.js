document.addEventListener('DOMContentLoaded', function() {

    // Função para editar uma partida
    window.editPartida = function(id) {
        fetch(`/api/partidas/${id}`)
        .then(response => {
            if (!response.ok) {
                console.log("!response.ok");
                throw new Error(`Erro ao buscar a partida com id ${id}`);
            }
            console.log("response", response)
            return response.json();
        })
        .then(partida => {
            partidaIdInput.value = partida.id;
            document.getElementById('casaGols').value = partida.casa_gols;
            document.getElementById('foraGols').value = partida.fora_gols;
            document.getElementById('data').value = partida.data;
            document.getElementById('timeCasa').value = partida.time_casa;
            document.getElementById('timeFora').value = partida.time_fora;
        })
        .catch(error => {
            console.error('Erro ao editar partida:', error);
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
    }
};
