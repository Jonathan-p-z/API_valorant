function showAbilityDetails(name, description, image, videoURL, isVideo) {
    console.log("Name:", name);
    console.log("Description:", description);
    console.log("Image:", image);
    console.log("VideoURL:", videoURL);
    console.log("IsVideo:", isVideo);

    // Fermer les détails de la compétence actuellement ouverte
    document.getElementById('ability-details').style.display = 'none';
    document.getElementById('ability-video').style.display = 'none';
    document.getElementById('ability-image').style.display = 'none';

    // Mettre à jour les détails de la nouvelle compétence
    document.getElementById('ability-name').innerText = name;
    document.getElementById('ability-description').innerText = description;
    if (isVideo) {
        document.getElementById('ability-video').style.display = 'block';
        document.getElementById('ability-video-source').src = videoURL;
        document.getElementById('ability-video').load(); // Recharger la vidéo pour appliquer la nouvelle source
    } else {
        document.getElementById('ability-image').style.display = 'block';
        document.getElementById('ability-image').src = videoURL;
    }
    document.getElementById('ability-details').style.display = 'block';
}