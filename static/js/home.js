document.addEventListener('DOMContentLoaded', function() {
    const API_URL = '/api/times';

    // Função para carregar os times
    function loadTimes() {
        fetch(API_URL)
            .then(response => {
                if (!response.ok) {
                    throw new Error(fmt.Sprintf("Network response was not ok for: %s", API_URL));
                }
                return response.json();
            })
            .then(data => {
                // console.log('Dados recebidos:', data);
                data.forEach(time => {
                    console.log(time);
                });

                // times chegando --> ver como mostrar
            })
            .catch(error => {
                console.error('Erro ao carregar partidas:', error);
            });
    }

    loadTimes();
});
