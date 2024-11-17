document.addEventListener('DOMContentLoaded', function () {
    CarregarDropsTimes();  // Carrega os times nos selects
    CarregarPartidaParaEdicao();  // Carrega os dados da partida se houver um ID
    CriarPartidaClick(); // Cria ou atualiza a partida
});

function CarregarDropsTimes() {
    const timeCasaSelect = document.getElementById('timeCasa');
    const timeForaSelect = document.getElementById('timeFora');

    loadTimes(false).then(() => {
        times.forEach(time => {
            const optionCasa = document.createElement('option');
            optionCasa.value = time.id;
            optionCasa.textContent = time.nome;
            timeCasaSelect.appendChild(optionCasa);

            const optionFora = document.createElement('option');
            optionFora.value = time.id;
            optionFora.textContent = time.nome;
            timeForaSelect.appendChild(optionFora);
        });
    });
};

// Função para carregar os dados da partida para edição
function CarregarPartidaParaEdicao() {
    const urlParams = new URLSearchParams(window.location.search);
    const partidaId = urlParams.get('id');

    if (partidaId) {
        fetch(`/api/partidas/${partidaId}`)
            .then(response => response.json())
            .then(partida => {
                // Preenche os campos do formulário com os dados da partida
                document.getElementById('rodada').value = partida.rodada;
                document.getElementById('timeCasa').value = partida.time_casa;
                document.getElementById('casaGols').value = partida.casa_gols;
                document.getElementById('timeFora').value = partida.time_fora;
                document.getElementById('foraGols').value = partida.fora_gols;
                document.getElementById('data').value = partida.data;  // Ajuste de formato de data pode ser necessário
                document.getElementById('formTitle').textContent = `Editar Partida ${partidaId}`;
            })
            .catch(error => {
                console.error('Erro ao carregar a partida:', error);
            });
    };
};

function CriarPartidaClick() {
    const formPartida = document.getElementById('formPartida');
    formPartida.addEventListener('submit', function (e) {
        e.preventDefault();

        const urlParams = new URLSearchParams(window.location.search);
        const partidaId = urlParams.get('id');
        const edicao = partidaId ? true : false;
        const rodada = parseInt(document.getElementById('rodada').value) || 0;
        const timeCasa = parseInt(document.getElementById('timeCasa').value) || 0;
        const timeFora = parseInt(document.getElementById('timeFora').value) || 0;
        const casaGols = parseInt(document.getElementById('casaGols').value) || 0;
        const foraGols = parseInt(document.getElementById('foraGols').value) || 0;
        const vencedor = (casaGols > foraGols) ? "Casa" : (foraGols > casaGols) ? "Fora" : "Empate";
        const data = document.getElementById('data').value.split("T")[0];

        ValidarTimesNaRodada(rodada, timeCasa, timeFora, edicao).then(isValid => {
            if (!isValid) return;

            const partida = {
                id: partidaId,
                rodada: rodada,
                time_casa: timeCasa,
                casa_gols: casaGols,
                time_fora: timeFora,
                fora_gols: foraGols,
                data: data,
                vencedor: vencedor
            };

            const method = edicao ? 'PUT' : 'POST';
            const url = edicao ? `/api/partidas/${partidaId}` : '/api/partidas';

            fetch(url, {
                method: method,
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(partida)
            })
            .then(response => {
                if (response.ok) {
                    toastr.success("Partida salva com sucesso!");
                    setTimeout(function() {
                        window.location.href = '/';
                    }, 3500);
                } else {
                    return response.text().then(message => { 
                        toastr.error(message);
                    });
                }
            })
            .catch(error => {
                toastr.error(error);
            });
        });
    });
};
