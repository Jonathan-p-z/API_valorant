window.addEventListener("load", function() {
    const loadingScreen = document.getElementById("loadingScreen");
    if (sessionStorage.getItem("loading")) {
        window.location.href = "/home";
    } else {
        sessionStorage.setItem("loading", "true");
    }

    function preloadImages(urls, callback) {
        let loadedCount = 0;
        const total = urls.length;

        urls.forEach(url => {
            const img = new Image();
            img.src = url;
            img.onload = () => {
                loadedCount++;
                if (loadedCount === total) {
                    callback();
                }
            };
            img.onerror = () => {
                loadedCount++;
                if (loadedCount === total) {
                    callback();
                }
            };
        });
    }

    const imageUrls = [
        "/static/img/cfae886e263126f685510e2f45b82970.jpg",
        "/static/img/MastersBangkok.avif",
        "/static/img/013dc5d8874cce55d4aa7b678049bcfaa52f127b-1920x1080.jpg",
        "/static/img/abyss.webp"
    ];
    
    function loadData(callback) {
        fetch("/api/load-data")
            .then(response => response.json())
            .then(data => {
                console.log("Data loaded:", data);
                callback();
            })
            .catch(error => {
                console.error("Error loading data:", error);
                callback();
            });
    }

    loadData(() => {
        preloadImages(imageUrls, () => {
            loadingScreen.style.opacity = "0";
            setTimeout(() => {
                loadingScreen.style.display = "none";
                window.location.href = "/home";
            }, 1000);
        });
    });
});

document.addEventListener("DOMContentLoaded", function() {
    var progressBar = document.querySelector(".progress");
    var width = 0;
    var interval = setInterval(function() {
        if (width >= 100) {
            clearInterval(interval);
            window.location.href = "/home"; // Redirect to home after loading
        } else {
            width += 1;
            progressBar.style.width = width + "%";
        }
    }, 50); // Adjust the interval time to control the speed of the progress bar
});