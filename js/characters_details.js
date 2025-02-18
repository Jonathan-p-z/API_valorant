function showAbilityDetails(name, description, image, videoURL) {
    document.getElementById('ability-name').innerText = name;
    document.getElementById('ability-description').innerText = description;
    document.getElementById('ability-image').src = image;
    document.getElementById('ability-video').src = videoURL;
    document.getElementById('ability-details').style.display = 'block';
}