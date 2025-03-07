let progress = 0;
const progressText = document.getElementById('progressText');
const circle = document.querySelector('.progress-ring__circle');
const radius = circle.r.baseVal.value;
const circumference = 2 * Math.PI * radius;

circle.style.strokeDasharray = `${circumference} ${circumference}`;
circle.style.strokeDashoffset = circumference;

function setProgress(percent) {
    const offset = circumference - (percent / 100) * circumference;
    circle.style.strokeDashoffset = offset;
    progressText.textContent = `${percent}%`;
}

function updateProgress() {
    progress += Math.random() * 10;
    if (progress > 100) progress = 100;
    setProgress(Math.floor(progress));

    if (progress < 100) {
        setTimeout(updateProgress, 500);
    } else {
        window.location.href = '/home';
    }
}

updateProgress();