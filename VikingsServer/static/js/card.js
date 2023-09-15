const photos = document.querySelectorAll('.photo');
photos.forEach(photo => {
    photo.addEventListener('click', function () {
        const photoID = this.id.split('-')[1];
        const newURL = `/cities?city=${photoID}`;
        window.location.href = newURL;
    });
});

const circles = document.querySelectorAll('.circle');
circles.forEach(circle => {
    circle.addEventListener('click', function (event) {
        event.stopPropagation();
        const card = this.closest('.card');
        const cityID = card.getAttribute('data-city-id');
        if (cityID !== null) {
            const confirmDelete = confirm('Вы уверены, что хотите удалить этот город?');

            if (confirmDelete) {
                const deleteURL = `api/v3/cities`;
                const requestData = {"id": parseInt(cityID)};
                fetch(deleteURL, {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(requestData)
                })
                    .then(response => {
                        if (response.status === 200) {
                            console.log(`Город с ID ${cityID} удален успешно.`);
                            window.location.reload();
                        } else {
                            console.error(`Ошибка при удалении города с ID ${cityID}.`);
                        }
                    })
                    .catch(error => {
                        console.error('Произошла ошибка при выполнении запроса:', error);
                    });
            }
        } else {
            console.error('cityID равен null.');
        }
    });
});