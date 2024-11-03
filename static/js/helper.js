let times = [];

const API_URL_TIMES = '/api/times';
const API_URL_PARTIDAS = '/api/partidas';

// Função para carregar os times
function loadTimes(loadTable) {
    return fetch(API_URL_TIMES)  // Retorna a Promise do fetch
        .then(response => {
            if (!response.ok) {
                throw new Error(`Network response was not ok for: ${API_URL_TIMES}`);
            }
            return response.json();
        })
        .then(data => {
            times = data;
            console.log('Dados recebidos:', times);
            if (loadTable) loadTimesTable(times);
        })
        .catch(error => {
            console.error('Erro ao carregar os times:', error);
        });
}

function loadTimesTable(times) {
    timesTable.innerHTML = '';
            
    times.forEach(time => {
        const row = timesTable.insertRow();
        row.innerHTML = `
            <td>${time.id}</td>
            <td>${time.nome}</td>
            <td>${time.cidade}</td>
        `;
    });
}

// Função para carregar todas as partidas
function loadPartidas() {
    const partidasTable = document.getElementById('partidasTable').getElementsByTagName('tbody')[0];

    fetch(API_URL_PARTIDAS)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Network response was not ok for: ${API_URL_PARTIDAS}`);
            }
            return response.json();
        })
        .then(data => {
            partidasTable.innerHTML = '';
            data.forEach(partida => {
                //console.log('Partida:', partida); // Log da partida
                const timeCasa = times.find(time => time.id === Number(partida.time_casa));
                const timeFora = times.find(time => time.id === Number(partida.time_fora));
                
                const row = partidasTable.insertRow();
                row.innerHTML = `
                    <td>${partida.id}</td>
                    <td>${partida.rodada}</td>
                    <td>${timeCasa ? timeCasa.nome : 'Desconhecido'}</td>
                    <td>${partida.casa_gols}</td>
                    <td>${timeFora ? timeFora.nome : 'Desconhecido'}</td>
                    <td>${partida.fora_gols}</td>
                    <td>${partida.data}</td>
                    <td>${partida.vencedor}</td>
                    <td>
                        <button onclick="editPartida(${partida.id})">Editar</button>
                        <button onclick="deletePartida(${partida.id})">Excluir</button>
                    </td>
                `;
            });
        })
        .catch(error => {
            console.error('Erro ao carregar partidas:', error);
        });
}