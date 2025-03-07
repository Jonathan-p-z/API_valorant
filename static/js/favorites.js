function addFavorite(id, name, image, type) {
    fetch('/api/favorites', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ id, name, image, type })
    })
    .then(response => {
        if (response.ok) {
            alert('Ajouté aux favoris');
            location.reload();
        } else {
            alert('Erreur lors de l\'ajout aux favoris');
        }
    })
    .catch(error => console.error('Erreur:', error));
}

function removeFavorite(id) {
    fetch(`/api/remove-favorite?id=${id}`, {
        method: 'DELETE'
    })
    .then(response => {
        if (response.ok) {
            alert('Supprimé des favoris');
            document.getElementById(`favorite-${id}`).remove();
        } else {
            alert('Erreur lors de la suppression des favoris');
        }
    })
    .catch(error => console.error('Erreur:', error));
}

function toggleFavorite(id, name, image, type) {
    const favoriteItem = document.getElementById(`favorite-${id}`);
    if (favoriteItem) {
        removeFavorite(id);
    } else {
        addFavorite(id, name, image, type);
    }
}

document.addEventListener('DOMContentLoaded', () => {
    fetch('/api/favorites')
        .then(response => response.json())
        .then(favorites => {
            const favoritesList = document.getElementById('favorites-list');
            favorites.forEach(favorite => {
                const favoriteItem = document.createElement('div');
                favoriteItem.id = `favorite-${favorite.id}`;
                favoriteItem.classList.add('favorite-item');
                favoriteItem.innerHTML = `
                    <div class="favorite-star" onclick="toggleFavorite('${favorite.id}', '${favorite.name}', '${favorite.image}', '${favorite.type}')">&#9733;</div>
                    <img src="${favorite.image}" alt="${favorite.name}" class="favorite-image">
                    <h3>${favorite.name}</h3>
                    <button onclick="removeFavorite('${favorite.id}')">Supprimer</button>
                `;
                favoritesList.appendChild(favoriteItem);
            });
        })
        .catch(error => console.error('Erreur lors du chargement des favoris:', error));
});