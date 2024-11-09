// OnInit()
document.addEventListener('DOMContentLoaded', function () {
    CarregarDropsTimes();
    CarregarDadosEdicao();
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
    });

};

function CarregarDadosEdicao() {

    const partidaIdInput = document.getElementById('partidaId');
    const urlParams = new URLSearchParams(window.location.search);
    const partidaId = urlParams.get('id');

    if (partidaId) {
        document.getElementById('formTitle').textContent = 'Editar Partida';

        fetch(`/api/partidas/${partidaId}`)
            .then(response => response.json())
            .then(partida => {
                partidaIdInput.value = partida.id;
                document.getElementById('casaGols').value = partida.casa_gols;
                document.getElementById('foraGols').value = partida.fora_gols;
                document.getElementById('data').value = partida.data;
                document.getElementById('timeCasa').value = partida.time_casa;
                document.getElementById('timeFora').value = partida.time_fora;
            });
    };

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
                toastr.success("Partida criada com sucesso!");
                
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
            console.error('Erro:', error);
            toastr.error("Ocorreu um erro ao criar a partida.");
        });
    });

};


