    document.addEventListener('DOMContentLoaded', () => {
    // Animation pour les éléments de la grille d'actualités
    const newsItems = document.querySelectorAll('.news-item');
    newsItems.forEach((item, index) => {
        item.style.opacity = '0';
        item.style.transform = 'translateY(20px)';
        item.style.transition = `opacity 0.5s ${index * 0.2}s, transform 0.5s ${index * 0.2}s`;
    });

    setTimeout(() => {
        newsItems.forEach(item => {
            item.style.opacity = '1';
            item.style.transform = 'translateY(0)';
        });
    }, 500);
});