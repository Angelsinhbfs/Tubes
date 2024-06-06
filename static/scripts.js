// scripts.js

let page = 1; // Track the current page

// Function to load more content
function loadMoreContent() {
    // Simulate an AJAX request
    setTimeout(() => {
        const content = document.getElementById('tube');
        for (let i = 0; i < 3; i++) {
            const newItem = document.createElement('div');
            newItem.className = 'item';
            newItem.textContent = `Item ${page * 3 + i + 1}`;
            content.appendChild(newItem);
        }
        page++;
    }, 1000); // Simulate network delay
}

// Function to check if the user has scrolled to the bottom
function checkScroll() {
    if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
        loadMoreContent();
    }
}

// Attach the scroll event listener
window.addEventListener('scroll', checkScroll);