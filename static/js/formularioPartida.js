// OnInit()
document.addEventListener('DOMContentLoaded', function () {
    CarregarDropsTimes();
    CriarPartidaClick();
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

        // Preencher os valores selecionados caso esteja editando a partida
        if (document.getElementById('partidaId').value) {
            const partidaId = document.getElementById('partidaId').value;
            fetch(`/api/partidas/${partidaId}`)
                .then(response => response.json())
                .then(partida => {
                    document.getElementById('timeCasa').value = partida.time_casa;
                    document.getElementById('timeFora').value = partida.time_fora;
                });
        }
    });
};

//sepa nao precisa
function ValidarTimesNaRodada(rodada, timeCasa, timeFora) {
    return fetch(`/api/partidas/validar?rodada=${rodada}&timeCasa=${timeCasa}&timeFora=${timeFora}`)
        .then(response => response.json())
        .then(data => {
            if (data.jaJogaram) {
                toastr.error("Os times já se enfrentaram nesta rodada.");
                return false;
            }
            return true;
        });
};

function CriarPartidaClick() {
    const formPartida = document.getElementById('formPartida');
    formPartida.addEventListener('submit', function (e) {
        e.preventDefault();

        // Captura os valores dos campos no momento do submit
        const rodada = parseInt(document.getElementById('rodada').value) || 0;
        const timeCasa = parseInt(document.getElementById('timeCasa').value) || 0;
        const timeFora = parseInt(document.getElementById('timeFora').value) || 0;
        const casaGols = parseInt(document.getElementById('casaGols').value) || 0;
        const foraGols = parseInt(document.getElementById('foraGols').value) || 0;
        const vencedor = (casaGols > foraGols) ? "Casa" : (foraGols > casaGols) ? "Fora" : "Empate";
        const data = document.getElementById('data').value.split("T")[0]; // Pega apenas a parte da data

        ValidarTimesNaRodada(rodada, timeCasa, timeFora).then(isValid => {
            if (!isValid) return;

            const partida = {
                rodada: rodada,
                time_casa: timeCasa,
                casa_gols: casaGols,
                time_fora: timeFora,
                fora_gols: foraGols,
                data: data,               // (YYYY-MM-DD)
                vencedor: vencedor,       // Pode ser "Casa", "Fora" ou "Empate"
            };

            const method = partida.id ? 'PUT' : 'POST';
            const url = partida.id ? `/api/partidas/${partida.id}` : '/api/partidas';

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



