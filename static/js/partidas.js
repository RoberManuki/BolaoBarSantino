document.addEventListener('DOMContentLoaded', function() {
    const partidasTable = document.getElementById('partidasTable').getElementsByTagName('tbody')[0];

    const API_URL = '/api/partidas'; // Atualizado para refletir a nova rota da API

    // Função para carregar todas as partidas
    function loadPartidas() {
        fetch(API_URL)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                console.log('Dados recebidos:', data); // Adicione este log para verificar os dados
                partidasTable.innerHTML = '';
                data.forEach(partida => {
                    const row = partidasTable.insertRow();
                    row.innerHTML = `
                        <td>${partida.id}</td>
                        <td>${partida.rodada}</td>
                        <td>${partida.time_casa}</td>
                        <td>${partida.casa_gols}</td>
                        <td>${partida.time_fora}</td>
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

    // Carregar partidas quando a página for carregada
    loadPartidas();
});
