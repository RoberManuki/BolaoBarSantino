let times = [];

// Função para carregar os times
function loadTimes(loadTable) {
    return fetch('/api/times')  // Retorna a Promise do fetch
        .then(response => {
            if (!response.ok) {
                throw new Error(`Network response was not ok for: ${'/api/times'}`);
            }
            return response.json();
        })
        .then(data => {
            times = data;
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

function loadPartidas(rodada) {
    const partidasTable = document.getElementById('partidasTable').getElementsByTagName('tbody')[0];

    fetch(`/api/partidas?rodada=${rodada}`)
        .then(response => response.json())
        .then(data => {
            partidasTable.innerHTML = '';
            data.forEach(partida => {
                const timeCasa = times.find(time => time.id === Number(partida.time_casa));
                const timeFora = times.find(time => time.id === Number(partida.time_fora));
                const vencedor = partida.casa_gols > partida.fora_gols ? 'Casa' : (partida.fora_gols > partida.casa_gols ? 'Fora' : 'Empate');

                const row = partidasTable.insertRow();
                row.innerHTML = `
                    <td>${partida.id}</td>
                    <td>${timeCasa ? timeCasa.nome : 'Desconhecido'}</td>
                    <td>${partida.casa_gols}</td>
                    <td>${timeFora ? timeFora.nome : 'Desconhecido'}</td>
                    <td>${partida.fora_gols}</td>
                    <td>${partida.data}</td>
                    <td>${vencedor}</td>
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

