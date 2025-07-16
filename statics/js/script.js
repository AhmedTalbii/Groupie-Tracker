var hoverSound = document.querySelector(".hoverSound");
var cards = document.querySelectorAll(".artistCard");

function onHover() {
    hoverSound.currentTime = 0;
    hoverSound.play();
}

cards.forEach(function (card) {
    card.addEventListener("mouseenter", onHover);
});