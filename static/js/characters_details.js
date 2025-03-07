function showAbilityDetails(name, description, image, videoURL, isVideo) {
    const abilityDetails = document.getElementById('ability-details');
    const abilityName = document.getElementById('ability-name');
    const abilityImage = document.getElementById('ability-image');
    const abilityDescription = document.getElementById('ability-description');
    const abilityVideo = document.getElementById('ability-video');
    const abilityVideoSource = document.getElementById('ability-video-source');

    abilityName.textContent = name;
    abilityDescription.textContent = description;

    if (isVideo) {
        abilityImage.style.display = 'none';
        abilityVideo.style.display = 'block';
        abilityVideoSource.src = videoURL;
        abilityVideo.load();
    } else {
        abilityImage.style.display = 'block';
        abilityImage.src = image;
        abilityVideo.style.display = 'none';
    }

    abilityDetails.style.display = 'block';
}