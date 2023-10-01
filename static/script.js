// static/script.js

function handleSubmitForUpload(event) {
    console.log("upload success")
} 

// Add an event listener to the form
const uploadForm = document.getElementById('uploadForm')
if (uploadForm){
    uploadForm.addEventListener('submit', handleSubmitForUpload)
} 

// Function to handle button click
function sayHello() {
    alert('Hello from JavaScript!');
}

// Function to change the background color of the header
function changeHeaderColor() {
    const header = document.querySelector('header');
    if (header) {
        header.style.backgroundColor = getRandomColor();
    }
}

// Generate a random color for background
function getRandomColor() {
    const letters = '0123456789ABCDEF';
    let color = '#';
    for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
}
