const ArtistsCards = document.querySelectorAll('.artistCard');

ArtistsCards.forEach((card) => {
    const modal = card.querySelector(".infoBox");
    const closeBtn = modal.querySelector(".close");

    card.addEventListener("click", showModal);
    function showModal() {
        modal.showModal();
    }

    closeBtn.addEventListener("click", (event) => {
        event.stopPropagation();
        modal.close();
    });
});
