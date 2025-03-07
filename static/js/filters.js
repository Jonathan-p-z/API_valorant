function filterWeapons() {
    const weaponType = document.getElementById('weapon-type').value;
    const fireRate = document.getElementById('fire-rate').value;
    const weaponCards = document.querySelectorAll('.weapon-card');

    let hasVisibleWeapons = false;

    weaponCards.forEach(card => {
        const cardType = card.getAttribute('data-type');
        const cardFireRate = parseFloat(card.getAttribute('data-fire-rate'));

        let typeMatch = !weaponType || weaponType === "Tous" || cardType === weaponType;
        let fireRateMatch = !fireRate || (fireRate === 'above' && cardFireRate > 8) || (fireRate === 'below' && cardFireRate < 8);

        if (typeMatch && fireRateMatch) {
            card.style.display = 'block';
            hasVisibleWeapons = true;
        } else {
            card.style.display = 'none';
        }
    });

    // Hide categories with no visible weapons
    const weaponCategories = document.querySelectorAll('.weapons-category');
    weaponCategories.forEach(category => {
        const visibleWeapons = category.querySelectorAll('.weapon-card[style="display: block;"]');
        if (visibleWeapons.length === 0) {
            category.style.display = 'none';
        } else {
            category.style.display = 'block';
        }
    });

    // Display a message if no weapons are visible
    const noResultsMessage = document.getElementById('no-results-message');
    if (!hasVisibleWeapons) {
        noResultsMessage.style.display = 'block';
    } else {
        noResultsMessage.style.display = 'none';
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const filterButton = document.querySelector('.filter-form button[type="submit"]');

    filterButton.addEventListener('click', (event) => {
        event.preventDefault();
        filterWeapons();
    });
});

