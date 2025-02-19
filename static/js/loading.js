window.addEventListener("load", function() {
    const loadingScreen = document.getElementById("loadingScreen");

    // Prevent multiple reloads
    if (sessionStorage.getItem("loading")) {
        window.location.href = "/home";
    } else {
        sessionStorage.setItem("loading", "true");
    }

    // Function to preload images
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

    // URLs of images to preload
    const imageUrls = [
        "/static/img/cfae886e263126f685510e2f45b82970.jpg",
        "/static/img/MastersBangkok.avif",
        "/static/img/013dc5d8874cce55d4aa7b678049bcfaa52f127b-1920x1080.jpg",
        "/static/img/abyss.webp"
    ];

    // Start loading data in the background
    fetch("/api/load-data")
        .then(response => response.json())
        .then(data => {
            console.log("Data loaded:", data);
            // Preload images
            preloadImages(imageUrls, () => {
                // Hide loading screen and redirect to the homepage
                loadingScreen.style.opacity = "0";
                setTimeout(() => {
                    loadingScreen.style.display = "none";
                    window.location.href = "/home";
                }, 1000);
            });
        })
        .catch(error => {
            console.error("Error loading data:", error);
            // Hide loading screen and redirect to the homepage even if there is an error
            loadingScreen.style.opacity = "0";
            setTimeout(() => {
                loadingScreen.style.display = "none";
                window.location.href = "/home";
            }, 1000);
        });
});